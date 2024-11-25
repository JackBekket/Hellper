## Package: database

### Imports:

- github.com/tmc/langchaingo/chains
- github.com/tmc/langchaingo/memory
- github.com/tmc/langchaingo/vectorstores

### External Data, Input Sources:

- UsersMap: A map that stores user data, where the key is the user's ID and the value is a User struct.
- UsageMap: A map that stores session usage data, where the key is the user's ID and the value is a SessionUsage struct.

### Code Summary:

#### User Struct:

The User struct represents a user in the database. It contains the user's ID, username, dialog status, admin status, AI session, network, topics, and a vector store.

#### SessionUsage Struct:

The SessionUsage struct represents the usage of an AI session. It contains the session ID and a map of usage data.

#### AiSession Struct:

The AiSession struct represents an AI session. It contains the GPT key, GPT model, AI type, chat session, base URL, and usage data.

#### ChatSession Struct:

The ChatSession struct represents a chat session. It contains a conversation buffer and a dialog thread.

#### Functions:

- AddUser: Adds a new user to the UsersMap.
- UpdateUserUsage: Updates the usage data for a user's AI session.
- UpdateSessionUsage: Updates the usage data for a session.
- GetSessionUsage: Retrieves the usage data for a session.
- NewChatSession: Creates a new chat session with a conversation buffer and dialog thread.

The code provides a basic framework for managing user data, AI sessions, and chat sessions. It includes functions for adding users, updating usage data, and creating new chat sessions.

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

