Here is the markdown summary of the provided package code:

**database**
================

### Overview

The `database` package provides a basic structure for storing and managing user data and session information. It defines several types and functions for user and session management.

### Configuration

* Environment variables:
	+ `AI_ENDPOINT`
	+ `EMBEDDINGS_DB_URL`
* Command-line arguments:
	* None
* Files and paths:
	* None

### Launching the Application

The package can be launched as a command-line interface (CLI) or as a main package. To launch the CLI, run the command `go run main.go`. To launch the main package, run the command `go run main.go`.

### Edge Cases

* None

### Package Structure

```
database/
newUserDataBase.go
user.go
lib/
database/
newUserDataBase.go
user.go
```

### Summary

The package provides a basic structure for storing and managing user data and session information. It defines several types, including `User`, `SessionUsage`, and `AiSession`, which are used to represent user information, session usage, and AI session details, respectively.

The `user.go` file provides a `SetContext` function that sets the context for a `User` instance, and a `ClearContext` function that resets the `VectorStore` field of the `User` instance to `nil`.

The `newUserDataBase.go` file defines a `NewUserDataBase` function that is used to create a new user database.

**