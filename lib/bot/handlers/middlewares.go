package handlers

import (
	"context"
	"strings"
	"sync"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// The global middleware checks for the user's presence in the cache and PostgreSQL.
// If the user is absent, it initiates the registration or data recovery process.
// The retrieval of user has been changed.
// Now it is passed through the child context to the handlers, and the context key is stored in the database package, in the contextKey.go file.
// User check in the cache is now performed once in the global middleware
func (h *handlers) IdentifyUserMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(parentCtx context.Context, tgb *bot.Bot, update *models.Update) {
		var chatID int64
		var username string
		var ctxWithUser context.Context
		user, found := h.cache.GetUser(chatID)
		if found {
			ctxWithUser = context.WithValue(parentCtx, database.UserCtxKey, user)
			next(ctxWithUser, tgb, update)
			return
		}

		switch {
		case update.Message != nil:
			chatID = update.Message.From.ID
			username = update.Message.From.Username
		case update.CallbackQuery != nil:
			chatID = update.CallbackQuery.From.ID
			username = update.CallbackQuery.From.Username
		default:
			//Other message formats are not used, so I am exiting the function
			return
		}

		if h.dbService.CheckSession(chatID) {
			user, err := restoreUserSessionFromDB(h.dbService, chatID, update.Message.From.Username)
			if err != nil {
				log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to restore user from the database. The user has been sent for registration")
				h.handleNewUserRegistration(parentCtx, tgb, update)
				return
			}

			h.cache.SetUser(chatID, *user) // add user from persistent db into cache
			log.Info().Int64("chat_id", chatID).Msg("User session successfully restored from the database. User added to the cache.")
			ctxWithUser = context.WithValue(parentCtx, database.UserCtxKey, user)
			next(ctxWithUser, tgb, update)
			return
		}

		//TODO: when we do the endpoints part, remove this hardcode
		if h.dbService.CheckToken(chatID, 1) {
			user, err := recoverUserAfterDrop(h.dbService, chatID, username, h.config.BaseURL)
			if err != nil {
				log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to restore user from the database. The user has been sent for registration")
				h.handleNewUserRegistration(ctxWithUser, tgb, update)
				return
			}

			h.cache.SetUser(chatID, *user)
			log.Info().Int64("chat_id", chatID).Msg("User successfully restored after drop")
			ctxWithUser = context.WithValue(parentCtx, database.UserCtxKey, user)
			h.handleSendAIModelSelectionKeyboard(ctxWithUser, tgb, update)
			return
		}

		log.Warn().Int64("chat_id", chatID).Msg("User not found in cache or database. Redirecting to registration.")
		h.handleNewUserRegistration(ctxWithUser, tgb, update)

	}
}

// this func ensures that a callback query is processed only once at a time per message
func callbackSingleExecutionMiddleWare(next bot.HandlerFunc) bot.HandlerFunc {
	sf := sync.Map{}
	return func(ctxWithUser context.Context, tgb *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			key := update.CallbackQuery.Message.Message.ID
			if _, loaded := sf.LoadOrStore(key, struct{}{}); loaded {
				return
			}
			defer sf.Delete(key)
			next(ctxWithUser, tgb, update)
		}
	}
}

// Middleware that filters messages in groups without bot mention
func (h *handlers) filterGroupMessagesMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, tgb *bot.Bot, update *models.Update) {
		if update.Message.Chat.ID < 0 {

			if update.Message.Chat.ID < 0 && update.Message.Voice != nil {
				return
			}
			if update.Message.Text != "" && !strings.Contains(update.Message.Text, h.botUsername) {
				return
			}

			if update.Message.Photo != nil && (update.Message.Caption == "" || !strings.Contains(update.Message.Caption, h.botUsername)) {
				return
			}
		}

		next(ctx, tgb, update)
	}
}
