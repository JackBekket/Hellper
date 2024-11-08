package stt

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleVoiceMessage(updateMessage *tgbotapi.Message, bot tgbotapi.BotAPI) (string, error) {

	fileID := updateMessage.Voice.FileID
	fileURL, err := GetFileURL(fileID, bot)
	if err != nil {
		log.Println("Error getting file URL:", err)
		return "", err
	}
	localFilePath := filepath.Join("tmp", "audio", updateMessage.Voice.FileID+".ogg")
	err = DownloadFile(fileURL, localFilePath)
	if err != nil {
		log.Println("Error downloading the file:", err)
		return "", err
	}

	return localFilePath, nil
}

func GetFileURL(fileID string, bot tgbotapi.BotAPI) (string, error) {

	fileConfig := tgbotapi.FileConfig{FileID: fileID}
	file, err := bot.GetFile(fileConfig)
	if err != nil {
		return "", err
	}

	// Telegram serves files via the URL like this:
	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)
	return fileURL, nil
}

func DownloadFile(url, localFilePath string) error {
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

// url, model
func GetEnvsForSST() (string, string) {
	url := os.Getenv("AI_ENDPOINT")
	URLSuffix := os.Getenv("VOICE_RECOGNITION_SUFFIX")
	if URLSuffix == "" {
		URLSuffix = "/v1/audio/transcriptions"
	}
	url += URLSuffix

	model := os.Getenv("VOICE_RECOGNITION_MODEL")
	if model == "" {
		model = "whisper-1"
	}
	return url, model
}
