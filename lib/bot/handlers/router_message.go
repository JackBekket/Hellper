package handlers

import (
	"context"
	"strings"

	"github.com/JackBekket/hellper/lib/agent"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	"github.com/JackBekket/hellper/lib/localai"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/llms"
)

// Router for text message handlers
func (h *handlers) textMessageRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	switch user.DialogStatus {
	case statusAIModelSelectionKeyboard:
		h.handleSendAIModelIKB(ctx, tgb, update)
	case statusAPIToken:
		h.handleAPIToken(ctx, tgb, update)
	case statusStartDialogSequence:
		go h.handleStartDialogSequence(ctx, tgb, update)
	default: // todo: error msg
	}
}

func (h *handlers) handleAPIToken(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	localAIToken := strings.TrimSpace(update.Message.Text)
	url := getURL(user.AiSession.BaseURL, h.config.ModelsListEndpoint)
	aiModelsList, err := localai.GetModelsList(url, localAIToken)
	if err != nil {
		if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Your token is invalid. Please enter a new token.",
		}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
		log.Warn().Err(err).Int64("chat_id", chatID).Str("url", url).Caller().Msg("error retrieving LLM models list")
		return
	}
	if err := h.dbService.InsertToken(chatID, user.AiSession.AuthMethod, localAIToken); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error inserting token")
		return
	}
	user.AiSession.AIToken = localAIToken

	user.DialogStatus = statusAIModelSelectionChoiceCallback
	h.cache.UpdateUser(user)

	if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Choose model",
		ReplyMarkup: renderAIModelsInlineKeyboard(aiModelsList),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

func (h *handlers) handleSendAIModelIKB(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	localAIToken, url := user.AiSession.AIToken, getURL(user.AiSession.BaseURL, h.config.ModelsListEndpoint)
	aiModelsList, err := localai.GetModelsList(url, localAIToken)
	if err != nil {
		if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Your token is invalid. Please enter a new token.",
		}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
		log.Warn().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving LLM models list")
		return
	}

	if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        msgChooseModel,
		ReplyMarkup: renderAIModelsInlineKeyboard(aiModelsList),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}

	user.DialogStatus = statusAIModelSelectionChoiceCallback
	h.cache.UpdateUser(user)
}

// Dialog_Status 6 -> 6 (loop) - status_StartDialogSequence
func (h *handlers) handleStartDialogSequence(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	// Handle the case where the connection is successful, but the user has exhausted their quota for using the model.
	// error in node agent:
	// API returned unexpected status code: 429: You exceeded your current quota, please check your plan and billing details.
	// For more information on this error, read the docs: https://platform.openai.com/docs/guides/error-codes/api-errors.
	model, prompt := user.AiSession.GptModel, update.Message.Text
	thread := user.AiSession.DialogThread
	log.Info().Str("gpt_model", model).Str("prompt", prompt).Msg("processing GPT request")
	post_session, resp, err := langchain.ContinueAgent(user.AiSession.AIToken, model, user.AiSession.BaseURL, prompt, &thread)
	if err != nil {
		videoMsg, err := getErrorMsgWithRandomVideo(chatID)
		if err != nil {
			log.Error().Err(err).Caller().Msg("error generating video message")
			return
		}
		if _, err := tgb.SendVideo(ctx, videoMsg); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending video message")
		}
		log.Warn().Int64("chat_id", chatID).Str("username", user.Username).Msg("The user was removed from the cache due to an authentication issue.")
		h.cache.DeleteUser(chatID)
		return
	}

	if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: resp, ParseMode: "MARKDOWN"}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user.DialogStatus = statusStartDialogSequence
	usage := database.GetSessionUsage(chatID)
	user.AiSession.Usage = usage
	user.AiSession.DialogThread = *post_session
	h.cache.UpdateUser(user)

	totalTurns := len(thread.ConversationBuffer)
	log.Info().Int("total_turns", totalTurns).Msg("conversation turns counted")

	// here we save user conversation to the db?
	// Update the user in the database here. Yes, everything that happens with the "user" should be immediately visible.
	buffer := post_session.ConversationBuffer
	last_msg := buffer[len(buffer)-1]
	humanType := agent.CreateMessageContentHuman(prompt)
	threadID := chatID

	if err := updateUserHistoryInDB(h.dbService, chatID, user.AiSession.ProviderID, threadID, model, humanType[0], last_msg); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Str("model", model).Caller().Msg("failed to update user history")
		return
	}

}

func updateUserHistoryInDB(db *database.Service, chatID, providerID, threadID int64, model string, humanType, last_msg llms.MessageContent) error {
	if err := db.UpdateHistory(chatID, providerID, chatID, threadID, model, humanType); err != nil {
		return err
	} //endpointID is hardcoded and why chatID is threadID?
	if err := db.UpdateHistory(chatID, providerID, chatID, threadID, model, last_msg); err != nil {
		return err
	}
	return nil
}
