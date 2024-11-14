package imageRecognition

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getEnvsForImgRec() (string, string, string) {
	url := os.Getenv("AI_ENDPOINT")
	URLSuffix := os.Getenv("IMAGE_RECOGNITION_SUFFIX")
	if URLSuffix == "" {
		URLSuffix = "/v1/chat/completions"
	}
	url += URLSuffix

	model := os.Getenv("IMAGE_RECOGNITION_MODEL")
	if model == "" {
		model = "bunny-llama-3-8b-v"
	}

	token := os.Getenv("OPENAI_API_KEY")
	return url, model, token

}

func RecognizeImage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) (string, error) {

	imgLink, err := handleImageMessage(bot, msg)
	if err != nil {
		return "", err
	}
	endpoint, model, token := getEnvsForImgRec()
	prompt := "What's in the image?"

	if msg.Caption != "" {
		prompt = msg.Caption
	}
	response, err := imageRecognitionLAI(endpoint, model, token, imgLink, prompt)
	if err != nil {
		return "", err
	}
	return response, nil

}

func handleImageMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) (string, error) {

	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: msg.Photo[0].FileID})
	if err != nil {
		return "", fmt.Errorf("could not get file info: %v", err)
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

	return fileURL, nil
}

func imageRecognitionLAI(url string, model string, token string, imgLink string, prompt string) (string, error) {

	client := &http.Client{}

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{"type": "text", "text": prompt},
					{"type": "image_url", "image_url": map[string]string{
						"url": imgLink,
					}},
				},
			},
		},
		"temperature": 0.3,
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := getMessageContent(resp)
	if err != nil {
		return "", err
	}
	return response, nil
}

func getMessageContent(resp *http.Response) (string, error) {
	type ResponseBody struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseBody ResponseBody
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return "", err
	}

	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no message content found")
}
