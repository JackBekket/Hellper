# database

This package provides a database for storing and managing user data, AI session data, and session usage data. It includes structs for representing users, AI sessions, and session usage, as well as functions for adding users, updating user and session usage, and retrieving session usage.

## File Structure

```
lib/database/
    newAiSessionDataBase.go
    newUserDataBase.go
```

## Major Code Parts

### User Data

The `User` struct represents a user in the database. It has fields for the user's ID, username, dialog status, admin status, AI session, and network. The `UsersMap` variable is a map that stores users, where the key is the user's ID and the value is the `User` struct.

The `AddUser` function adds a new user to the `UsersMap`.

### AI Session Data

The `AiSession` struct represents an AI session. It has fields for the GPT key, GPT model, AI type, dialog thread, base URL, and usage. The `AiSessionMap` variable is a map that stores AI sessions, where the key is an integer and the value is the `AiSession` struct.

### Session Usage Data

The `SessionUsage` struct represents the usage of a session. It has fields for the session ID and usage data. The `UsageMap` variable is a map that stores session usage, where the key is the session ID and the value is the `SessionUsage` struct.

The `UpdateSessionUsage` function updates the usage of a session with a given ID in the `UsageMap`. The `GetSessionUsage` function returns the usage of a session with a given ID from the `UsageMap`.

### Chat Session Data

The `ChatSession` struct represents a chat session. It has fields for the conversation buffer and dialog thread.

### Edge Cases

The package does not explicitly handle any edge cases for launching the application.

### Environment Variables, Flags, and Command-Line Arguments

The package does not use any environment variables, flags, or command-line arguments for configuration.

