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
  
### Database Structure  
  
The package provides a basic structure for storing and managing user data and session information. It defines several types, including `User`, `SessionUsage`, and `AiSession`, which are used to represent user information, session usage, and AI session details, respectively.  
  
### User Management  
  
The package provides functions for adding and updating user information, including `AddUser` and `UpdateUserUsage`. These functions allow for storing and modifying user data, including AI session usage.  
  
### Session Management  
  
The package also provides functions for managing session information, including `UpdateSessionUsage` and `GetSessionUsage`. These functions allow for updating and retrieving session usage data.  
  
### Chat Session Graph  
  
The package defines a `ChatSessionGraph` type, which is used to represent a graph of chat sessions. It also provides a function `NewChatSessionGraph` for creating a new instance of this type.  
  
### Missing Function  
  
The package is missing a function `NewChatSession` that is expected to be implemented. This function is currently commented out.  
  
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
  
