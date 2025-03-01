package handlers

import (
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// The function prepares a message with a random video
// old func name - errorMessage.
func getErrorMsgWithRandomVideo(chatID int64) (*bot.SendVideoParams, error) {
	// Send helper video error
	// Get a list of all files in the media directory
	files, err := func() ([]fs.FileInfo, error) {
		f, err := os.Open("../../media/")
		if err != nil {
			return nil, err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, err
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name() < list[j].Name()
		})
		return list, nil
	}()
	if err != nil {
		return &bot.SendVideoParams{}, err
	}

	// Select a random file
	//rand.Seed(time.Now().UnixNano())
	randomFile := files[rand.Intn(len(files))]

	// Open the video file
	videoFile, err := os.Open(filepath.Join("../../media/", randomFile.Name()))
	if err != nil {
		return &bot.SendVideoParams{}, err
	}
	defer videoFile.Close()

	return &bot.SendVideoParams{
		ChatID: chatID,
		Video: &models.InputFileUpload{
			Filename: randomFile.Name(),
			Data:     videoFile,
		},
	}, nil

}
