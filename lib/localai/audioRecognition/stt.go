package stt

import (
	"os"
)

// // Deprecated, move to handlers
// func HandleVoiceMessage(updateMessage *tgbotapi.Message, bot tgbotapi.BotAPI) (string, error) {

// 	fileID := updateMessage.Voice.FileID
// 	filebaseURL, err := GetFileURL(fileID, bot)
// 	if err != nil {
// 		log.Println("Error getting file URL:", err)
// 		return "", err
// 	}
// 	localFilePath := filepath.Join("tmp", "audio", updateMessage.Voice.FileID+".ogg")
// 	err = DownloadFile(fileURL, localFilePath)
// 	if err != nil {
// 		log.Println("Error downloading the file:", err)
// 		return "", err
// 	}

// 	return localFilePath, nil
// }
// // Deprecated, move to handlers
// func GetFileURL(fileID string, bot tgbotapi.BotAPI) (string, error) {

// 	fileConfig := tgbotapi.FileConfig{FileID: fileID}
// 	file, err := bot.GetFile(fileConfig)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Telegram serves files via the URL like this:
// 	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)
// 	return fileURL, nil
// }
// // Deprecated, move to handlers
// func DownloadFile(url, localFilePath string) error {
// 	// Create the file
// 	out, err := os.Create(localFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Write the body to the file
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }

// url, model

// REFACTOR
func GetEnvsForSST() (string, string) {
	baseURL := os.Getenv("AI_BASEURL")
	endpoint := os.Getenv("VOICE_RECOGNITION_ENDPOINT")
	if endpoint == "" {
		endpoint = "/v1/audio/transcriptions"
	}
	baseURL += endpoint

	model := os.Getenv("VOICE_RECOGNITION_MODEL")
	if model == "" {
		model = "whisper-1"
	}
	return baseURL, model
}
