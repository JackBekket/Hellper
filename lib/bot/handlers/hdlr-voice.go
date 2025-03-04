package handlers

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/localai"
	stt "github.com/JackBekket/hellper/lib/localai/audioRecognition"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// The handler transcribes the voice message and sends the result to the user
func (h *handlers) handleVoiceTranscriber(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	params := &bot.GetFileParams{FileID: update.Message.Voice.FileID}
	msgFailedVoiceFunc := func() {
		msg := &bot.SendMessageParams{ChatID: chatID, Text: errMsgFailedTrascribeVoice}
		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
	}
	file, err := tgb.GetFile(ctx, params)
	if err != nil {
		msgFailedVoiceFunc()
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("failed to get file from Telegram API")
		return
	}

	fileURL := urlTelegramServeFilesConstructor(tgb.Token(), file.FilePath)

	localFilePath := filepath.Join("tmp", "audio", update.Message.Voice.FileID+".ogg")
	err = downloadFile(fileURL, localFilePath)
	if err != nil {
		log.Error().Err(err).Str("file_url", fileURL).Str("file_path", localFilePath).Caller().Msg("failed to download file")
		msgFailedVoiceFunc()
		return
	}

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Msg("user not found in context")
		return
	}

	url, model := stt.GetEnvsForSST()
	transcription, err := localai.TranscribeWhisper(url, model, localFilePath, user.AiSession.LocalAIToken)
	if err != nil {
		msgFailedVoiceFunc()
		log.Error().Err(err).Str("url", url).Str("model", model).Str("file_path", localFilePath).Caller().Msg("failed to transcribe audio")
		return
	}

	msg := &bot.SendMessageParams{ChatID: chatID, Text: transcription}
	_, err = tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	if err := os.Remove(localFilePath); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error deleting image")
		return
	}
}

func downloadFile(url, localFilePath string) error {
	// Create the file
	out, err := os.Create(localFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to the file
	_, err = io.Copy(out, resp.Body)
	return err
}
