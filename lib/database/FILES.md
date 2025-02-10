# lib/database/newUserDataBase.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/vectorstores`  
  
## External Data and Input Sources  
The external data and input sources are not explicitly defined in the provided code. However, it can be inferred that the data is related to user interactions, such as:  
* User ID (int64)  
* Username (string)  
* Dialog status (int8)  
* Admin status (bool)  
* AI session data (AiSession struct)  
* Network (string)  
* Topics (array of int)  
  
## TODO Comments  
The list of TODO comments is:  
* User should be a fully functional user class and all operations with the user should be placed in a separate user.go package.  
  
## Code Summary  
### Data Structures  
The code defines several data structures, including:  
* `User` struct: represents a user with attributes such as ID, username, dialog status, admin status, AI session data, network, and topics.  
* `SessionUsage` struct: represents the usage of a session with attributes such as ID and usage map.  
* `AiSession` struct: represents an AI session with attributes such as GPT key, GPT model, AI type, dialog thread, base URL, and usage map.  
* `ChatSessionGraph` struct: represents a chat session graph with attributes such as conversation buffer.  
  
### Functions  
The code defines several functions, including:  
* `AddUser`: adds a user to the `UsersMap`.  
* `UpdateUserUsage`: updates the usage of a user's AI session.  
* `UpdateSessionUsage`: updates the usage of a session.  
* `GetSessionUsage`: returns the usage of a session.  
* `NewChatSessionGraph`: creates a new chat session graph.  
  
### Variables  
The code defines two variables:  
* `UsersMap`: a map of user IDs to `User` structs.  
* `UsageMap`: a map of session IDs to `SessionUsage` structs.  
  
  
  
# lib/database/user.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* "log"  
* "os"  
* "github.com/JackBekket/hellper/lib/embeddings" (as e)  
* "github.com/joho/godotenv"  
* "github.com/tmc/langchaingo/vectorstores/pgvector"  
  
## External Data and Input Sources  
The external data and input sources are:  
* Environment variables:  
	+ "AI_ENDPOINT"  
	+ "EMBEDDINGS_DB_URL"  
* AI session data:  
	+ "GptKey" (api token)  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### SetContext Function  
The `SetContext` function sets the context for a user by loading environment variables, creating a vector store, and assigning it to the user. It takes a `collectionName` string as input and returns an error.  
  
### ClearContext Function  
The `ClearContext` function clears the context for a user by setting the `VectorStore` to nil.  
  
  
  
