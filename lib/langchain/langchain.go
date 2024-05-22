package langchain

//package langchain_controller
//package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/JackBekket/hellper/lib/database"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
)

/** DEV NOTE
	 OAI -- openAI, LAI -- localAI
	 if your IDE says it won't compile just try to build from terminal first
	 if it says there no methods "Run" or "Predict" in LLM class -- it is weird bug, just compile it from terminal
**/

/*
	you can get conversation logs by docker logs -f local-ai
	(if you run local-ai in DEBUG mode)
*/
/*
// I use it for fast testing
func main()  {
	//ctx := context.Background()
	env.Load()
	//env_data := env.LoadAdminData()
	token := env.GetAdminToken()
	//model_name := "gpt-3.5-turbo"	// using openai for tests
	model_name := "wizard-uncensored-13b"

	user_initial_promt := "Hello, my name is Bekket, I am working on a new project called 'Andromeda'."

	result, err :=TestChatWithContextNoLimit(token,model_name)		// works with both OAI and LAI
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)


	session, err := StartNewChat(token,model_name,"localai",user_initial_promt,)
	if err != nil {
		log.Println(err)
	}
	res1,err := ContinueChatWithContextNoLimit(session,"What's my name and what is the name of the project I currently working on?")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("answer 1",res1)
}
*/

/*
// return ContentResponse instead of single string result
func GenerateContent(api_token string, model_name string, promt string, network string) (*llms.ContentResponse, error) {
	ctx := context.Background()
	token := api_token
	var llm_ *openai.LLM
	if network == "localai" {
		//base_url = "http://localhost:8080/v1/"
		llm, err := openai.New(
			openai.WithToken(token),
			openai.WithModel(model_name),
			//llms.WithOptions()
			openai.WithBaseURL("http://localhost:8080/v1/"),
			openai.WithAPIVersion("v1"),
		)
		if err != nil {
			return nil, err
		}
		llm_ = llm
	}
	if network == "openai" {
		//base_url
		llm, err := openai.New(
			openai.WithToken(token),
			openai.WithModel(model_name),
		)
		if err != nil {
			return nil, err
		}
		llm_ = llm
	}

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a helpfull assistant who help in whatever task human ask you about"),
		llms.TextParts(schema.ChatMessageTypeHuman, promt),
	}

	completion, err := llm_.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	return completion, nil
}
*/

// Example using call with few inputs
/*
		translatePrompt := prompts.NewPromptTemplate(
		"Translate the following text from {{.inputLanguage}} to {{.outputLanguage}}. {{.text}}",
		[]string{"inputLanguage", "outputLanguage", "text"},
	)
	llmChain = chains.NewLLMChain(llm, translatePrompt)

	// Otherwise the call function must be used.
	outputValues, err := chains.Call(ctx, llmChain, map[string]any{
		"inputLanguage":  "English",
		"outputLanguage": "French",
		"text":           "I love programming.",
	})
	if err != nil {
		return err
	}

	out, ok := outputValues[llmChain.OutputKey].(string)
	if !ok {
		return fmt.Errorf("invalid chain return")
	}
	fmt.Println(out)
*/

type ChainCallbackHandler struct{}

// HandleAgentAction implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleAgentAction(ctx context.Context, action schema.AgentAction) {
	panic("unimplemented")
}

// HandleAgentFinish implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleAgentFinish(ctx context.Context, finish schema.AgentFinish) {
	//panic("unimplemented")
}

// HandleChainEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {
	//panic("unimplemented")
}

// HandleChainError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleChainStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {
	//panic("unimplemented")
}

// HandleLLMError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleLLMGenerateContentStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent) {
	//panic("unimplemented")
}

// HandleLLMStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMStart(ctx context.Context, prompts []string) {
	//panic("unimplemented")
}

// HandleRetrieverEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document) {
	//panic("unimplemented")
}

// HandleRetrieverStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleRetrieverStart(ctx context.Context, query string) {
	//panic("unimplemented")
}

// HandleStreamingFunc implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleStreamingFunc(ctx context.Context, chunk []byte) {
	//panic("unimplemented")
}

// HandleToolEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolEnd(ctx context.Context, output string) {
	//panic("unimplemented")
}

// HandleToolError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleToolStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolStart(ctx context.Context, input string) {
	//panic("unimplemented")
}

func (h *ChainCallbackHandler) HandleText(ctx context.Context, text string) {
	// Implement this method if needed
}

// Implement other methods...

func (h *ChainCallbackHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {
	/*
	  // Extract the headers you're interested in
	  tokensUsed := res.Header.Get("Openai-Usage-Tokens")
	  promptTokens := res.Header.Get("Openai-Usage-Prompt-Tokens")
	  completionTokens := res.Header.Get("Openai-Usage-Completion-Tokens")

	  fmt.Println("Tokens Used:", tokensUsed)
	  fmt.Println("Prompt Tokens:", promptTokens)
	  fmt.Println("Completion Tokens:", completionTokens)

	*/
	LogResponseContentChoice(res)
}

func LogResponseContentChoice(resp *llms.ContentResponse) {
	//choice *llms.ContentChoice
	choice := resp.Choices[0]
	log.Println("Content: ", choice.Content)
	log.Println("Stop Reason: ", choice.StopReason)
	//t :=resp.Usage.CompletionTokens

	// GenerationInfo is a map that could contain complex/nested structures,
	// so we'll marshal it into a JSON string for a cleaner log message.
	// This step is optional and depends on your preference for log clarity.
	genInfo, err := json.Marshal(choice.GenerationInfo)
	if err != nil {
		log.Println("Error marshaling GenerationInfo: ", err)
		return
	}
	log.Println("Generation Info: ", string(genInfo))

	// If you have specific fields you expect in GenerationInfo, you can log them individually:
	// Example: log.Println("Some specific gen info: ", choice.GenerationInfo["someKey"])
	log.Println("Some specific gen info: ", choice.GenerationInfo["PromptTokens"])
	log.Println("Some specific gen info: ", choice.GenerationInfo["CompletionTokens"])
	log.Println("Some specific gen info: ", choice.GenerationInfo["TotalTokens"])

	// Note: Since FuncCall is a pointer to a schema.FunctionCall, ensure you check for nil to avoid panics.
	if choice.FuncCall != nil {
		// Assuming FuncCall has fields you want to log, replace 'FieldName' with actual fields.
		log.Printf("Function Call: %+v\n", choice.FuncCall)
		// For specific field: log.Println("FuncCall field: ", choice.FuncCall.FieldName)
	} else {
		log.Println("No Function Call requested.")
	}
}

// Initialize New Dialog thread with User with no limitation for token usage (may fail, use with limit)  initial_promt is first user message, (workaround for bug with LAI context)
func InitializeNewChatWithContextNoLimit(api_token string, model_name string, base_url string, user_initial_promt string) (*db.ChatSession, error) {
	//ctx := context.Background()

	cb := &ChainCallbackHandler{}

	if base_url == "" {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
			//openai.WithCallback()
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

func StartNewChat(api_token string, model_name string, base_url string, user_initial_promt string) (string, *db.ChatSession, error) {
	session, err1 := InitializeNewChatWithContextNoLimit(api_token, model_name, base_url, user_initial_promt)
	if err1 != nil {
		return "", nil, err1
	}
	result, post_session, err := RunChain(session, user_initial_promt)
	if err != nil {
		return "", nil, err
	}
	return result, post_session, nil
}

func RunChain(session *db.ChatSession, prompt string) (string, *db.ChatSession, error) {
	ctx := context.Background()
	result, err := chains.Run(ctx, session.DialogThread, prompt)
	if err != nil {
		return "", nil, err
	}
	//chains.
	return result, session, nil
}

// Continue Dialog with memory included, so user can chat with remembering context of previouse messages
func ContinueChatWithContextNoLimit(session *db.ChatSession, prompt string) (string, *db.ChatSession, error) {
	ctx := context.Background()
	result, err := chains.Run(ctx, session.DialogThread, prompt)
	if err != nil {
		return "", nil, err
	}
	return result, session, nil
}

/*
	 Main function for generating from single promt (without memory and context) --> this will result as Instruction, because it will not use langchain as template..

		Below is an instruction that describes a task. Write a response that appropriately completes the request.
	    Instruction: {{.Input}}
	    Response:
*/
func GenerateContentInstruction(promt string, model_name string, api_token string, network string) (string, error) {
	ctx := context.Background()
	var result string
	if network == "local" {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithBaseURL("http://localhost:8080"),
			openai.WithModel(model_name),
			openai.WithAPIVersion("v1"),
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

		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, promt)
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
