## Package: database

### Imports:
- github.com/tmc/langchaingo/llms
- github.com/tmc/langchaingo/vectorstores

### External Data, Input Sources:
- UsersMap: A map that stores user data, where the key is the Telegram user ID and the value is a User struct.
- UsageMap: A map that stores session usage data, where the key is the session ID and the value is a SessionUsage struct.

### Major Code Parts:

#### User Struct:
- Represents a user in the database.
- Contains fields for user ID, username, dialog status, admin status, AI session, network, topics, and a vector store.

#### SessionUsage Struct:
- Represents a session's usage data.
- Contains fields for session ID and a map of usage statistics.

#### AiSession Struct:
- Represents an AI session.
- Contains fields for GPT key, GPT model, AI type, dialog thread, base URL, and usage statistics.

#### ChatSessionGraph Struct:
- Represents a chat session graph.
- Contains a field for the conversation buffer, which stores the messages in the chat session.

#### Functions:
- AddUser: Adds a new user to the UsersMap.
- UpdateUserUsage: Updates the usage statistics for a user's AI session.
- UpdateSessionUsage: Updates the usage statistics for a session.
- GetSessionUsage: Retrieves the usage statistics for a session.
- NewChatSessionGraph: Creates a new ChatSessionGraph with a given conversation buffer.

The code provides a basic framework for managing user data, AI sessions, and chat session graphs. It includes data structures and functions for adding users, updating usage statistics, and creating chat session graphs.

lib/database/user.go
## Package: database

### Imports:

- log
- os
- e (github.com/JackBekket/hellper/lib/embeddings)
- godotenv
- pgvector (github.com/tmc/langchaingo/vectorstores/pgvector)

### External Data and Input Sources:

- AI_ENDPOINT: Environment variable containing the API endpoint for the AI service.
- GPT_KEY: User's AI session GPT key.
- EMBEDDINGS_DB_URL: Environment variable containing the URL for the embeddings database.

### Code Summary:

#### SetContext Function:

This function sets the context for a user by establishing a connection to the embeddings database and initializing a vector store. It first loads the environment variables and retrieves the necessary information for connecting to the AI service and the embeddings database. Then, it uses the provided information to create a vector store using the `GetVectorStoreWithOptions` function from the `embeddings` package. The vector store is stored in the user's `VectorStore` field. Finally, it sets up a defer function to ensure that the vector store is closed when the function exits.

#### ClearContext Function:

This function clears the context for a user by setting the `VectorStore` field to nil. This effectively disconnects the user from the embeddings database and closes the vector store.

