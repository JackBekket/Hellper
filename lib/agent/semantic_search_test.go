package agent_test

import (
	"log"
	"os"
	"testing"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/agent"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)


func Test_Search (t *testing.T) {
  //t.Deadline(5 * time.Second)
  // Testing OnePunch semantic_search agent. it will stop when it finds the answer
  log.Println("testing one-shot search")
  model := createGenericLLM()
  result1 := agent.OneShotRun("Call semanticSearch tool. Collection Name: 'Hellper' Query: How does embeddings package works?",model)
  log.Println("OneShotAskRun",result1)
}

func TestMemory(t *testing.T) {
  log.Println("testing with memory")
  initialstate := []llms.MessageContent{
      llms.TextParts(llms.ChatMessageTypeSystem, "Below a current conversation between user and helpful AI assistant. Your task will be in the next system message"),
      llms.TextParts(llms.ChatMessageTypeHuman, "Hello my name is Yemet! I'm excited to see what we can do together."),
      llms.TextParts(llms.ChatMessageTypeAI, "Hey there! Let me know how I can help you out."),
      llms.TextParts(llms.ChatMessageTypeAI, "I'm here to assist you with any task you may have. Just give me a command and I'll do my best to help."),
  }
  model := createGenericLLM()
  result := agent.OneShotRun("Collection Name: 'Hellper' Query: How does main package works? Call semanticSearch tool. Also do you remember what is my name?",model ,initialstate...)
  log.Println("Result:", result)

  // Assert or compare the result with the expected output
}




func createGenericLLM() openai.LLM{
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


  /*
  completion_test, err := model.GenerateContent(context.Background(),intialState)
  if err != nil {
	log.Println("error with simple generate content",err)
  }
  log.Println("completion test: ", completion_test)
  */