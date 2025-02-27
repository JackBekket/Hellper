package handlers

import (
	"context"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// old func name: AddNewUserToMap.
// Creates a session for a new user. Sends a welcome message.
func (h *handlers) handleNewUser(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	user := database.User{
		ID:           chatID,
		Username:     update.Message.From.Username,
		DialogStatus: 3,
		Admin:        false,
		AiSession: database.AiSession{
			Base_url: h.ai_endpoint,
		},
	}

	h.cache.data[chatID] = user

	log.Info().
		Int64("chat_id", user.ID).
		Str("username", user.Username).
		Msg("new user")

	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msgTemplates["hello"],
	}

	_, err := tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

func (h *handlers) handleRecoverUserAfterDrop(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	username := update.Message.From.Username

	log.Info().
		Int64("chat_id", chatID).
		Str("username", username).
		Msg("Restoring a registered user")

	user := database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: 4,
		Admin:        false,
		AiSession: database.AiSession{
			Base_url: h.ai_endpoint,
		},
	}

	apiKey, err := h.db_service.GetToken(chatID, 1)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Msg("Ошибка получения API-ключа")
		return
	}
	user.AiSession.GptKey = apiKey

	h.cache.data[chatID] = user

	h.RenderModelsForRegisteredUser(ctx, tgb, update) //todo
}
