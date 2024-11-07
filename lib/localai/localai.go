package localai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	prompt := "How are you?"
	modelName := "wizard-uncensored-13b"
	url := "http://localhost:8080/v1/chat/completions"

	resp, err := GenerateCompletion(prompt, modelName, url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Assistant's response:", resp.Choices[0].Message.Content)
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

func GenerateCompletionWithPWD(prompt, modelName string, url string, s_pwd string, u_pwd string) (*ChatResponse, error) {
	if u_pwd != s_pwd {
		err := &WrongPwdError{"wrong password"}
		return nil, err
	} else {
		result, err := GenerateCompletion(prompt, modelName, url)
		if err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
}

func GenerateImageStableDiffusion(prompt, size, url string) (string, error) {
	//TODO: envvvvv
	url += "/v1/images/generations"
	fmt.Println("Request URL:", url)
	payload := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Size   string `json:"size"`
	}{
		//TODO: update to get it from .env
		Model:  "stablediffusion-cpp",
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
