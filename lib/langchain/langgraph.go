package langchain

import (
	"github.com/JackBekket/hellper/lib/agent"
	db "github.com/JackBekket/hellper/lib/database"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
)


func RunNewAgent(api_token string, model_name string, base_url string, user_promt string) (*db.ChatSessionGraph,string ,error) {
	cb := &ChainCallbackHandler{}

	if base_url == "" {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
			openai.WithCallback(cb),
		)
		if err != nil {
			return nil, "error",err
		}
		dialog_state, output_text := agent.RunThread(user_promt,*llm)
		//last_msg := dialog_state[len(dialog_state)-1]
		

		return &db.ChatSessionGraph{
			ConversationBuffer: dialog_state,
		}, output_text ,nil
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
			return nil,"error", err
		}

		dialog_state, output_text := agent.RunThread(user_promt,*llm)
		return &db.ChatSessionGraph{
			ConversationBuffer: dialog_state,
		}, output_text ,nil
	}
}


