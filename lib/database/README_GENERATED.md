# database

This package provides a database for storing and managing user data, including user sessions and their usage. It also includes a database for managing AI session data.

## File structure

```
lib/database/
    newUserDataBase.go
    newAiSessionDataBase.go
```

## User data

The `newUserDataBase.go` file defines a database for storing user data. It includes functions for adding new users, updating user usage, and retrieving session usage.

- `AddUser(user User)`: Adds a new user to the `UsersMap`.
- `UpdateUserUsage(id int64, usage map[string]int)`: Updates the usage of a user's AiSession.
- `UpdateSessionUsage(id int64, usage map[string]int)`: Updates the usage of a session.
- `GetSessionUsage(id int64) (map[string]int)`: Returns the usage of a session.

The `User` struct represents a user and includes fields for ID, username, dialog status, admin status, AiSession, and network.

- `AiSession`: Represents an AI session and includes fields for GPT key, GPT model, AI type, dialog thread, base URL, and usage.

The `UsersMap` is a map that stores users by their ID.

## AI session data

The `newAiSessionDataBase.go` file defines a database for storing AI session data. It includes an `AiSessionMap` that stores AI sessions by their ID.

- `main()`: Prints "Hello world!"

The `AiSession` struct represents an AI session and includes fields for GPT key, GPT model, AI type, dialog thread, base URL, and usage.

## Edge cases

There are no edge cases mentioned in the provided code.

## Conclusion

This package provides a database for storing and managing user data, including user sessions and their usage. It also includes a database for managing AI session data. The package includes functions for adding new users, updating user usage, and retrieving session usage. The package also includes a database for managing AI session data.