# lib/agent/semantic_search_test.go  
## agent_test  
  
### Imports  
  
```  
flag  
fmt  
log  
os  
testing  
github.com/JackBekket/hellper/lib/agent  
github.com/joho/godotenv  
github.com/tmc/langchaingo/llms  
github.com/tmc/langchaingo/llms/openai  
```  
  
### External Data, Input Sources  
  
- Environment variables: AI_ENDPOINT, ADNIN_KEY, EMBEDDINGS_DB_URL  
  
### TODOs  
  
- Should be global? (referring to ai_url)  
- Settable? (referring to model_name)  
  
### Summary  
  
The `agent_test` package contains tests for the `agent` package, which seems to be a component for interacting with a large language model (LLM) and performing semantic search.  
  
The tests include:  
  
1. `Test_Search`: Tests the one-shot semantic search agent by calling the `OneShotRun` function with a query and a generic LLM.  
  
2. `TestMemory`: Tests the agent with memory by providing an initial state and calling `OneShotRun` with a query.  
  
3. `TestLongConversation`: Tests a longer conversation by calling `OneShotRun` multiple times with different queries and appending the responses to the initial state.  
  
4. `Test5Conversation`: Similar to `TestLongConversation`, but with 5 turns of conversation.  
  
5. `createGenericLLM`: Creates a generic LLM using the OpenAI API and sets up the necessary parameters, such as the model name, API token, and base URL.  
  
The tests use the `OneShotRun` function to perform semantic search and retrieve information from a collection called "Hellper". The results are then compared to expected outputs or assertions.  
  
