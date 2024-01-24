package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChatRequest struct {
 Model      string     `json:"model"`
 Messages   []Message  `json:"messages"`
 Temperature float64    `json:"temperature"`
}

type Message struct {
 Role    string `json:"role"`
 Content string `json:"content"`
}

type ChatResponse struct {
 Created int            `json:"created"`
 Object  string         `json:"object"`
 ID      string         `json:"id"`
 Model   string         `json:"model"`
 Choices []Choice       `json:"choices"`
 Usage   UsageStatistics `json:"usage"`
}

type Choice struct {
 Index        int     `json:"index"`
 FinishReason string  `json:"finish_reason"`
 Message      Message `json:"message"`
}

type UsageStatistics struct {
 PromptTokens    int `json:"prompt_tokens"`
 CompletionTokens int `json:"completion_tokens"`
 TotalTokens     int `json:"total_tokens"`
}

func main() {
 prompt := "How are you?"
 modelName := "wizard-uncensored-13b"

 resp, err := GenerateCompletion(prompt, modelName)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 fmt.Println("Assistant's response:", resp.Choices[0].Message.Content)
}

func GenerateCompletion(prompt, modelName string) (*ChatResponse, error) {
 url := "http://localhost:8080/v1/chat/completions"

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

 return &chatResp, nil
}
