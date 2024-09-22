# Package: hellper/lib/bot

### Imports:

- context
- log
- os
- strconv
- github.com/JackBekket/hellper/lib/bot/command
- github.com/JackBekket/hellper/lib/bot/dialog
- github.com/JackBekket/hellper/lib/bot/env
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5
- github.com/joho/godotenv

### External Data, Input Sources:

- OPENAI_API_KEY (local key for localai)
- PG_LINK (not used in the code)
- TG_KEY (telegram bot token)
- ADMIN_ID (admin user ID)
- AI_ENDPOINT (local AI endpoint)

### Summary:

The package provides a Telegram bot that interacts with users and manages admin information. It initializes a Telegram bot instance using the provided TG_KEY, loads environment variables, and sets up a database for storing user information. The bot handles incoming updates from Telegram, checks if the user is new, and adds them to the database if necessary. It also uses a command.Commander to handle bot commands and user interactions.

#### Initialization:

1. Loads environment variables using godotenv.Load() and retrieves values for OPENAI_API_KEY, TG_KEY, ADMIN_ID, and AI_ENDPOINT.
2. Creates a new instance of the Telegram bot using the retrieved TG_KEY.
3. Initializes a map called adminData to store admin information, including the admin ID and GPT key.

#### Database and Commander:

1. Initializes a database for storing user information using the database.UsersMap.
2. Creates a new instance of the command.Commander, which handles bot commands and user interactions.

#### Update Handling:

1. Sets up a channel for handling incoming updates from the Telegram bot.
2. Starts a goroutine to handle updates using the dialog.HandleUpdates function.
3. Iterates through incoming updates and checks if the user is new. If so, the user is added to the database.

#### Main Loop:

1. Continuously listens for new updates from the Telegram bot.
2. For each update, checks if the user is new and adds them to the database if necessary.
3. Sends the update to the update handler goroutine for processing.

