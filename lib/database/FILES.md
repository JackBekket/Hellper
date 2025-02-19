# lib/database/database.go  
# Package/Component Name  
The package/component name is **database**.  
  
## Imports  
The imports in this package are:  
* `"database/sql"`  
* `"fmt"`  
* `"log"`  
* `"github.com/tmc/langchaingo/llms"`  
  
## External Data/Input Sources  
The external data/input sources in this package are:  
* Database ( PostgreSQL )  
* `llms` package for language model related functionality  
  
## TODO Comments  
The TODO comments in this package are:  
* `TODO: debug` in the `DropHistory` function  
* `TODO: this variable is not set anywhere in user object, so it equals to 0. In our db we use 1, that's why query fails.` in the `DropHistory` function  
* `also create new session?` in the `UpdateEndpoint` function  
  
## Summary of Major Code Parts  
### Database Schema  
The database schema consists of several tables:  
* `auth_methods`  
* `endpoints`  
* `auth`  
* `ai_sessions`  
* `chat_sessions`  
* `last_usage`  
* `usage`  
* `chat_messages`  
  
### Functions  
The major functions in this package are:  
#### CreateTables  
Creates the database tables if they do not exist.  
#### UpdateHistory  
Updates the chat history for a given user and endpoint.  
#### UpdateUsage  
Updates the usage statistics for a given user and endpoint.  
#### DropUsage  
Drops the usage statistics for a given user and endpoint.  
#### GetUsage  
Gets the usage statistics for a given user and endpoint.  
#### GetHistory  
Gets the chat history for a given user and endpoint.  
#### DropHistory  
Drops the chat history for a given user and endpoint.  
#### UpdateModel  
Updates the model for a given user.  
#### UpdateEndpoint  
Updates the endpoint for a given user.  
#### GetSession  
Gets the session for a given user.  
#### CheckSession  
Checks if a session exists for a given user.  
#### CreateLSession  
Creates a new session for a given user.  
#### DeleteLSession  
Deletes a session for a given user.  
#### GetToken  
Gets the token for a given user and auth method.  
#### CheckToken  
Checks if a token exists for a given user and auth method.  
#### InsertToken  
Inserts a token for a given user and auth method.  
#### DeleteToken  
Deletes a token for a given user and auth method.  
#### GetEndpoints  
Gets the endpoints.  
  
  
  
# lib/database/init.go  
# Package: database  
## Imports  
The package imports the following modules:  
* `database/sql`  
* `github.com/lib/pq` (as a side-effect import, likely for PostgreSQL driver registration)  
  
## External Data and Input Sources  
The package uses the following external data and input sources:  
* A PostgreSQL database connection string, passed to the `NewHandler` function  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Handler Struct  
The package defines a `Handler` struct with a single field `DB` of type `*sql.DB`, representing a database connection.  
  
### NewHandler Function  
The `NewHandler` function creates a new instance of the `Handler` struct, taking a PostgreSQL connection string as input. It:  
1. Opens a database connection using the provided connection string.  
2. Returns a pointer to the `Handler` instance, or an error if the connection fails.  
  
  
  
# lib/database/models.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* "encoding/json"  
* "net/http"  
* "net/url"  
  
## External Data and Input Sources  
The external data and input sources are:  
* **endpoint**: a string representing the endpoint URL  
* **token**: a string representing the authorization token  
* **http.DefaultClient**: the default HTTP client used to send requests  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Data Structures  
The code defines two data structures:  
* **OpenAIDataObject**: a struct representing an OpenAI data object with **ID** and **Object** fields  
* **OpenAIModelsResponse**: a struct representing an OpenAI models response with a **Data** field containing a slice of **OpenAIDataObject** instances  
  
### Function GetModelsList  
The **GetModelsList** function:  
* takes **endpoint** and **token** as input parameters  
* constructs a URL path by joining the **endpoint** with the "models" path  
* sends a GET request to the constructed URL with the **token** in the Authorization header  
* decodes the response into an **OpenAIModelsResponse** instance  
* extracts the model IDs from the response and returns them as a slice of strings  
  
  
  
# lib/database/newUserDataBase.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/vectorstores`  
  
## External Data and Input Sources  
The external data and input sources are not explicitly defined in the provided code. However, the types and functions suggest that the data may come from:  
* Telegram user IDs  
* AI session usage data  
* Vector store data  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### User and Session Types  
The code defines several types:  
* `User`: represents a user with an ID, username, dialog status, admin status, AI session, network, topics, and vector store.  
* `SessionUsage`: represents the usage of an AI session with an ID and a map of usage data.  
* `AiSession`: represents an AI session with a GPT key, model, type, dialog thread, base URL, and usage data.  
* `ChatSessionGraph`: represents a chat session graph with a conversation buffer.  
  
### Functions  
The code defines several functions:  
* `AddUser`: adds a user to the `UsersMap`.  
* `UpdateUserUsage`: updates the usage data for a user in the `UsersMap`.  
* `UpdateSessionUsage`: updates the usage data for a session in the `UsageMap`.  
* `GetSessionUsage`: retrieves the usage data for a session from the `UsageMap`.  
* `NewChatSessionGraph`: creates a new chat session graph with a given conversation buffer.  
* `ClearBuffer`: clears the conversation buffer of a chat session graph.  
  
### Variables  
The code defines two variables:  
* `UsersMap`: a map of users with their IDs as keys.  
* `UsageMap`: a map of session usage data with their IDs as keys.  
  
  
  
# lib/database/service.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* `errors`  
* `sync`  
* `github.com/tmc/langchaingo/llms/openai`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `userId` of type `int64`  
* `token`, `model`, and `endpointURL` of type `string`  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Service Struct  
The `Service` struct has two fields: `LLMHandlers` of type `sync.Map` and `DBHandler` of type `*Handler`.  
  
### NewAIService Function  
The `NewAIService` function creates a new instance of the `Service` struct and calls the `CreateTables` method.  
  
### GetHandler Function  
The `GetHandler` function retrieves a handler for a given `userId` from the `LLMHandlers` map.  
  
### DropHandler Function  
The `DropHandler` function removes a handler for a given `userId` from the `LLMHandlers` map.  
  
### UpdateHandler Function  
The `UpdateHandler` function updates a handler for a given `userId` with new `token`, `model`, and `endpointURL` values.  
  
  
  
# lib/database/user.go  
# Package Name and Imports  
The package name is **database**. The imports are:  
* `fmt`  
* `log`  
* `os`  
* `github.com/JackBekket/hellper/lib/embeddings` (as `e`)  
* `github.com/joho/godotenv`  
* `github.com/tmc/langchaingo/vectorstores/pgvector`  
  
## External Data and Input Sources  
The external data and input sources are:  
* Environment variables:  
	+ `AI_ENDPOINT`  
	+ `EMBEDDINGS_DB_URL`  
* `godotenv` library to load environment variables from a .env file  
* `pgvector` library to interact with a PostgreSQL database  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### User Struct Methods  
The `User` struct has several methods:  
* `SetContext`: sets the context for the user by loading environment variables, creating a vector store, and storing it in the `User` struct  
* `ClearContext`: clears the context for the user by setting the `VectorStore` to `nil`  
* `FlushThread`: currently empty, but intended to flush the thread  
* `FlushMemory`: drops the history for the user from the database  
* `Kill`: kills the user session by flushing the thread, flushing the memory, deleting the token, deleting the language session, and removing the user from the `UsersMap`  
* `DropSession`: drops the language session for the user and flushes the thread and memory  
  
### Vector Store  
The vector store is created using the `e.GetVectorStoreWithOptions` function, which takes the AI endpoint, API token, database link, and collection name as options.  
  
### Error Handling  
Error handling is done using `log.Println` statements to print error messages.  
  
  
  
