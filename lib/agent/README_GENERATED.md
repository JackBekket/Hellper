# Package: agent_test

### Imports:
- log
- testing
- github.com/JackBekket/hellper/lib/agent
- github.com/tmc/langchaingo/llms

### External Data, Input Sources:
- Collection Name: 'Hellper'
- Query: How does embeddings package works?
- initialstate: []llms.MessageContent

### Test_Search:
This function tests the autonomous semantic_search agent. It runs the agent with the given query and logs the result.

### TestMemory:
This function tests the agent's memory by providing an initial state containing a conversation between a user and the AI assistant. The agent is then run with the query "How does embeddings package works? Also do you remember what is my name?" and the initial state. The result is logged, and the expected output should be compared to the actual result.

