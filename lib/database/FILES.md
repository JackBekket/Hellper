# lib/database/newUserDataBase.go  
**Package Name:** database  
  
**Imports:**  
  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/vectorstores`  
  
**External Data/Inputs:**  
  
* None  
  
**TODO Comments:**  
  
* `// user should be fully functional user class and all operation with user should be placed here (in separate user.go package)`  
  
**Summary:**  
  
### Data Structures  
  
The package defines several data structures:  
  
* `User`: represents a user with various attributes (ID, username, dialog status, admin status, AI session, network, and topics)  
* `SessionUsage`: represents a session usage with a map of usage data  
* `AiSession`: represents an AI session with various attributes (GptKey, GptModel, AI type, dialog thread, and base URL)  
* `ChatSessionGraph`: represents a chat session graph with a conversation buffer  
  
### Functions  
  
The package provides several functions:  
  
* `AddUser`: adds a user to the `UsersMap`  
* `UpdateUserUsage`: updates the usage data for a user  
* `UpdateSessionUsage`: updates the session usage data  
* `GetSessionUsage`: retrieves the session usage data for a given ID  
* `NewChatSessionGraph`: creates a new `ChatSessionGraph` instance  
  
### Notes  
  
The package seems to be related to natural language processing and AI-powered chat sessions. The `User` struct is incomplete and is marked as a TODO, indicating that it should be fully functional and moved to a separate package.  
  
**  
  
# lib/database/user.go  
**Package Name:** database  
  
**Imports:**  
  
* `log`  
* `os`  
* `github.com/JackBekket/hellper/lib/embeddings` (e)  
* `github.com/joho/godotenv`  
* `github.com/tmc/langchaingo/vectorstores/pgvector`  
  
**External Data/Inputs:**  
  
* `collectionName` (string)  
* `os` environment variables:  
	+ `AI_ENDPOINT`  
	+ `EMBEDDINGS_DB_URL`  
* `u.AiSession.GptKey` (assuming `u` is an instance of `User`)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### SetContext Function  
  
The `SetContext` function sets the context for a `User` instance. It loads environment variables, retrieves a vector store using the `embeddings` package, and assigns it to the `VectorStore` field of the `User` instance. The function also closes the vector store using the `pgvector` package.  
  
### ClearContext Function  
  
The `ClearContext` function resets the `VectorStore` field of the `User` instance to `nil`.  
  
**  
  
