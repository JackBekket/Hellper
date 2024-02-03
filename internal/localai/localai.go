package localai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type GenerationResponse struct {
    Created int64 `json:"created"`
    ID      string `json:"id"`
    Data    []GenerationData `json:"data"`
    Usage   GenerationUsage `json:"usage"`
}

type GenerationData struct {
    Embedding interface{} `json:"embedding"`
    Index     int `json:"index"`
    URL       string `json:"url"`
}
   
type GenerationUsage struct {
    PromptTokens      int `json:"prompt_tokens"`
    CompletionTokens  int `json:"completion_tokens"`
    TotalTokens       int `json:"total_tokens"`
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
    if (u_pwd != s_pwd) {
        err := &WrongPwdError{"wrong password"}
        return nil, err
    } else {
        result, err := GenerateCompletion(prompt,modelName,url)
        if err != nil {
            return nil, err
        } else {
            return result,nil
        }
    }
}


    func GenerateImageStableDissusion(prompt, size string) (string, error) {
        url := "http://localhost:8080/v1/images/generations"
       
        payload := struct {
         Prompt string `json:"prompt"`
         Size   string `json:"size"`
        }{
         Prompt: prompt,
         Size:   size,
        }
       
        payloadBytes, err := json.Marshal(payload)
        if err != nil {
         return "", err
        }
       
        resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
        if err != nil {
         return "", err
        }
        defer resp.Body.Close()
       
        if resp.StatusCode != http.StatusOK {
         return "", fmt.Errorf("Request failed with status code %d", resp.StatusCode)
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
        image_url := generationResp.Data[0].URL
        fmt.Println("image url: ", image_url)
       
        uploadURL, err := uploadToTelegraph(generationResp.Data[0].URL)
        if err != nil {
         return "", err
        }
       
        return uploadURL, nil
       }


func uploadToTelegraph(fileURL string) (string, error) {
        // Implement the logic to upload the file to Telegraph and get the generated link
        // Here's just a placeholder returning the fileURL as the result
        return fileURL, nil
       }


    
