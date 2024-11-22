# Package: main

### Imports:

* context
* log
* os
* strconv
* github.com/JackBekket/hellper/lib/bot/command
* github.com/JackBekket/hellper/lib/bot/dialog
* github.com/JackBekket/hellper/lib/bot/env
* github.com/JackBekket/hellper/lib/database
* github.com/go-telegram-bot-api/telegram-bot-api/v5
* github.com/joho/godotenv

### External Data, Input Sources:

* OPENAI_API_KEY (local key for localai)
* PG_LINK (not used in the code)
* TG_KEY (telegram bot token)
* ADMIN_ID (admin user ID)
* AI_ENDPOINT (endpoint for AI model)

### Summary:

#### Initialization:

1. Loads environment variables using godotenv.Load().
2. Retrieves the telegram bot token from the TG_KEY environment variable.
3. Retrieves the admin user ID from the ADMIN_ID environment variable and parses it as an integer.
4. Retrieves the AI endpoint from the AI_ENDPOINT environment variable.
5. Creates a new Telegram bot instance using the retrieved token.

#### Database and Commander:

1. Initializes a database for storing user data using the database.UsersMap variable.
2. Creates a new command commander instance using the bot, database, and a context.

#### Update Handling:

1. Sets up a channel for handling incoming updates from the Telegram bot.
2. Starts a goroutine to handle updates using the dialog.HandleUpdates function.
3. Continuously listens for updates from the bot and checks if the user is new. If the user is new, the entry in the database is created.

#### Inline Keyboards:

1. Handles inline keyboards by checking if the update contains a callback query.
2. If a callback query is present, it retrieves the chat ID from the message.
3. If the user is new, the update is sent to the update channel for processing.

#### Edge Cases:

1. If the TG_KEY environment variable is not set, the bot will not be able to connect to Telegram.
2. If the ADMIN_ID environment variable is not set or is not a valid integer, the bot will not be able to identify the admin user.
3. If the AI_ENDPOINT environment variable is not set or is not a valid URL, the bot will not be able to access the AI model.

#### File Structure:

```
main.go
```

