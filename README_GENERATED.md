# Package: main

### Imports:

- context
- log
- os
- github.com/JackBekket/hellper/lib/bot/command
- github.com/JackBekket/hellper/lib/bot/dialog
- github.com/JackBekket/hellper/lib/database
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
- godotenv "github.com/joho/godotenv"

### External Data, Input Sources:

- Environment variable: TG_KEY (telegram bot token)

### TODOs:

- None found

### Summary:

#### Initialization:

The code starts by loading environment variables using `godotenv.Load()`. It then retrieves the telegram bot token from the TG_KEY environment variable. A new Telegram bot instance is created using the token, and an error is logged if the token is missing.

#### Database and Commander Initialization:

A database instance is initialized using `database.UsersMap`. A context is created using `context.Background()`. A new commander instance is created using the bot, database, and context.

#### Bot Authorization and Logging:

The bot is authorized, and a log message is printed indicating the authorized account.

#### Update Handling:

An update channel is created to handle incoming updates. A new update is created with a timeout of 60 seconds. The bot's GetUpdatesChan method is used to retrieve updates from the Telegram API.

#### Dialog Handling:

A goroutine is started to handle updates using the `dialog.HandleUpdates` function. This function takes the update channel, bot instance, and commander instance as arguments.

#### Main Loop:

The main loop iterates over the updates received from the Telegram API. Each update is sent to the update channel, which is then handled by the goroutine responsible for dialog handling.

