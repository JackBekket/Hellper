# Database Package
## Summary
The database package is responsible for managing the interaction between the application and the PostgreSQL database. It provides functions for creating tables, updating history and usage statistics, and handling user sessions.

## Environment Variables
* `AI_ENDPOINT`
* `EMBEDDINGS_DB_URL`

## Flags/Cmdline Arguments
None

## Files and Paths
* `database.go`
* `init.go`
* `models.go`
* `newUserDataBase.go`
* `service.go`
* `user.go`
* `lib/database/database.go`

## Edge Cases
The application can be launched in the following ways:
* By running the `NewHandler` function in `init.go` to create a new database connection
* By calling the `CreateTables` function in `database.go` to create the database tables
* By running the `NewAIService` function in `service.go` to create a new instance of the `Service` struct

## Project Package Structure
* `database`
	+ `database.go`
	+ `init.go`
	+ `models.go`
	+ `newUserDataBase.go`
	+ `service.go`
	+ `user.go`

## Relations Between Code Entities
The `Handler` struct in `init.go` is used to create a new database connection, which is then used by the functions in `database.go` to interact with the database. The `Service` struct in `service.go` is used to manage the user sessions and language models. The `User` struct in `user.go` is used to represent a user and their context.

## Unclear Places/Dead Code
The `FlushThread` method in `user.go` is currently empty and may need to be implemented. The `TODO` comments in `database.go` indicate that some functions may need to be debugged or completed.

