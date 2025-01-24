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
 call := OneShotRun(prompt,*&llm)
 log.Println(call)
}


// this function recive previouse history message state and append new user prompt, than run agent
func RunThread(prompt string, model openai.LLM, history ...llms.MessageContent) ([]llms.MessageContent, string){
	
	//model := createGenericLLM()
	call := OneShotRun(prompt, model, history...)
	log.Println(call)
	lastResponse := createMessageContent(call)
	if len(history) > 0 { 
		state := append(history, lastResponse...)
		return state,call
	} else {
		state := lastResponse
		return state,call
	}
}


func createMessageContent (content string) []llms.MessageContent{
	intialState := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeAI, content),
	  }
	return intialState
}


func CreateGenericLLM() openai.LLM{
	model_name := "tiger-gemma-9b-v1-i1"    // should be settable?
	_ = godotenv.Load()
			ai_url := os.Getenv("AI_ENDPOINT")          //TODO: should be global?
			api_token := os.Getenv("ADNIN_KEY")
			//db_link := os.Getenv("EMBEDDINGS_DB_URL")
	model, err := openai.New(
	  openai.WithToken(api_token),
	  //openai.WithBaseURL("http://localhost:8080"),
	  openai.WithBaseURL(ai_url),
	  openai.WithModel(model_name),
	  openai.WithAPIVersion("v1"),
	)
	if err != nil {
	  log.Fatal(err)
	}
	return *model
}