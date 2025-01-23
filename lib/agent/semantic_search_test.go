package agent_test

import (
	"log"
	"testing"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/agent"
	"github.com/tmc/langchaingo/llms"
)


func Test_Search (t *testing.T) {
  //t.Deadline(5 * time.Second)
  // Testing autonomouse semantic_search agent. it will stop when it finds the answer
  result1 := agent.OneShotRun("Collection Name: 'Hellper' Query: How does embeddings package works?")
  log.Println("OneShotAskRun",result1)
}

func TestMemory(t *testing.T) {
  initialstate := []llms.MessageContent{
      llms.TextParts(llms.ChatMessageTypeSystem, "Below a current conversation between user and helpful AI assistant. Your task will be in the next system message"),
      llms.TextParts(llms.ChatMessageTypeHuman, "Hello my name is Yemet! I'm excited to see what we can do together."),
      llms.TextParts(llms.ChatMessageTypeAI, "Hey there! Let me know how I can help you out."),
      llms.TextParts(llms.ChatMessageTypeAI, "I'm here to assist you with any task you may have. Just give me a command and I'll do my best to help."),
  }

  result := agent.OneShotRun("Collection Name: 'Hellper' Query: How does embeddings package works? Also do you remember what is my name?", initialstate...)
  log.Println("Result:", result)

  // Assert or compare the result with the expected output
}





  /*
  completion_test, err := model.GenerateContent(context.Background(),intialState)
  if err != nil {
	log.Println("error with simple generate content",err)
  }
  log.Println("completion test: ", completion_test)
  */