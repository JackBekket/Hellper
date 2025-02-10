## Package: database

### Imports:
- github.com/tmc/langchaingo/llms
- github.com/tmc/langchaingo/vectorstores

### External Data, Input Sources:
- UsersMap: A map that stores user data, where the key is the telegram user ID and the value is a User struct.
- UsageMap: A map that stores session usage data, where the key is the telegram user ID and the value is a SessionUsage struct.

### TODOs:
- User should be fully functional user class and all operation with user should be placed here (in separate user.go package)

### Summary:
The database package manages user data and session usage. It includes structs for User, SessionUsage, and ChatSessionGraph. The package also provides functions for adding users, updating user usage, updating session usage, and getting session usage.

#### User Struct:
The User struct represents a user in the system. It contains the user's ID, username, dialog status, admin status, AI session, network, topics, and vector store.

#### SessionUsage Struct:
The SessionUsage struct represents the usage of an AI session. It contains the session ID and a map of usage data.

#### ChatSessionGraph Struct:
The ChatSessionGraph struct represents a chat session graph. It contains a conversation buffer of messages.

#### Functions:
- AddUser: Adds a new user to the UsersMap.
- UpdateUserUsage: Updates the usage data for a user in the UsersMap.
- UpdateSessionUsage: Updates the usage data for a session in the UsageMap.
- GetSessionUsage: Retrieves the usage data for a session from the UsageMap.
- NewChatSessionGraph: Creates a new ChatSessionGraph with a given buffer of messages.

lib/database/newUserDataBase.go
- newUserDataBase.go
- user.go

