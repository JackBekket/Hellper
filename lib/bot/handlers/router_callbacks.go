package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

func (h *handlers) callbackRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.CallbackQuery.From.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	fmt.Println("ПОЛУЧЕН КОЛБЕК")

	switch user.DialogStatus {
	case statusAIModelSelectionChoice:
		h.handleAIModelSelectionCallback(ctx, tgb, update)
	case statusConnectingToAiWithLang:
		h.handleConnectingToAiWithLangCallback(ctx, tgb, update)
	default: // todo: error msg

		// проверка языка

		messageID := update.CallbackQuery.ID
		callbackResponse := &bot.AnswerCallbackQueryParams{
			CallbackQueryID: messageID,
			Text:            "Неизвестный Хендлер",
		}
		_, err := tgb.AnswerCallbackQuery(ctx, callbackResponse)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
			return

		}

		if user.AiSession.LocalAIToken == "" {
			h.handleNewUserRegistration(ctx, tgb, update)
		}

		h.handleConnectingToAiWithLangCallback(ctx, tgb, update)

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
	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msgSessionModel + aiModelName,
	}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	langMsg := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        msgChooseLang,
		ReplyMarkup: renderLangInlineKeyboard(),
	}

	_, err = tgb.SendMessage(ctx, langMsg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	// status_ConnectingToAiWithLang
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.DialogStatus = statusConnectingToAiWithLang
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
		Text:   msgConnectingAINode,
	}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	// I commented out the line because the context with the value is not used anywhere
	//ctxWithValue := context.WithValue(ctx, "user", user)
	langPrompt := getInitialLangPrompt(lang)
	log.Info().Int64("chat_id", chatID).Str("language", lang).Str("endpoint", h.config.BaseURL).
		Msg("Starting AI conversation")

	go h.handleStartAiConversationWithLang(ctx, tgb, chatID, langPrompt)
}

// old name func - SetupSequenceWithKey
func (h *handlers) handleStartAiConversationWithLang(ctx context.Context, tgb *bot.Bot, chatID int64, langPrompt string) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	model := user.AiSession.GptModel
	probe, response, err := langchain.RunNewAgent(user.AiSession.LocalAIToken, model, h.config.BaseURL, langPrompt)
	if err != nil {
		videoMsg, err := getErrorMsgWithRandomVideo(chatID)
		if err != nil {
			log.Error().Err(err).Caller().Msg("")
			return
		}
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

	user.DialogStatus = statusStartDialogSequence
	user.AiSession.DialogThread = *probe

	// TODO: Replace with a thread-safe one
	usage := database.GetSessionUsage(user.ID)
	user.AiSession.Usage = usage

	h.dbService.CreateLSession(chatID, model, 1)

	h.cache.UpdateUser(user)
	log.Info().Int64("chat_id", chatID).Str("username", user.Username).Str("BaseURL", h.config.BaseURL).
		Msg("AI conversation completed successfully")
}

// Returns the initial prompt with the selected language
func getInitialLangPrompt(lang string) string {
	switch lang {
	case "English":
		return basePromptLangEN
	case "Russian":
		return basePromptLangRU
	default:
		return basePromptLangEN
	}
}
