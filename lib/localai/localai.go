package localai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

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

func GetModelsList(endpoint, token string) ([]string, error) {
	modelsList := []string{}
	urlPath, err := url.JoinPath(endpoint, "models")
	if err != nil {
		return modelsList, err
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return modelsList, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return modelsList, err
	}

	modelsResp := OpenAIModelsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&modelsResp)
	if err != nil {
		return modelsList, err
	}

	for _, obj := range modelsResp.Data {
		modelsList = append(modelsList, obj.ID)
	}

	return modelsList, nil
}




func GenerateCompletion(prompt, modelName string, url string) (*ChatResponse, error) {

	//url := "http://localhost:8080/v1/chat/completions"

	// Create the request body
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

	// Convert request body to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// log raw response
	log.Println("raw response: ", resp)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var chatResp ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return nil, err
	}

	// log unmarshalled response
	log.Println(chatResp)

	return &chatResp, nil
}

func GenerateImageStableDiffusion(prompt, size, url, model string) (string, error) {
	fmt.Println("Request URL:", url)
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
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Get the API key from the environment and add it to the Authorization header
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		return "", fmt.Errorf("API key not found in environment variables")
	}
	req.Header.Set("Authorization", "Bearer "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log if the request fails
	if resp.StatusCode != http.StatusOK {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(errorBody))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var generationResp GenerationResponse
	err = json.Unmarshal(body, &generationResp)
	if err != nil {
		return "", err
	}
	imageURL := generationResp.Data[0].URL
	fmt.Println("Image URL from localai pkg:", imageURL)

	return imageURL, nil
}

func TranscribeWhisper(url, model, path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err = writer.WriteField("model", model)
	if err != nil {
		fmt.Println("Error adding model field:", err)
		return "", err
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		fmt.Println("Error creating file field:", err)
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file data:", err)
		return "", err
	}

	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	var response struct {
		Text string `json:"text"`
	}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	text := response.Text

	//uncomment to remove [BLANK_AUDIO] from output.
	text = cleanText(text)

	return text, nil
}

func cleanText(input string) string {
	if strings.Contains(input, "[BLANK_AUDIO]") && len(strings.Trim(input, "[BLANK_AUDIO]")) == 0 {
		return "[BLANK_AUDIO]"
	}
	return strings.ReplaceAll(input, "[BLANK_AUDIO]", "")
}