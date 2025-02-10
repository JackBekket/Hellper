# Database Package
The database package provides functionality for managing user data and AI sessions. It includes data structures for representing users, sessions, and chat session graphs, as well as functions for adding users, updating usage, and creating new chat session graphs.

## Environment Variables
* `AI_ENDPOINT`
* `EMBEDDINGS_DB_URL`
* `GptKey` (api token)

## Flags and Cmdline Arguments
None

## Files and Paths
* `newUserDataBase.go`
* `user.go`
* `lib/database/newUserDataBase.go`
* `lib/database/user.go`

## Edge Cases for Launching the Application
The application can be launched as a command-line interface (CLI) using the following edge cases:
* Running the `AddUser` function to add a new user
* Running the `UpdateUserUsage` function to update a user's AI session usage
* Running the `NewChatSessionGraph` function to create a new chat session graph

## Project Package Structure
* `database/`
	+ `newUserDataBase.go`
	+ `user.go`
* `lib/database/`
	+ `newUserDataBase.go`
	+ `user.go`

## Relations between Code Entities
The `User` struct is related to the `AiSession` struct through the `AI session data` attribute. The `SessionUsage` struct is related to the `AiSession` struct through the `usage map` attribute. The `ChatSessionGraph` struct is related to the `User` struct through the `conversation buffer` attribute.

## Unclear Places or Dead Code
The `ClearContext` function seems to be unclear in its purpose, as it only sets the `VectorStore` to nil without any additional functionality.

