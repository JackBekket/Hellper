# agent

- duck_search_agent.go
- semantic_search_agent.go
- semantic_search_test.go
- superagent.go
- lib/agent/semantic_search_test.go

The `agent` package provides functionality for interacting with a large language model (LLM) and performing semantic search. It includes tests for the `agent` package, which seems to be a component for interacting with a large language model (LLM) and performing semantic search.

The tests include:

1. `Test_Search`: Tests the one-shot semantic search agent by calling the `OneShotRun` function with a query and a generic LLM.

2. `TestMemory`: Tests the agent with memory by providing an initial state and calling `OneShotRun` with a query.

3. `TestLongConversation`: Tests a longer conversation by calling `OneShotRun` multiple times with different queries and appending the responses to the initial state.

4. `Test5Conversation`: Similar to `TestLongConversation`, but with 5 turns of conversation.

5. `createGenericLLM`: Creates a generic LLM using the OpenAI API and sets up the necessary parameters, such as the model name, API token, and base URL.

The tests use the `OneShotRun` function to perform semantic search and retrieve information from a collection called "Hellper". The results are then compared to expected outputs or assertions.

The `agent` package also includes a `superagent` file, which may contain additional functionality related to the LLM and semantic search.

The tests are located in the `agent_test` package, which is likely used for testing the functionality of the `agent` package.

The `lib/agent/semantic_search_test.go` file contains additional tests for the semantic search functionality of the `agent` package.

The `duck_search_agent.go` file contains the implementation of a duck search agent, which may be used for performing semantic search on a collection of documents.

The `semantic_search_agent.go` file contains the implementation of a semantic search agent, which may be used for performing semantic search on a collection of documents.

The `semantic_search_test.go` file contains tests for the semantic search functionality of the `agent` package.

The `agent` package provides a comprehensive set of tools and functionalities for interacting with an LLM and performing semantic search. The tests included in the package ensure that the functionality is working as expected and can be used to retrieve information from a collection called "Hellper".

