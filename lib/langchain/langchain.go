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
// Initialize New Dialog thread with User with no limitation for token usage (may fail, use with limit)
func InitializeNewChatWithContextNoLimit(api_token string, model_name string, base_url string, user_initial_promt string) (*db.ChatSession, error) {
	//ctx := context.Background()

	cb := &ChainCallbackHandler{}

	if base_url == "" {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, err
		}

		memoryBuffer := memory.NewConversationBuffer()
		conversation := chains.NewConversation(llm, memoryBuffer) // create new conversation, which means langchain is modify initial promt in this moment. It is important, that your own template at local-ai side is also modifiyng template, so there might be a template collision.

		return &db.ChatSession{
			ConversationBuffer: *memoryBuffer,
			DialogThread:       conversation,
		}, nil
	} else {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
			//openai.WithBaseURL("http://localhost:8080"),
			openai.WithBaseURL(base_url),
			openai.WithAPIVersion("v1"),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, err
		}

		memoryBuffer := memory.NewConversationBuffer()
		conversation := chains.NewConversation(llm, memoryBuffer) // create new conversation, which means langchain is modify initial promt in this moment. It is important, that your own template at local-ai side is also modifiyng template, so there might be a template collision.

		return &db.ChatSession{
			ConversationBuffer: *memoryBuffer,
			DialogThread:       conversation,
		}, nil
	}

}

// Main function to start new conversation
func StartNewChat(ctx context.Context, api_token string, model_name string, base_url string, user_initial_promt string) (string, *db.ChatSession, error) {
	session, err1 := InitializeNewChatWithContextNoLimit(api_token, model_name, base_url, user_initial_promt)
	if err1 != nil {
		return "", nil, err1
	}
	result, post_session, err := RunChain(ctx, session, user_initial_promt)
	if err != nil {
		return "", nil, err
	}
	return result, post_session, nil
}

func RunChain(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error) {
	//ctx := context.Background()
	result, err := chains.Run(ctx, session.DialogThread, prompt)
	if err != nil {
		return "", nil, err
	}
	//chains.
	return result, session, nil
}



// TODO: HERE WE NEED TO REFACTOR TO WORK WITH NEW AGENTS SCHEMATICS. AFTER THAT WE CAN START DECREASE CODEBASE
// Continue Dialog with memory included, so user can chat with remembering context of previouse messages
func ContinueChatWithContextNoLimit(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error) {
	//ctx := context.Background()
	result, err := chains.Run(ctx, session.DialogThread, prompt)
	if err != nil {
		return "", nil, err
	}
	return result, session, nil
}

*/

/*
		HIGH LEVEL INSTRUCT (SUDO)
	 Main function for generating from single promt (without memory and context) --> this will result as Instruction, because it will not use langchain as template..

		Below is an instruction that describes a task. Write a response that appropriately completes the request.
	    Instruction: {{.Input}}
	    Response:
*/
func GenerateContentInstruction(base_url string, promt string, model_name string, api_token string, network string, options ...llms.CallOption) (string, error) {
	ctx := context.Background()
	var result string
	if network == "local" {
		llm, err := openai.New(
			openai.WithToken(api_token),
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
			openai.WithToken(api_token),
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
