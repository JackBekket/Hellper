package handlers

import (
	"context"

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
		msg := &bot.SendMessageParams{ChatID: chatID, Text: errMsgFailedRecognizeImage}
		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
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

	url := getURL(h.config.BaseURL, fileURL)
	model := h.config.ImageRecognitionModel
	recognize, err := imageRecognition.ImageRecognitionLAI(url, model, tgb.Token(), fileURL, prompt)
	if err != nil {
		msgFailedRecognizeFunc()
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed recognize image")
		return
	}

	msg := &bot.SendMessageParams{ChatID: chatID, Text: recognize}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}
