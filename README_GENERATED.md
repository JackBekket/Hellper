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
* TG_KEY (Telegram bot token)
* ADMIN_ID (admin user ID)
* AI_ENDPOINT (AI endpoint)

### Summary:

#### Initialization:

1. Loads environment variables using `godotenv.Load()`.
2. Retrieves the Telegram bot token from the environment variable `TG_KEY`.
3. Retrieves the admin user ID from the environment variable `ADMIN_ID` and parses it as an integer.
4. Retrieves the AI endpoint from the environment variable `AI_ENDPOINT`.

#### Bot Initialization:

1. Creates a new Telegram bot instance using the retrieved token.
2. Creates a map of admin data, including the admin ID and their local AI key.

#### Database and Commander Initialization:

1. Initializes the database using `database.UsersMap`.
2. Creates a new command commander using the bot, database, and a context.

#### Update Handling:

1. Sets up a channel to handle incoming updates from the Telegram bot.
2. Starts a goroutine to handle updates using the `dialog.HandleUpdates` function.
3. Iterates through incoming updates and checks if the user is new. If so, adds the user to the database.

#### Inline Keyboard Logic:

1. Handles inline keyboard interactions by checking for callback queries.
2. Retrieves the chat ID from the update and checks if the user is already in the database.
3. If the user is new, adds them to the database and sends the update to the update channel.

#### End of main function:

1. The main function ends, and the program continues to handle updates and inline keyboard interactions.

