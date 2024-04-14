//package langchain

//package langchain_controller
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JackBekket/uncensoredgpt_tgbot/lib/bot/env"
	db "github.com/JackBekket/uncensoredgpt_tgbot/lib/database"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

/** DEV NOTE
	 OAI -- openAI, LAI -- localAI
	 if your IDE says it won't compile just try to build from terminal first
	 if it says there no methods "Run" or "Predict" in LLM class -- it is weird bug, just compile it from terminal
**/

// I use it for fast testing
func main()  {
	//ctx := context.Background()
	env.Load()
	//env_data := env.LoadAdminData()
	token := env.GetAdminToken()
	//model_name := "gpt-3.5-turbo"	// using openai for tests
	model_name := "wizard-uncensored-13b"


	user_initial_promt := "Hello, my name is Bekket, I am working on a new project called 'Andromeda'."
	//ai_initial_promt := "Hello Bekket, seems like a great name for a project!"
	//check_promt := "What is my name and what project am I currently working on?"

	/*
	result, err :=TestChatWithContextNoLimit(token,model_name)		// works with both OAI and LAI
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
	*/

	
	
	//bug
	
	session, err := StartNewChat(token,model_name,"localai",user_initial_promt,)
	if err != nil {
		log.Println(err)
	}
	

	
	//memory := session.ConversationBuffer		// tha'ts a weird thing, Initialize and Continue works in pair only if I get memory buffer in here

	//memory.ChatHistory.AddUserMessage(ctx,"I am working on a new project called 'Andromeda'")
	//memory.ChatHistory.AddAIMessage(ctx,"I like it!")
	

	
	//res1,err := ContinueChatWithContextNoLimit(session,"I am working on a new golang project called 'Andromeda', do you like this project name?")
	res1,err := ContinueChatWithContextNoLimit(session,"What's my name and what is the name of the project I currently working on?")
	//res1,err := ContinueChatWithContextNoLimit(session,"What is my name?")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("answer 1",res1)
	

	
	// bug in sequential chain
	
	res2, err := ContinueChatWithContextNoLimit(session,"What is the name of the project I currently working on?")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("answer2",res2)
	
	

	/*
	//memory := session.ConversationBuffer
	log.Println("check if it's stored in messages, printing messages:")
	history, err := memory.ChatHistory.Messages(ctx)
	if err != nil {
		log.Println(err)
	}
	//log.Println(history)
	total_turns := len(history)
	log.Println("total number of turns: ", total_turns)
	// Iterate over each message and print
    log.Println("Printing messages:")
    for _, msg := range history {
        log.Println(msg.GetContent())
    }
	*/
	
}


// TODO: make universal function to OAI and LOI, add base_url as argument probably
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
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a helpfull assistant who help in whatever task human ask you about"),
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
	return completion, nil
}


// chat without context
func GenerateContentLAI(api_token string, model_name string, promt string) (*llms.ContentResponse, error) {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel(model_name),
		//llms.WithOptions()
		openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithAPIVersion("v1"),
	)
	if err != nil {
	  log.Fatal(err)
	}
	

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a helpfull assistant who help in whatever task human ask you about"),
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

	return completion, nil
}



// DEBUG NOTE -- this thing work
// chat with context without limitation of token to use
//  use it only to fast testing
func TestChatWithContextNoLimit(api_token string, model_name string) (string, error) {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel(model_name),
		//llms.WithOptions()
		openai.WithBaseURL("http://localhost:8080/v1/"),	// comment this and next line to call OAI, if not then call to LAI
		openai.WithAPIVersion("v1"),
	)
	if err != nil {
	  log.Fatal(err)
	}

	memory_buffer := memory.NewConversationBuffer()

	//test data
	// First dialogue pair
	inputValues1 := map[string]any{"input": "Hi"}				// ignore linter
	outputValues1 := map[string]any{"output": "What's up"}

	memory_buffer.SaveContext(ctx,inputValues1,outputValues1)	//initial messages should be put like this

	memory_buffer.ChatHistory.AddUserMessage(ctx, "Not much, just hanging")  	// next messages from conversation could be added like this
	memory_buffer.ChatHistory.AddAIMessage(ctx,"Cool")
	memory_buffer.ChatHistory.AddUserMessage(ctx, "I am working at my new exiting golang AI project called 'Andromeda'")
	memory_buffer.ChatHistory.AddUserMessage(ctx, "My name is Bekket btw")
	
	conversation := chains.NewConversation(llm,memory_buffer) 	// build chain, start new conversation thread
	

	// Run is used when we have only one input (promt for example).   If there are need in passing few inputs then use chains.Call instead
	result, err := chains.Run(ctx,conversation,"what is my name and what project am I currently working on?")	//ignore linter error
	if err != nil {
		return "", err
	}

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

	log.Println("AI answer:")
	log.Println(result)

	log.Println("check if it's stored in messages, printing messages:")
	history, err := memory_buffer.ChatHistory.Messages(ctx)
	if err != nil {
		return "", err
	}
	//log.Println(history)
	total_turns := len(history)
	log.Println("total number of turns: ", total_turns)
	// Iterate over each message and print
    log.Println("Printing messages:")
    for _, msg := range history {
        log.Println(msg.GetContent())
    }

	return result,err
}


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
		//memoryBuffer.ChatHistory.AddUserMessage(ctx,user_initial_promt)

		conversation := chains.NewConversation(llm, memoryBuffer)	// create new conversation, which means langchain is modify initial promt in this moment. It is important, that your own template at local-ai side is also modifiyng template, so there might be a template collision.
		
	
		return &db.ChatSession{
			ConversationBuffer: *memoryBuffer,
			DialogThread: conversation,
		}, nil
	}

}

func StartNewChat(api_token string, model_name string, base_url string,user_initial_promt string) (*db.ChatSession, error) {
	session, err1 :=InitializeNewChatWithContextNoLimit(api_token, model_name, base_url,user_initial_promt)
	if err1 != nil {
		return nil, err1
	}
	_,err,post_session :=RunChain(session,user_initial_promt)
	if err != nil {
	return nil, err
	}
	return post_session,nil
}

func RunChain(session *db.ChatSession, prompt string) (string, error,*db.ChatSession) {
	ctx := context.Background()
    result, err := chains.Run(ctx, session.DialogThread, prompt)
    if err != nil {
        return "", err, nil
    }
    return result, nil, session
}

// Continue Dialog with memory included, so user can chat with remembering context of previouse messages
func ContinueChatWithContextNoLimit(session *db.ChatSession, prompt string) (string, error) {
	ctx := context.Background()
    result, err := chains.Run(ctx, session.DialogThread, prompt)
    if err != nil {
        return "", err
    }
    return result, nil
}



// TODO: remove or transfer this into tests
func TestOAI(api_token string)  {
	ctx := context.Background()
	token := api_token

	llm, err := openai.New(
		openai.WithToken(token),
		openai.WithModel("gpt-3.5-turbo"),
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



// TODO: make one function for both OAI & LAI, add baseUrl as argument
// Main function for generating from single promt (without memory and context)
func GenerateFromSinglePromtLocal(prompt string, model_name string) (string,error) {
	ctx := context.Background()
	llm, err := openai.New(
		//openai.WithToken()
		openai.WithBaseURL("http://localhost:8080"),
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

