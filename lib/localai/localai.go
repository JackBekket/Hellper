package localai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

var ErrConnectionFailure = errors.New("сonnection error")

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Created int             `json:"created"`
	Object  string          `json:"object"`
	ID      string          `json:"id"`
	Model   string          `json:"model"`
	Choices []Choice        `json:"choices"`
	Usage   UsageStatistics `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type UsageStatistics struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GenerationResponse struct {
	Created int64            `json:"created"`
	ID      string           `json:"id"`
	Data    []GenerationData `json:"data"`
	Usage   GenerationUsage  `json:"usage"`
}

type GenerationData struct {
	Embedding interface{} `json:"embedding"`
	Index     int         `json:"index"`
	URL       string      `json:"url"`
}

type GenerationUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type WrongPwdError struct {
	message string
}

func (e *WrongPwdError) Error() string {
	return e.message
}

type OpenAIDataObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

type OpenAIModelsResponse struct {
	Data []OpenAIDataObject `json:"data"`
}

func GetModelsList(url, token string) ([]string, error) {
	modelsList := []string{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error retrieving the list of models ")
		return modelsList, ErrConnectionFailure
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error retrieving the response")
		return modelsList, ErrConnectionFailure
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return modelsList, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var modelsResp OpenAIModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
		return modelsList, err
	}

	for _, obj := range modelsResp.Data {
		modelsList = append(modelsList, obj.ID)
	}

	return modelsList, nil
}

func GenerateCompletion(prompt, modelName, url string) (*ChatResponse, error) {
	data := ChatRequest{
		Model: modelName,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.9,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to marshal chat request")
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Caller().Msg("error creating new request for chat completion")
		return nil, ErrConnectionFailure
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error sending request for chat completion")
		return nil, ErrConnectionFailure
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status_code", resp.StatusCode).Caller().Msg("unexpected status code for chat completion")
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to read response body")
		return nil, err
	}

	var chatResp ChatResponse
	if err = json.Unmarshal(body, &chatResp); err != nil {
		log.Error().Err(err).Caller().Msg("failed to unmarshal chat response")
		return nil, err
	}

	log.Info().Msg("chat completion response received successfully")
	return &chatResp, nil
}

func GenerateImageStableDiffusion(prompt, size, url, model, localAIToken string) (string, error) {
	payload := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Size   string `json:"size"`
	}{
		Model:  model,
		Prompt: prompt,
		Size:   size,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to marshal payload")
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to create new request")
		return "", ErrConnectionFailure
	}
	req.Header.Set("Content-Type", "application/json")

	if localAIToken == "" {
		err := fmt.Errorf("localAIToken not found")
		log.Error().Err(err).Caller().Msg("authorization token missing")
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+localAIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to send request")
		return "", ErrConnectionFailure
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		log.Error().Int("status_code", resp.StatusCode).Msgf("unexpected status code: %s", string(errorBody))
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(errorBody))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("failed to read response body")
		return "", err
	}

	var generationResp GenerationResponse
	if err = json.Unmarshal(body, &generationResp); err != nil {
		log.Error().Err(err).Caller().Msg("failed to unmarshal response")
		return "", err
	}

	if len(generationResp.Data) == 0 {
		err := fmt.Errorf("no image data returned")
		log.Error().Err(err).Caller().Msg("empty response data")
		return "", err
	}

	imageURL := generationResp.Data[0].URL
	log.Info().Msgf("Image URL from localai pkg: %s", imageURL)
	return imageURL, nil
}
func TranscribeWhisper(url, model, path, localAIToken string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error opening file")
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err = writer.WriteField("model", model); err != nil {
		log.Error().Err(err).Caller().Msg("error adding model field")
		return "", err
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		log.Error().Err(err).Caller().Msg("error creating form file")
		return "", err
	}
	if _, err = io.Copy(part, file); err != nil {
		log.Error().Err(err).Caller().Msg("error copying file data")
		return "", err
	}

	if err = writer.Close(); err != nil {
		log.Error().Err(err).Caller().Msg("error closing writer")
		return "", err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error creating request")
		return "", ErrConnectionFailure
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("accept", "application/json")
	if localAIToken == "" {
		err := fmt.Errorf("localAIToken not found")
		log.Error().Err(err).Caller().Msg("authorization token missing")
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+localAIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error sending request")
		return "", ErrConnectionFailure
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		log.Error().Int("status_code", resp.StatusCode).Msgf("unexpected status code: %s", string(errorBody))
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(errorBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Caller().Msg("error reading response")
		return "", err
	}

	var response struct {
		Text string `json:"text"`
	}
	if err = json.Unmarshal(respBody, &response); err != nil {
		log.Error().Err(err).Caller().Msg("error unmarshalling response")
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	text := cleanText(response.Text)
	log.Info().Msg("transcription completed successfully")
	return text, nil
}

func cleanText(input string) string {
	if strings.Contains(input, "[BLANK_AUDIO]") && len(strings.Trim(input, "[BLANK_AUDIO]")) == 0 {
		return "[BLANK_AUDIO]"
	}
	return strings.ReplaceAll(input, "[BLANK_AUDIO]", "")
}
