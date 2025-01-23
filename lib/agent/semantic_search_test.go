package agent_test

import (
	"testing"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/agent"
)


func Test_Search (t *testing.T) {
  //t.Deadline(5 * time.Second)
  // Testing autonomouse semantic_search agent. it will stop when it finds the answer
  agent.OneShotRun("Collection Name: 'Hellper' Query: How does embeddings package works?")
}

  /*
  completion_test, err := model.GenerateContent(context.Background(),intialState)
  if err != nil {
	log.Println("error with simple generate content",err)
  }
  log.Println("completion test: ", completion_test)
  */