package handlers

import (
	"context"
	"strings"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// The global middleware checks for the user's presence in the cache and PostgreSQL.
// If the user is absent, it initiates the registration or data recovery process
func (h *handlers) IdentifyUserMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, tgb *bot.Bot, update *models.Update) {
		var chatID int64
		var username string

		switch {
		case update.Message != nil:
			chatID = update.Message.From.ID
			username = update.Message.From.Username
		case update.CallbackQuery != nil:
			chatID = update.CallbackQuery.From.ID
			username = update.CallbackQuery.From.Username
		default:
			next(ctx, tgb, update)
			return
		}

		_, found := h.cache.GetUser(chatID)
		if found {
			next(ctx, tgb, update)
			return
		}

		if h.db_service.CheckSession(chatID) {
			user, err := restoreUserSessionFromDB(h.db_service, chatID, update.Message.From.Username)
			if err != nil {
				log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to restore user from the database. The user has been sent for registration")
				h.handleNewUserRegistration(ctx, tgb, update)
				return
			}

			h.cache.SetUser(chatID, *user) // add user from persistent db into cache
			log.Info().Int64("chat_id", chatID).Msg("User session successfully restored from the database. User added to the cache.")
			next(ctx, tgb, update)
			return
		}

		//TODO: when we do the endpoints part, remove this hardcode
		if h.db_service.CheckToken(chatID, 1) {
			user, err := recoverUserAfterDrop(h.db_service, chatID, username, h.config.AI_endpoint)
			if err != nil {
				log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to restore user from the database. The user has been sent for registration")
				h.handleNewUserRegistration(ctx, tgb, update)
				return
			}

			h.cache.SetUser(chatID, *user)
			log.Info().Int64("chat_id", chatID).Msg("User successfully restored after drop")
			h.handleSendAIModelSelectionKeyboard(ctx, tgb, update)
			return
		}

		log.Warn().Int64("chat_id", chatID).Msg("User not found in cache or database. Redirecting to registration.")
		h.handleNewUserRegistration(ctx, tgb, update)

	}
}

// this func ensures that a callback query is processed only once at a time per message
func callbackSingleExecutionMiddleWare(next bot.HandlerFunc) bot.HandlerFunc {
	sf := sync.Map{}
	return func(ctx context.Context, tgb *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			key := update.CallbackQuery.Message.Message.ID
			if _, loaded := sf.LoadOrStore(key, struct{}{}); loaded {
				return
			}
			defer sf.Delete(key)
			next(ctx, tgb, update)
		}
	}
}

// Middleware that parses the msg into a command and arguments
// to pass them to the next handler via context
func cmdHandlerMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil || update.Message.Text == "" {
			return
		}
		command, arg := extractCommandAndArg(update.Message.Text)
		if command == "" {
			return
		}

		log.Info().Str("command", command).Str("arg", arg).Int64("chat_id", update.Message.Chat.ID).Msg("processing command")

		ctx = context.WithValue(ctx, context_BotCommand, command)
		ctx = context.WithValue(ctx, context_CommandArg, arg)

		next(ctx, b, update)
	}
}

// Function to extract the command and argument.
// Also removes the bot's name if the message was sent in a group chat.
// At the moment, only one argument is allowed
func extractCommandAndArg(msg string) (string, string) {
	msg = strings.TrimSpace(msg)

	if len(msg) == 0 || msg[0] != '/' {
		return "", ""
	}

	parts := strings.Fields(msg)
	command := strings.Split(parts[0], "@")[0]
	arg := strings.TrimSpace(strings.Join(parts[1:], " "))

	return command, arg
}
