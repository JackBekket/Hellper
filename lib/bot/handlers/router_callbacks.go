package handlers

import (
	"context"
	"strings"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
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

func (h *handlers) handleConnectingToAiWithLangCallback(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.CallbackQuery.From.ID
	lang := update.CallbackQuery.Data

	log.Info().Int64("chat_id", chatID).Str("lang", lang).Msg("User initiated AI connection via callback")

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

	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msg_Connecting_AI_node,
	}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	// I commented out the line because the context with the value is not used anywhere
	//ctxWithValue := context.WithValue(ctx, "user", user)
	langPromt := getInitialLangPromt(lang)
	log.Info().Int64("chat_id", chatID).Str("language", lang).Str("ai_endpoint", h.baseURL).
		Msg("Starting AI conversation")

	go h.handleStartAiConversationWithLang(ctx, tgb, chatID, langPromt)
}

// old name func - SetupSequenceWithKey
func (h *handlers) handleStartAiConversationWithLang(ctx context.Context, tgb *bot.Bot, chatID int64, langPromt string) {
	user, ok := h.cache.GetUser(chatID)
	if !ok {
		log.Error().Int64("chat_id", chatID).Msg("user not found in cache")
		return
		// todo: Add actions in case the user is not found in the cache
	}

	probe, response, err := langchain.RunNewAgent(user.AiSession.GptKey, user.AiSession.GptModel, h.baseURL, langPromt)
	if err != nil {
		msg := &bot.SendMessageParams{ChatID: chatID, Text: msg_AI_client_failure}
		_, err = tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}

		videoMsg, err := getMsgRandomErrorVideo(chatID)
		_, err = tgb.SendVideo(ctx, videoMsg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending video message")
		}
		log.Warn().Int64("chat_id", chatID).Str("username", user.Username).Msg("The user was removed from the cache due to an authentication issue.")
		h.cache.DeleteUser(chatID)
		return
	}

	msg := &bot.SendMessageParams{ChatID: chatID, Text: response}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user.DialogStatus = 6
	user.AiSession.DialogThread = *probe

	// TODO: Replace with a thread-safe one
	usage := database.GetSessionUsage(user.ID)
	user.AiSession.Usage = usage

	h.cache.UpdateUser(user)
	log.Info().Int64("chat_id", chatID).Str("username", user.Username).Str("ai_endpoint", h.baseURL).
		Msg("AI conversation completed successfully")
}

// Returns the initial prompt with the selected language
func getInitialLangPromt(lang string) string {
	switch lang {
	case "English":
		return initialPromt_Lang_EN
	case "Russian":
		return initialPromt_Lang_RU
	default:
		return initialPromt_Lang_EN
	}
}
