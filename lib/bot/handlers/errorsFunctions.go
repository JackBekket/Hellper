package handlers

import (
	"context"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// The function prepares a message with a random video
// old func name - errorMessage.
func getErrorMsgWithRandomVideo(ctx context.Context, tgb *bot.Bot, chatID int64) {
	// Send helper video error
	// Get a list of all files in the media directory
	filePath := "./media"
	files, err := func() ([]fs.FileInfo, error) {
		f, err := os.Open(filePath)
		if err != nil {
			log.Error().Err(err).Msg("failed to open media directory")
			return nil, err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to read media directory")
			return nil, err
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name() < list[j].Name()
		})
		return list, nil
	}()
	if err != nil {
		return
	}

	// Select a random file
	randomFile := files[rand.Intn(len(files))]
	videoFile, err := os.Open(filepath.Join(filePath, randomFile.Name()))
	if err != nil {
		log.Error().Err(err).Msg("failed to open selected video file")
	}
	defer videoFile.Close()

	if _, err := tgb.SendVideo(ctx, &bot.SendVideoParams{
		ChatID:  chatID,
		Caption: msgAIclientFailure,
		Video: &models.InputFileUpload{
			Filename: randomFile.Name(),
			Data:     videoFile,
		},
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending video message")
	}

}
