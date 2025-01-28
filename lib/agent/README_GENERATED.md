# agent

This package contains code for an agent that can perform semantic search and maintain a conversation history. The agent uses a generic LLM to process queries and retrieve information from a database.

### Imports

```
import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/JackBekket/hellper/lib/agent"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)
```

### External Data, Input Sources

- Environment variables: `AI_ENDPOINT`, `ADNIN_KEY`, `EMBEDDINGS_DB_URL`

### Code Summary

#### Test_Search

This test case focuses on the semantic search functionality of the agent. It creates a generic LLM using the `createGenericLLM` function and then calls the `OneShotRun` function with a query to search for information on how the embeddings package works. The result is then printed to the console.

#### TestMemory

This test case tests the agent's ability to maintain a conversation history. It initializes a conversation state with a few messages and then calls the `OneShotRun` function with a query to search for information on how the main package works. The result is then printed to the console.

#### TestLongConversation

This test case simulates a longer conversation by appending the agent's response to the conversation state and then calling the `OneShotRun` function again with a new query. This process is repeated for multiple turns, and the results are printed to the console.

#### createGenericLLM

This function creates a generic LLM using the OpenAI API. It first loads environment variables for the API endpoint, API token, and database URL. Then, it creates an instance of the OpenAI LLM using the provided parameters and returns the LLM object.

File structure:

- duck_search_agent.go
- semantic_search_agent.go
- semantic_search_test.go
- superagent.go
- lib/agent/semantic_search_test.go

