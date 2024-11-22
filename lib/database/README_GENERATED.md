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

lib/database/newUserDataBase.go
