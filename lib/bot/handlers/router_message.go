package handlers

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

const (
	status_MainHandlerAfterUserIdentification = "mainHandlerAfterUserIdentification" // old dialogStatus = 6
	status_AIModelSelectionKeyboard           = "AIModelSelectionKeyboard"           // old dialogStatus = 3
)

// Router for text message handlers
func (h *handlers) textMessageRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {

}

// 3 - status_AIModelSelectionKeyboard
func (h *handlers) handleSendAIModelSelectionKeyboard(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	var gptKey string
	user := h.cache.data[chatID]

	if update.Message.Text != "" && user.AiSession.GptKey == "" {
		gptKey = strings.TrimSpace(update.Message.Text)
		h.db_service.InsertToken(chatID, 1, gptKey)
		user.AiSession.GptKey = gptKey
	}
	baseURL := h.ai_endpoint
	aiModelsList, err := h.db_service.GetModelsList(baseURL, gptKey)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving LLM models list")
		h.handleNewUserRegistration(ctx, tgb, update)
		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to LLM models list error")
		return
	}
	//mainHandlerAfterUserIdentification
	user.DialogStatus = 4
	h.cache.data[chatID] = user

	msg := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Choose model",
		ReplyMarkup: CreateAIModelsMarkup(aiModelsList),
	}

	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}
