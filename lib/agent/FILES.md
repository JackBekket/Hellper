# lib/agent/semantic_search_test.go  
**Package/Component Name:** agent_test  
  
**Imports:**  
  
* `flag`  
* `fmt`  
* `log`  
* `os`  
* `testing`  
* `github.com/JackBekket/hellper/lib/agent`  
* `github.com/joho/godotenv`  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langgraphgo/graph` (commented out)  
  
**External Data/Inputs:**  
  
* `flag` is used to set a test timeout  
* `os` is used to load environment variables  
* `godotenv` is used to load environment variables  
* `openai` is used to interact with the OpenAI API  
* `llms` is used to create message content  
* `agent` is used to run the OneShotRun function  
  
**TODO Comments:**  
  
* `// TODO: should be global?` in the `createGenericLLM` function  
* `// db_link := os.Getenv("EMBEDDINGS_DB_URL")` in the `createGenericLLM` function (commented out)  
  
**Summary:**  
  
### Overview  
  
The `agent_test` package contains a set of tests for the agent's OneShotRun functionality. The tests simulate conversations with the agent, testing its ability to answer questions and remember previous conversations.  
  
### Test Functions  
  
* `Test_Search`: Tests the OneShotRun function with a single query  
* `TestMemory`: Tests the OneShotRun function with a conversation that spans multiple turns  
* `Test5Conversation`: Tests the OneShotRun function with a 5-turn conversation  
  
### Code Highlights  
  
* The `createGenericLLM` function is used to create a generic LLM model, which is then used in the tests  
* The tests use the `agent` package to run the OneShotRun function, which interacts with the OpenAI API  
* The tests also use the `llms` package to create message content and the `godotenv` package to load environment variables  
  
**  
  
