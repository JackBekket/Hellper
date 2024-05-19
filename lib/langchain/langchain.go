package langchain

//package langchain_controller
//package main

import (
	"context"
	"fmt"
	"log"

	db "github.com/JackBekket/uncensoredgpt_tgbot/lib/database"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"

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

// Initialize New Dialog thread with User with no limitation for token usage (may fail, use with limit)  initial_promt is first user message, (workaround for bug with LAI context)
func InitializeNewChatWithContextNoLimit(api_token string, model_name string, base_url string,user_initial_promt string) (*db.ChatSession, error)  {
	//ctx := context.Background()

	if base_url == "" {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
		)
		if err != nil {
			return nil, err
		}

		memoryBuffer := memory.NewConversationBuffer()
		conversation := chains.NewConversation(llm, memoryBuffer)	// create new conversation, which means langchain is modify initial promt in this moment. It is important, that your own template at local-ai side is also modifiyng template, so there might be a template collision.

		return &db.ChatSession{
			ConversationBuffer: *memoryBuffer,
			DialogThread: conversation,
		}, nil
	} else {
		llm, err := openai.New(
			openai.WithToken(api_token),
			openai.WithModel(model_name),
			openai.WithBaseURL("http://localhost:8080"),
			openai.WithAPIVersion("v1"),
		)
		if err != nil {
			return nil, err
		}
	
		memoryBuffer := memory.NewConversationBuffer()
		conversation := chains.NewConversation(llm, memoryBuffer)	// create new conversation, which means langchain is modify initial promt in this moment. It is important, that your own template at local-ai side is also modifiyng template, so there might be a template collision.
		
	
		return &db.ChatSession{
			ConversationBuffer: *memoryBuffer,
			DialogThread: conversation,
		}, nil
	}

}

func StartNewChat(api_token string, model_name string, base_url string,user_initial_promt string) (string,*db.ChatSession, error) {
	session, err1 :=InitializeNewChatWithContextNoLimit(api_token, model_name, base_url,user_initial_promt)
	if err1 != nil {
		return "",nil, err1
	}
	result,post_session,err := RunChain(session,user_initial_promt)
	if err != nil {
	return "",nil, err
	}
	return result,post_session,nil
}

func RunChain(session *db.ChatSession, prompt string) (string,*db.ChatSession, error) {
	ctx := context.Background()
    result, err := chains.Run(ctx, session.DialogThread, prompt)
    if err != nil {
        return "", nil, err
    }
    return result, session, nil
}

// Continue Dialog with memory included, so user can chat with remembering context of previouse messages
func ContinueChatWithContextNoLimit(session *db.ChatSession, prompt string) (string,*db.ChatSession, error) {
	ctx := context.Background()
    result, err := chains.Run(ctx, session.DialogThread, prompt)
    if err != nil {
        return "",nil, err
    }
    return result, session,nil
}





/* Main function for generating from single promt (without memory and context) --> this will result as Instruction, because it will not use langchain as template..
	     
	Below is an instruction that describes a task. Write a response that appropriately completes the request.
    Instruction: {{.Input}}
    Response:
	*/
func GenerateContentInstruction(promt string, model_name string,api_token string,network string) (string,error) {
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
	
	return result,nil
}

