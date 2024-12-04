package agent_test

import (
	"testing"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/agent"
)


func Test_Search (t *testing.T) {

  // Testing autonomouse semantic_search agent. it will stop when it finds the answer
  agent.SearchRun()
}



/*
func TestRun(t *testing.T) {
  // Create a new agent
  //agentRun := agent.Run()

  // Initialize the state
  //intialState := make([]llms.MessageContent, 0)

  // Add an initial message to the state
  //intialState = append(intialState, llms.TextParts(llms.ChatMessageTypeHuman, "Hello"))

  // Create a channel
  ch := make(chan llms.MessageContent)
  go func() {
      for i := 0; i < 5; i++ {
          msg := llms.TextParts(llms.ChatMessageTypeHuman, fmt.Sprintf("What do you think of it? %d", i))
          ch <- msg
          time.Sleep(100 * time.Millisecond)
      }
      close(ch)
  }()

  agent.Run(ch)

  
}
*/