# database

## Summary

This code package provides a database for managing AI sessions. It defines a struct called `AiSession` that stores information about each session, including the GPT key, GPT client, and GPT model. A map called `AiSessionMap` is used to store these sessions, with an integer as the key and the `AiSession` struct as the value. This package is designed to help manage and track AI sessions efficiently.



## Summary

This code package provides a database for managing user data and session usage in a chat application. It defines several data structures, including User, AiSession, and ChatSession, to store information about users, their AI sessions, and chat threads. The package also includes functions for adding users, updating user usage, and retrieving session usage.

The User struct stores information about a user, such as their Telegram ID, username, dialog status, admin status, AI session, and network. The AiSession struct stores details about the user's AI session, including the GPT key, GPT model, AI type, dialog thread, base URL, and usage. The ChatSession struct contains the conversation buffer and dialog thread for the user's chat session.

The package uses two maps to store user and session usage data: UsersMap and UsageMap. The AddUser function adds a new user to the UsersMap, while UpdateUserUsage and UpdateSessionUsage functions update the usage information for a given user or session. The GetSessionUsage function retrieves the usage information for a specific session.

In summary, this code package provides a database for managing user data and session usage in a chat application, allowing for efficient storage and retrieval of user information and session usage statistics.



