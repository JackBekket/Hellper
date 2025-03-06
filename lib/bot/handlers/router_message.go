package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/JackBekket/hellper/lib/agent"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
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
	case statusAIModelSelectionKeyboardForNewUser:
		h.handleSendAIModelSelectionKeyboardForNewUser(ctx, tgb, update)
	case statusAIModelSelectionKeyboardForExistUser:
		h.handleSendAIModelSelectionKeyboardForNewUser(ctx, tgb, update)
	case statusStartDialogSequence:
		go h.handleStartDialogSequence(ctx, tgb, update)
	default: // todo: error msg
	}

}

// 3 - status_AIModelSelectionKeyboard
func (h *handlers) handleSendAIModelSelectionKeyboardForNewUser(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	var localAIToken string
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	if update.Message.Text == "" && user.AiSession.LocalAIToken == "" {
		msg := &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Please input your local-ai API token",
		}

		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
		return
	}

	// сделать проверку валидности токена
	localAIToken = strings.TrimSpace(update.Message.Text)
	h.dbService.InsertToken(chatID, 1, localAIToken)
	user.AiSession.LocalAIToken = localAIToken

	url := getURL(h.config.BaseURL, h.config.ModelsListEndpoint)
	aiModelsList, err := h.dbService.GetModelsList(url, localAIToken)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving LLM models list")
		h.handleNewUserRegistration(ctx, tgb, update)
		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to LLM models list error")
		return
	}
	user.DialogStatus = statusAIModelSelectionChoice
	h.cache.UpdateUser(user)

	msg := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Choose model",
		ReplyMarkup: renderAIModelsInlineKeyboard(aiModelsList),
	}

	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

// 3 - status_AIModelSelectionKeyboard
func (h *handlers) handleSendAIModelSelectionKeyboardForExistUser(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	// сделать проверку валидности токена
	localAIToken := user.AiSession.LocalAIToken
	fmt.Println("sendkbexist", user)
	if localAIToken == "" {
		msg := &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Your token is invalid. Please enter a new token.",
		}

		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}

		user.DialogStatus = statusAIModelSelectionChoice
		h.cache.UpdateUser(user)
		return
	}

	url := getURL(h.config.BaseURL, h.config.ModelsListEndpoint)
	aiModelsList, err := h.dbService.GetModelsList(url, localAIToken)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving LLM models list")
		h.handleNewUserRegistration(ctx, tgb, update)
		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to LLM models list error")
		return
	}
	user.DialogStatus = statusAIModelSelectionChoice
	h.cache.UpdateUser(user)

	msg := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Choose model",
		ReplyMarkup: renderAIModelsInlineKeyboard(aiModelsList),
	}

	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

// Dialog_Status 6 -> 6 (loop) - status_StartDialogSequence
func (h *handlers) handleStartDialogSequence(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	model := user.AiSession.GptModel
	prompt := update.Message.Text
	thread := user.AiSession.DialogThread
	log.Info().Str("gpt_model", model).Str("prompt", prompt).Msg("processing GPT request")

	post_session, resp, err := langchain.ContinueAgent(user.AiSession.LocalAIToken, model, h.config.BaseURL, prompt, &thread)
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

	msg := &bot.SendMessageParams{ChatID: chatID, Text: resp, ParseMode: "MARKDOWN"}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user.DialogStatus = statusStartDialogSequence

	usage := database.GetSessionUsage(chatID)
	// if !ok {
	// 	log.Error().Int64("chat_id", chatID).Caller().Msg("usage not found in cache")
	// 	return
	// 	// todo: Add actions in case the user is not found in the cache
	// }

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

	if err := updateUserHistoryInDB(h.dbService, chatID, threadID, model, humanType[0], last_msg); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Str("model", model).Caller().Msg("failed to update user history")
		return
	}

}

func updateUserHistoryInDB(db *database.Service, chatID, threadID int64, model string, humanType, last_msg llms.MessageContent) error {
	if err := db.UpdateHistory(chatID, 1, chatID, threadID, model, humanType); err != nil {
		return err
	} //endpointID is hardcoded and why chatID is threadID?
	if err := db.UpdateHistory(chatID, 1, chatID, threadID, model, last_msg); err != nil {
		return err
	}
	return nil
}
