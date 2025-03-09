package imageRecognition

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Deprecated, move to handlers

// func getEnvsForImgRec() (string, string, string) {
// 	baseURL := os.Getenv("AI_ENDPOINT")
// 	endpoint := os.Getenv("IMAGE_RECOGNITION_ENDPOINT")
// 	if endpoint == "" {
// 		endpoint = "/v1/chat/completions"
// 	}
// 	baseURL += endpoint

// 	model := os.Getenv("IMAGE_RECOGNITION_MODEL")
// 	if model == "" {
// 		model = "bunny-llama-3-8b-v"
// 	}

// 	token := os.Getenv("OPENAI_API_KEY")
// 	return baseURL, model, token

// }

// func RecognizeImage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) (string, error) {

// 	imgLink, err := handleImageMessage(bot, msg)
// 	if err != nil {
// 		return "", err
// 	}
// 	endpoint, model, token := getEnvsForImgRec()
// 	prompt := "What's in the image?"

// 	if msg.Caption != "" {
// 		prompt = msg.Caption
// 	}
// 	response, err := ImageRecognitionLAI(endpoint, model, token, imgLink, prompt)
// 	if err != nil {
// 		return "", err
// 	}
// 	return response, nil

// }

// Deprecated, move to handlers
// func handleImageMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) (string, error) {

// 	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: msg.Photo[0].FileID})
// 	if err != nil {
// 		return "", fmt.Errorf("could not get file info: %v", err)
// 	}

// 	filebaseURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

// 	return filebaseURL, nil
// }

func ImageRecognitionLAI(url, model, token, imgLink, prompt string) (string, error) {
	client := &http.Client{}

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{"type": "text", "text": prompt},
					{"type": "image_url", "image_url": map[string]string{"url": imgLink}},
				},
			},
		},
		"temperature": 0.3,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to marshal payload")
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to create new request")
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to send request")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		log.Error().Int("status_code", resp.StatusCode).Msgf("unexpected status code: %s", string(errorBody))
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(errorBody))
	}

	response, err := getMessageContent(resp)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to get message content")
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to read response body")
		return "", err
	}

	var responseBody ResponseBody
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		log.Error().Err(err).Caller().Msg("failed to unmarshal response body")
		return "", err
	}

	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message.Content, nil
	}

	err = fmt.Errorf("no message content found")
	log.Error().Err(err).Caller().Msg("response contains no message content")
	return "", err
}
