// This will be prototype to superagent (autonomouse agent, which work with memory and have similar functionality to langchain chains.Run method)
package agent

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// This function fire One-Shot agent without history context
func OnePunch(prompt string) {

	llm := CreateGenericLLM()
	call := OneShotRun(prompt, llm)
	log.Println(call)
}

// this function recive previouse history message state and append new user prompt, than run agent
func RunThread(prompt string, model openai.LLM, history ...llms.MessageContent) ([]llms.MessageContent, string) {

	//model := createGenericLLM()
	call := OneShotRun(prompt, model, history...)
	log.Println(call)
	lastResponse := CreateMessageContentAi(call)
	if len(history) > 0 {
		user_msg := CreateMessageContentHuman(prompt)
		state := append(history, user_msg[0])
		state = append(state, lastResponse...)
		return state, call
	} else {
		user_msg := CreateMessageContentHuman(prompt)
		state := user_msg
		state = append(state, lastResponse...)
		return state, call
	}
}

func CreateMessageContentAi(content string) []llms.MessageContent {
	intialState := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeAI, content),
	}
	return intialState
}

func CreateMessageContentHuman(content string) []llms.MessageContent {
	intialState := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, content),
	}
	return intialState
}

// refactor
func CreateGenericLLM() openai.LLM {
	model_name := "tiger-gemma-9b-v1-i1" // should be settable?
	_ = godotenv.Load()
	baseURL := os.Getenv("AI_BASEURL") //TODO: should be global?
	localAIToken := os.Getenv("AI_BASEURL")
	//db_link := os.Getenv("EMBEDDINGS_DB_URL")
	model, err := openai.New(
		openai.WithToken(localAIToken),
		//openai.WithBaseURL("http://localhost:8080"),
		openai.WithBaseURL(baseURL),
		openai.WithModel(model_name),
		openai.WithAPIVersion("v1"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return *model
}
