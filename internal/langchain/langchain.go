//package langchain_controller
package main

import (
	"context"
	"fmt"
	"log"

	//langchain "github.com/tmc/langchaingo"
	"github.com/JackBekket/uncensoredgpt_tgbot/internal/bot/env"
	"github.com/tmc/langchaingo/llms"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func main()  {
	//ctx := context.Background()
	env.Load()
	//env_data := env.LoadAdminData()
	token := env.GetAdminToken()

	/*
	completion,err := GenerateContentOAI(token,"gpt-3.5-turbo","What would be a good company name a company that makes colorful socks? Write at least 10 options")
	if err != nil {
		log.Println(err)
	}
	*/

	completion, err := GenerateContentLAI(token,"wizard-uncensored-13b", "What would be a good company name a company that makes colorful socks? Write at least 10 options")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(completion.Choices[0].Content)
}

func GenerateContentOAI(api_token string, model_name string, promt string) (*llms.ContentResponse, error) {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel(model_name),
		//llms.WithOptions()
		//openai.WithBaseURL("http://localhost:8000"),
	)
	if err != nil {
	  log.Fatal(err)
	}

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a helpfull assistant who execute whatever task human gives you"),
		llms.TextParts(schema.ChatMessageTypeHuman, promt),
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//fmt.Println(completion)
	return completion, nil
}


func GenerateContentLAI(api_token string, model_name string, promt string) (*llms.ContentResponse, error) {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel(model_name),
		//llms.WithOptions()
		openai.WithBaseURL("http://localhost:8000/v1/"),
		openai.WithAPIVersion("v1"),
	)
	if err != nil {
	  log.Fatal(err)
	}
	//log.Println(llm.)

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a helpfull assistant who execute whatever task human gives you"),
		llms.TextParts(schema.ChatMessageTypeHuman, promt),
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//llms.WithOptions()
	//fmt.Println(completion)
	return completion, nil
}

func TestOAI(api_token string)  {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel("gpt-3.5-turbo"),
		//llms.WithOptions()
		//openai.WithBaseURL("http://localhost:8000"),
	)
	if err != nil {
	  log.Fatal(err)
	}
	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a company branding design wizard."),
		llms.TextParts(schema.ChatMessageTypeHuman, "What would be a good company name a company that makes colorful socks? Write at least 10 options"),
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}


func GenerateFromSinglePromtLocal(prompt string, model_name string) (string,error) {
	ctx := context.Background()
	llm, err := openai.New(
		//openai.WithToken()
		openai.WithBaseURL("http://localhost:8000"),
		openai.WithModel(model_name),
	)
	if err != nil {
	  log.Fatal(err)
	}
	
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
	 // log.Fatal(err)
	 return "", err
	}
	fmt.Println(completion)
	return completion, nil
}

func GenerateFromSinglePromtOAI(promt string, model_name string, api_token string) (string , error) {
	ctx := context.Background()
	llm, err := openai.New(
		openai.WithToken(api_token),
		//openai.WithBaseURL("http://localhost:8000"),
		openai.WithModel(model_name),
	)
	if err != nil {
	  log.Fatal(err)
	}
	
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, promt)
	if err != nil {
	 // log.Fatal(err)
	 return "", err
	}
	fmt.Println(completion)
	return completion, nil
}