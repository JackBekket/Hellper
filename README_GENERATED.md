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

- Environment variables: OPENAI_API_KEY, PG_LINK, TG_KEY, ADMIN_ID, AI_ENDPOINT

### Code Summary:

1. Loads environment variables using godotenv.Load().
2. Retrieves API token from OPENAI_API_KEY environment variable.
3. Initializes a Telegram bot using the TG_KEY environment variable.
4. Retrieves admin ID and GPT key from environment variables.
5. Initializes a map of admin data.
6. Initializes a database for users using database.UsersMap.
7. Creates a context for the bot.
8. Creates a new commander using the bot, database, and context.
9. Sets up a channel for handling updates.
10. Starts a goroutine to handle updates using dialog.HandleUpdates.
11. Continuously listens for updates from the bot and sends them to the update channel.
12. Checks if a user is new based on their ID in the database.
13. If a user is new, their entry is created in the database.
14. Handles inline keyboards by checking for callback queries.
15. Retrieves the chat ID from the update.
16. Checks if the user is new based on their ID in the database.
17. If the user is new, their entry is created in the database.

#### Project Package Structure:

- hellper/lib/bot/command
- hellper/lib/bot/dialog
- hellper/lib/bot/env
- hellper/lib/database

