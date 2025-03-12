package handlers

import (
	"context"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/localai/imageRecognition"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// Image recognition handler
func (h *handlers) handleRecognizeImage(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	params := &bot.GetFileParams{FileID: update.Message.Photo[0].FileID}
	msgFailedRecognizeFunc := func() {
		if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: errMsgFailedRecognizeImage}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
	}
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	file, err := tgb.GetFile(ctx, params)
	if err != nil {
		msgFailedRecognizeFunc()
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to get file from Telegram API")
		return
	}

	fileURL := urlTelegramServeFilesConstructor(tgb.Token(), file.FilePath)
	prompt := update.Message.Caption
	if update.Message.Caption == "" {
		prompt = basePromptRecognizeImage
	}

	url := getURL(user.AiSession.BaseURL, h.config.ImageRecognitionEndpoint)
	model := h.config.ImageRecognitionModel
	recognize, err := imageRecognition.ImageRecognitionLAI(url, model, tgb.Token(), fileURL, prompt)
	if err != nil {
		msgFailedRecognizeFunc()
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed recognize image")
		return
	}

	if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: recognize}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}
