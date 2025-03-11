package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

func (h *handlers) callbackRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	callback := update.CallbackQuery
	chatID := callback.From.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	switch user.DialogStatus {
	case statusAIModelSelectionChoiceCallback:
		h.handleAIModelSelectionCallback(ctx, tgb, callback)
	//case statusConnectingToAiWithLangCallback:
	// Can also accept a text message with the first prompt
	//h.handleConnectingToAiWithLangCallback(ctx, tgb, update)
	case statusAuthMethodCallback:
		h.handleAuthMethodCallback(ctx, tgb, callback)
	case statusLocalAIProviderCallback:
		h.handleLocalAIProviderCallback(ctx, tgb, callback)
	default:
		if _, err := tgb.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: callback.ID,
			Text:            "🐈💨",
		}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
			return
		}
		if user.AiSession.AIToken == "" {
			h.handleNewUserRegistration(ctx, tgb, update)
		}

	}

}

func (h *handlers) handleAuthMethodCallback(ctx context.Context, tgb *bot.Bot, callback *models.CallbackQuery) {
	chatID := callback.From.ID
	data := callback.Data

	if _, err := tgb.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
		Text:            "🐈💨",
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
		return
	}
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	authMethod, err := strconv.Atoi(data)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("сonversion error")
		return
	}
	aiProviders, err := h.dbService.GetAIProvidersName(int64(authMethod))
	if err != nil || len(aiProviders) == 0 {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving providers name from the database")
		return
	}
	if authMethod == database.AuthMethodOpenAI {
		// OpenAI has a single base URL for the API
		h.handleOpenAIAuthMethod(ctx, tgb, chatID, aiProviders[0])
		return
	}
	//Currently, only localAI has multiple base URLs
	if _, err := tgb.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      chatID,
		MessageID:   callback.Message.Message.ID,
		ReplyMarkup: renderLocalAIProvidersInlineKeyboard(aiProviders),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error editing replyMarkup")
		return
	}

	user.AiSession.AuthMethod = database.AuthMethodLocalAI
	user.DialogStatus = statusLocalAIProviderCallback
	h.cache.UpdateUser(user)

}

func (h *handlers) handleOpenAIAuthMethod(ctx context.Context, tgb *bot.Bot, chatID int64, providerName string) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf(msgEnterAPIToken, providerName),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
	provider, err := h.dbService.GetAIProvider(providerName)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving provider from the database")
		return
	}
	user.DialogStatus = statusAPIToken
	user.AiSession.ProviderID = provider.ID
	user.AiSession.BaseURL = provider.BaseURL
	user.AiSession.ProviderName = provider.Name
	user.AiSession.AuthMethod = database.AuthMethodOpenAI

	h.cache.UpdateUser(user)
}

func (h *handlers) handleLocalAIProviderCallback(ctx context.Context, tgb *bot.Bot, callback *models.CallbackQuery) {
	chatID := callback.From.ID
	providerName := callback.Data

	if _, err := tgb.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
		Text:            "🐈💨",
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
		return
	}

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf(msgEnterAPIToken, providerName),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	provider, err := h.dbService.GetAIProvider(providerName)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving provider from the database")
		return
	}

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	user.DialogStatus = statusAPIToken
	user.AiSession.ProviderID = provider.ID
	user.AiSession.BaseURL = provider.BaseURL
	user.AiSession.ProviderName = provider.Name
	h.cache.UpdateUser(user)

}

// 4 - status_AIModelSelectionChoice. Old func name - HandleModelChoose
func (h *handlers) handleAIModelSelectionCallback(ctx context.Context, tgb *bot.Bot, callback *models.CallbackQuery) {
	chatID := callback.From.ID
	content := callback.Data

	if _, err := tgb.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
		Text:            "🐈💨",
	}); err != nil {
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
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf(msgSessionModelFormat, aiModelName),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msgFirstPrompt,
		//ReplyMarkup: renderLangInlineKeyboard(),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	// status_ConnectingToAiWithLang
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.DialogStatus = statusConnectingToAIWithFirstPrompt
	user.AiSession.GptModel = aiModelName
	h.cache.UpdateUser(user)

}

// func (h *handlers) handleConnectingToAiWithLangCallback(ctx context.Context, tgb *bot.Bot, update *models.Update) {
// 	var lang, langPrompt string
// 	var chatID int64
// 	if update.Message == nil {
// 		chatID = update.CallbackQuery.From.ID
// 		lang = update.CallbackQuery.Data
// 		log.Info().Int64("chat_id", chatID).Str("lang", lang).Msg("User initiated AI connection via callback")
// 		if _, err := tgb.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
// 			CallbackQueryID: update.CallbackQuery.ID,
// 			Text:            "🐈💨",
// 		}); err != nil {
// 			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error answering callback query")
// 			return
// 		}
// 		langPrompt = getInitialLangPrompt(lang)
// 	} else {
// 		langPrompt = update.Message.Text
// 	}

// 	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
// 		ChatID: chatID,
// 		Text:   msgConnectingAINode,
// 	}); err != nil {
// 		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
// 		return
// 	}

// 	go h.handleStartAiConversationWithLang(ctx, tgb, chatID, langPrompt)
// }

// // old name func - SetupSequenceWithKey
// func (h *handlers) handleStartAiConversationWithLang(ctx context.Context, tgb *bot.Bot, chatID int64, langPrompt string) {
// 	user, ok := ctx.Value(database.UserCtxKey).(database.User)
// 	if !ok {
// 		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
// 		return
// 	}
// 	log.Info().Int64("chat_id", chatID).Str("endpoint", user.AiSession.BaseURL).
// 		Msg("Starting AI conversation")

// 	model := user.AiSession.GptModel
// 	probe, response, err := langchain.RunNewAgent(user.AiSession.AIToken, model, user.AiSession.BaseURL, langPrompt)
// 	if err != nil {
// 		videoMsg, err := getErrorMsgWithRandomVideo(chatID)
// 		if err != nil {
// 			log.Error().Err(err).Caller().Msg("error generating video message")
// 			return
// 		}
// 		if _, err := tgb.SendVideo(ctx, videoMsg); err != nil {
// 			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending video message")
// 		}
// 		log.Warn().Int64("chat_id", chatID).Str("username", user.Username).Msg("The user was removed from the cache due to an authentication issue.")
// 		h.cache.DeleteUser(chatID)
// 		return
// 	}

// 	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: response}); err != nil {
// 		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
// 		return
// 	}

// 	user.DialogStatus = statusStartDialogSequence
// 	user.AiSession.DialogThread = *probe

// 	// TODO: Replace with a thread-safe one
// 	usage := database.GetSessionUsage(user.ID)
// 	user.AiSession.Usage = usage

// 	h.dbService.CreateAISession(chatID, user.AiSession.GptModel, user.AiSession.ProviderID)
// 	h.cache.UpdateUser(user)
// 	log.Info().Int64("chat_id", chatID).Str("username", user.Username).Str("BaseURL", user.AiSession.BaseURL).
// 		Msg("AI conversation completed successfully")
// }

// // Returns the initial prompt with the selected language
// func getInitialLangPrompt(lang string) string {
// 	switch lang {
// 	case langEnglish:
// 		return basePromptLangEN
// 	case langRussian:
// 		return basePromptLangRU
// 	default:
// 		return basePromptLangEN
// 	}
// }
