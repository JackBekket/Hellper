# Package: dialog

### Imports:

- log
- os
- regexp
- strings
- github.com/JackBekket/hellper/lib/bot/command
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External Data, Input Sources:

- Updates from Telegram bot API
- AI endpoint (AI_ENDPOINT environment variable)

### TODOs:

- Implement callback logic for inlines (in the else block of the main function)

### Summary:

The `dialog` package handles updates from the Telegram bot API and manages user dialogs. It uses a command interface (`command.Commander`) to interact with other components, such as a database for storing user data and an AI endpoint for processing user input.

The `HandleUpdates` function iterates over incoming updates and processes them based on the user's dialog status. It first checks if the update is a callback query, and if so, it handles the callback logic for inlines. Otherwise, it processes the update as a regular message.

For each message, the package extracts the chat ID, retrieves the corresponding user data from the database, and determines the appropriate dialog status. Based on the dialog status, the package either chooses a model, handles a wrong response, or initiates a dialog sequence with the AI endpoint.

The package also handles group chats and voice messages, ensuring that the bot only responds to messages that are relevant to the bot's purpose. It also removes mentions of the bot from messages to avoid unnecessary responses.

dialog directory file structure:
- dialog.go
 lib/bot/dialog/dialog.go

