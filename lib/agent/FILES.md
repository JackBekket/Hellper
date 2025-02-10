# lib/agent/semantic_search_test.go  
# Package Name and Imports  
The package name is `agent_test`. The imports are:  
- `flag`  
- `fmt`  
- `log`  
- `os`  
- `testing`  
- `github.com/JackBekket/hellper/lib/agent`  
- `github.com/joho/godotenv`  
- `github.com/tmc/langchaingo/llms`  
- `github.com/tmc/langchaingo/llms/openai`  
  
## External Data and Input Sources  
The external data and input sources are:  
- Environment variables:  
  - `AI_ENDPOINT`  
  - `ADNIN_KEY`  
  - `EMBEDDINGS_DB_URL` (not used in this code snippet)  
- Model name: `tiger-gemma-9b-v1-i1` (should be settable)  
- Initial state for conversations:  
  - System message: "Below a current conversation between user and helpful AI assistant. Your task will be in the next system message"  
  - Human message: "Hello my name is Yemet! I'm excited to see what we can do together."  
  - AI messages: "Hey there! Let me know how I can help you out." and "I'm here to assist you with any task you may have. Just give me a command and I'll do my best to help."  
  
## TODO Comments  
The TODO comments are:  
- `// should be settable?` (model name)  
- `//TODO: should be global?` (ai_url)  
  
## Code Summary  
### Test Functions  
The code includes several test functions:  
- `Test_Search`: Tests the one-shot search functionality of the agent.  
- `TestMemory`: Tests the agent's memory by asking it to recall the user's name.  
- `TestLongConversation`: Tests a long conversation with the agent.  
- `Test5Conversation`: Tests a conversation with 5 turns.  
  
### createGenericLLM Function  
The `createGenericLLM` function creates a new instance of the `openai.LLM` model with the specified model name, API token, and base URL.  
  
### Agent Functionality  
The agent's functionality is tested through the `OneShotRun` function, which takes a query, a model, and an initial state as input and returns a result.  
  
  
  
