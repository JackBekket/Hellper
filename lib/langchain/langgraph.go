package langchain

import (
	"github.com/JackBekket/hellper/lib/agent"
	db "github.com/JackBekket/hellper/lib/database"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
)

func RunNewAgent(aiToken string, model string, baseURL string, prompt string) (*db.ChatSessionGraph, string, error) {
	cb := &ChainCallbackHandler{}

	if baseURL == "" {
		llm, err := openai.New(
			openai.WithToken(aiToken),
			openai.WithModel(model),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, "error", err
		}
		dialogState, outputText := agent.RunThread(prompt, *llm)
		//last_msg := dialogState[len(dialogState)-1]

		return &db.ChatSessionGraph{
			ConversationBuffer: dialogState,
		}, outputText, nil
	} else {
		llm, err := openai.New(
			openai.WithToken(aiToken),
			openai.WithModel(model),
			openai.WithBaseURL(baseURL),
			openai.WithAPIVersion("v1"),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, "error", err
		}

		dialogState, outputText := agent.RunThread(prompt, *llm)
		return &db.ChatSessionGraph{
			ConversationBuffer: dialogState,
		}, outputText, nil
	}
}

func ContinueAgent(aiToken string, model string, baseURL string, prompt string, state *db.ChatSessionGraph) (*db.ChatSessionGraph, string, error) {
	cb := &ChainCallbackHandler{}

	if baseURL == "" {
		llm, err := openai.New(
			openai.WithToken(aiToken),
			openai.WithModel(model),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, "error", err
		}
		dialogState, outputText := agent.RunThread(prompt, *llm, state.ConversationBuffer...)

		return &db.ChatSessionGraph{
			ConversationBuffer: dialogState,
		}, outputText, nil
	} else {
		llm, err := openai.New(
			openai.WithToken(aiToken),
			openai.WithModel(model),
			//openai.WithBaseURL("http://localhost:8080"),
			openai.WithBaseURL(baseURL),
			openai.WithAPIVersion("v1"),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, "error", err
		}

		dialogState, outputText := agent.RunThread(prompt, *llm, state.ConversationBuffer...)
		return &db.ChatSessionGraph{
			ConversationBuffer: dialogState,
		}, outputText, nil
	}
}
