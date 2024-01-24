package openaibot

import (
	"errors"
	"log"

	gogpt "github.com/sashabaranov/go-openai"
	//"github.com/JackBekket/telegram-gpt/internal/bot/env"
)

func CreateClient(gptKey string) *gogpt.Client {
	return gogpt.NewClient(gptKey)
	
}

//
func CreateDefConfig(gptKey string, baseURL string) *gogpt.Client {
	cfg := gogpt.DefaultConfig(gptKey)
	cfg.BaseURL = baseURL
	return gogpt.NewClientWithConfig(cfg)	
}


// create anonymouse client with baseURL
func CreateCustomBaseConfig(baseURL string) *gogpt.Client {
	cfg := gogpt.DefaultConfig("")
	cfg.BaseURL = baseURL
	return gogpt.NewClientWithConfig(cfg)	
}


func CreateLocalhostClient() *gogpt.Client {
	cfg := gogpt.DefaultConfig("")
	cfg.BaseURL = "http://127.0.0.1:8080"
	return gogpt.NewClientWithConfig(cfg)
}

func CreateLocalhostClientWithCheck(lpwd string,user_promt string) (*gogpt.Client, error) {
	if (lpwd == user_promt) {
		log.Println(lpwd)
		log.Println(user_promt)
		log.Println("creating localhost client")
		cfg := gogpt.DefaultConfig(user_promt)
		cfg.BaseURL = "http://127.0.0.1:8080"
		return gogpt.NewClientWithConfig(cfg),nil
	} else {
		log.Println("password to access localhost AI is not correct")
		//return CreateClient(user_promt)
		return nil, errors.New("local pwd incorrect")
	}

}
