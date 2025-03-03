package langchain

// OBSOLETE! USE LANGGRAPH.GO

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
)

/*
		HIGH LEVEL INSTRUCT (SUDO)
	 Main function for generating from single promt (without memory and context) --> this will result as Instruction, because it will not use langchain as template..

		Below is an instruction that describes a task. Write a response that appropriately completes the request.
	    Instruction: {{.Input}}
	    Response:
*/
func GenerateContentInstruction(base_url string, promt string, model_name string, localAIToken string, network string, options ...llms.CallOption) (string, error) {
	ctx := context.Background()
	var result string
	if network == "local" {
		llm, err := openai.New(
			openai.WithToken(localAIToken),
			//openai.WithBaseURL("http://localhost:8080"),
			openai.WithBaseURL(base_url),
			openai.WithModel(model_name),
			openai.WithAPIVersion("v1"),
		)
		if err != nil {
			log.Fatal(err)
		}

		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, promt, options...)
		if err != nil {
			// log.Fatal(err)
			return "", err
		}
		fmt.Println(completion)
		result = completion
		return completion, nil
	}
	if network == "openai" {
		llm, err := openai.New(
			openai.WithToken(localAIToken),
			openai.WithModel(model_name),
		)
		if err != nil {
			log.Fatal(err)
		}

		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, promt, options...)
		if err != nil {
			// log.Fatal(err)
			return "", err
		}
		fmt.Println(completion)
		result = completion
		return completion, nil
	}

	return result, nil
}
