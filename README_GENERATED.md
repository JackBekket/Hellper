# hellper/main

This package provides the main entry point for the application, which is a Telegram bot. It initializes the bot, database, and commander, and handles incoming updates from the Telegram API.

## Environment Variables

- OPENAI_API_KEY
- PG_LINK
- TG_KEY
- ADMIN_ID
- AI_ENDPOINT

## File Structure

- main.go

## Code Summary

1. Initialization:
   - Loads environment variables using godotenv.Load().
   - Retrieves OPENAI_API_KEY, TG_KEY, ADMIN_ID, and AI_ENDPOINT from environment variables.
   - Creates a new instance of the Telegram bot API using the retrieved TG_KEY.
   - Creates a map of admin data with ADMIN_ID as the key and env.AdminData as the value.
   - Initializes the database and commander.

2. Update Handling:
   - Logs the authorized account.
   - Creates a new update object with a timeout of 60 seconds.
   - Creates a channel for handling updates.
   - Starts a goroutine to handle updates using the dialog.HandleUpdates function.
   - Iterates over incoming updates and checks if the user ID is in the database.
   - If the user ID is not in the database, it is added to the database and the update is sent to the update channel.
   - If the user ID is already in the database, the update is sent to the update channel.

3. Command Handling:
   - The code uses the github.com/JackBekket/hellper/lib/bot/command package to handle commands from users.
   - The commander is initialized and used to process incoming commands.

4. Dialog Handling:
   - The code uses the github.com/JackBekket/hellper/lib/bot/dialog package to handle dialogs with users.
   - The dialog.HandleUpdates function is used to process incoming updates and handle dialogs.

5. Database Interaction:
   - The code uses the github.com/JackBekket/hellper/lib/database package to interact with the database.
   - The database is initialized and used to store user data and other relevant information.

6. AI Integration:
   - The code uses the AI_ENDPOINT environment variable to connect to an external AI service.
   - The OPENAI_API_KEY environment variable is used to authenticate with the AI service.

7. Logging:
   - The code uses the log package to log messages and events.
   - The log level can be configured using environment variables or command-line arguments.

8. Error Handling:
   - The code includes basic error handling using the context package.
   - Errors are logged and handled appropriately.

9. Dead Code:
   - There is no apparent dead code in the provided code.

10. Edge Cases:
   - The application can be launched with the following command-line arguments:
     - OPENAI_API_KEY
     - PG_LINK
     - TG_KEY
     - ADMIN_ID
     - AI_ENDPOINT

