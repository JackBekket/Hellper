package handlers

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

//todo

const (
	status_AIModelSelectionChoice = "AIModelSelectionChoice" // old dialogStatus - 4
	status_ConnectingToAiWithLang = "connectingToAiWithLang" // old dialogStatus - 5

)

func (h *handlers) callbackRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	// Stub for the user cache structure
	dialogStatus := status_AIModelSelectionChoice
	switch dialogStatus {
	case status_AIModelSelectionChoice:
		h.handleAIModelSelectionCallback(ctx, tgb, update)
	case status_ConnectingToAiWithLang:
		h.handleConnectingToAiWithLangCallback(ctx, tgb, update)

	default: // todo: error msg
	}

}

// 4 - status_AIModelSelectionChoice. Old func name - HandleModelChoose
func (h *handlers) handleAIModelSelectionCallback(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.CallbackQuery.From.ID
	content := update.CallbackQuery.Data

	messageID := update.CallbackQuery.ID
	callbackResponse := &bot.AnswerCallbackQueryParams{
		CallbackQueryID: messageID,
		Text:            "🐈💨",
	}
	_, err := tgb.AnswerCallbackQuery(ctx, callbackResponse)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
		return

	}

	// Use strings.Split to separate the string by "_".
	parts := strings.Split(content, "_")
	if len(parts) < 2 {
		log.Error().Int64("chat_id", chatID).Str("content", content).Caller().Msg("invalid callback data format")
		return
	}
	aiModelName := parts[1]

	user, ok := h.cache.GetUser(chatID)
	if !ok {
		log.Error().Int64("chat_id", chatID).Msg("user not found in cache")
		return
		// todo: Add actions in case the user is not found in the cache
	}

	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msg_Session_model + aiModelName,
	}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	langMsg := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        msg_Choose_lang,
		ReplyMarkup: renderLangInlineKeyboard(),
	}

	_, err = tgb.SendMessage(ctx, langMsg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
	// status_ConnectingToAiWithLang
	user.DialogStatus = 5
	user.AiSession.GptModel = aiModelName
	h.cache.UpdateUser(user)

}
