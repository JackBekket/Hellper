package handlers

import (
	"context"
	"sync"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// The global middleware checks for the user's presence in the cache and PostgreSQL.
// If the user is absent, it initiates the registration or data recovery process
func (h *handlers) IdentifyUserMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, tgb *bot.Bot, update *models.Update) {
		var chatID int64
		if update.Message != nil {
			chatID = update.Message.From.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.From.ID
		} else {
			next(ctx, tgb, update)
			return
		}

		_, found := h.cache.data[chatID]
		if found {
			next(ctx, tgb, update)
			return
		}

		if h.db_service.CheckSession(chatID) {
			user, err := restoreUserSessionFromDB(h.db_service, chatID, update.Message.From.Username)
			if err != nil {
				log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to restore user from the database. The user has been sent for registration")
				h.handleNewUser(ctx, tgb, update)
				return
			}

			database.AddUser(*user) // add user from persistent db into cache
			next(ctx, tgb, update)
			return
		}

		//TODO: when we do the endpoints part, remove this hardcode
		if h.db_service.CheckToken(chatID, 1) {
			h.handleRecoverUserAfterDrop(ctx, tgb, update)
			return
		}

		// todo
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
