# lib/agent/duck_search_agent.go  
**Package/Component Name:** `agent`  
  
**Imports:**  
  
* `context`  
* `encoding/json`  
* `log`  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/llms/openai`  
* `github.com/tmc/langchaingo/tools/duckduckgo`  
* `github.com/JackBekket/langgraphgo/graph`  
  
**External Data/Inputs:**  
  
* None  
  
**TODO Comments:**  
  
* None  
  
**Summary:**  
  
### Overview  
The `agent` package provides a workflow-based chatbot agent that uses OpenAI's GPT-4o model to generate responses. The agent is designed to interact with a user and perform a search using the Duck Duck Go search engine.  
  
### Workflow  
The package defines a graph-based workflow that consists of two nodes: `agent` and `search`. The `agent` node is responsible for generating a response based on the input provided by the user, while the `search` node performs a search using the Duck Duck Go API. The workflow is compiled and executed to generate a response.  
  
### Code Breakdown  
  
#### Agent Node  
The `agent` node is responsible for generating a response based on the input provided by the user. It uses the OpenAI model to generate a response and appends the response to the state.  
  
#### Search Node  
The `search` node performs a search using the Duck Duck Go API. It takes the query as an input and returns the search results.  
  
#### Condition Function  
The `shouldSearch` function is used to determine whether the `search` node should be executed. It checks if the last message part is a tool call with the name "search" and returns "search" if true.  
  
### Conclusion  
The `agent` package provides a basic chatbot agent that can interact with a user and perform a search using the Duck Duck Go search engine. The workflow-based architecture allows for easy extension and modification of the agent's behavior.  
  
**  
  
# lib/agent/semantic_search_agent.go  
**Package/Component Name:** `agent`  
  
**Imports:**  
  
* `context`  
* `encoding/json`  
* `fmt`  
* `log`  
* `os`  
* `github.com/joho/godotenv`  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/llms/openai`  
* `github.com/JackBekket/hellper/lib/embeddings`  
* `github.com/JackBekket/langgraphgo/graph`  
  
**External Data/Inputs:**  
  
* `prompt` (string)  
* `model` (openai.LLM)  
* `history_state` (optional, []llms.MessageContent)  
  
**TODO Comments:**  
  
* `TODO: there should NOT exist arguments which called NAME cause it cause COLLISION with actual function name.`  
* `TODO: ALWAYS CHECK THIS JSON REFERENCE WHEN ALTERING VARS`  
  
**Summary:**  
  
The `agent` package is a OneShot agent example that does not have memory by itself, but can be passed a history of previous messages as an optional parameter. The main logic of this package is a workflow graph, which ensures that the message stack is processed as intended. The graph has multiple nodes, including an "agent" node, "semanticSearch" node, and conditional edges.  
  
The `OneShotRun` function is the main entry point of the package. It takes a prompt, a model, and an optional history state as input. It initializes a graph and sets the entry point to the "agent" node. The graph is then compiled and invoked, and the result is returned.  
  
The `agent` function is responsible for deciding whether to call a tool (e.g., `semanticSearch`) based on the user's input. If the decision is to call the tool, it appends the tool call to the message state and returns the updated state.  
  
The `shouldSearchDocuments` function is a handler for tool calls, specifically for the `semanticSearch` tool. It retrieves a vector store based on the input arguments and performs a semantic search. The results are then formatted and appended to the message state.  
  
The `semanticSearch` function is a real implementation of the `semanticSearch` tool, which performs a search in a vector store and returns the results.  
  
**Code Part Summaries:**  
  
### Main Workflow  
The main workflow is implemented in the `OneShotRun` function, which initializes a graph, compiles it, and invokes it. The graph is used to process the message stack as intended.  
  
### Agent Node  
The `agent` function is responsible for deciding whether to call a tool (e.g., `semanticSearch`) based on the user's input. If the decision is to call the tool, it appends the tool call to the message state and returns the updated state.  
  
### Tool Handlers  
The `shouldSearchDocuments` function is a handler for tool calls, specifically for the `semanticSearch` tool. It retrieves a vector store based on the input arguments and performs a semantic search. The results are then formatted and appended to the message state.  
  
### Tool Implementation  
The `semanticSearch` function is a real implementation of the `semanticSearch` tool, which performs a search in a vector store and returns the results.  
  
**  
  
# lib/agent/superagent.go  
**Package Name:** agent  
  
**Imports:**  
  
* `log`  
* `os`  
* `github.com/joho/godotenv`  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/llms/openai`  
  
**External Data/Inputs:**  
  
* `prompt` string (input to `OnePunch` and `RunThread` functions)  
* `model_name` string (settable, used in `CreateGenericLLM` function)  
* Environment variables:  
	+ `AI_ENDPOINT` (used in `CreateGenericLLM` function)  
	+ `API_TOKEN` (used in `CreateGenericLLM` function)  
	+ `EMBEDDINGS_DB_URL` (not used in the provided code)  
  
**TODOs:**  
  
* `// TODO: should be global?` (in `CreateGenericLLM` function)  
* `// TODO: should be settable?` (in `CreateGenericLLM` function)  
  
**Summary:**  
  
### Overview  
  
The provided code is part of the `agent` package and appears to be a prototype for a superagent that works with memory and has similar functionality to LangChain's chains. The code defines several functions for interacting with a language model.  
  
### Functions  
  
#### `OnePunch`  
  
This function takes a `prompt` string as input and calls the `OneShotRun` function with a generic LLM (created by `CreateGenericLLM`) and logs the result.  
  
#### `RunThread`  
  
This function takes a `prompt` string, a `model` (an `openai.LLM` instance), and an optional `history` of `llms.MessageContent` as input. It runs the `OneShotRun` function with the provided model and history, and returns the updated history and the result of the run.  
  
#### `CreateMessageContentAi` and `CreateMessageContentHuman`  
  
These functions create initial message content for AI and human inputs, respectively.  
  
#### `CreateGenericLLM`  
  
This function creates a generic LLM instance by loading environment variables and creating an `openai.LLM` instance. It returns the LLM instance.  
  
**  
  
