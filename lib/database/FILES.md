# lib/database/newUserDataBase.go  
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
  
# lib/database/user.go  
## Package: database  
  
### Imports:  
  
- log  
- os  
- e (github.com/JackBekket/hellper/lib/embeddings)  
- godotenv  
- pgvector (github.com/tmc/langchaingo/vectorstores/pgvector)  
  
### External Data, Input Sources:  
  
- os.Getenv("AI_ENDPOINT")  
- os.Getenv("EMBEDDINGS_DB_URL")  
- u.AiSession.GptKey  
  
### TODOs:  
  
- None  
  
### Code Summary:  
  
#### SetContext Function:  
  
This function sets the context for a User object by initializing a vector store. It first loads environment variables using godotenv.Load(). Then, it retrieves the AI endpoint, API token, and embeddings database URL from the environment variables.  
  
Next, it creates a vector store using the provided information and the GetVectorStoreWithOptions function from the embeddings package. The vector store is then assigned to the User object's VectorStore field.  
  
Finally, the function uses a defer statement to ensure that the vector store is closed when the function exits. This is done by checking if the vector store implements the pgvector.Store interface and calling its Close() method.  
  
#### ClearContext Function:  
  
This function clears the context for a User object by setting its VectorStore field to nil.  
  
