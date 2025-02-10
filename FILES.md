# main.go  
**Package Name:** main  
  
**Imports:**  
  
* `context`  
* `log`  
* `os`  
* `github.com/JackBekket/hellper/lib/bot/command`  
* `github.com/JackBekket/hellper/lib/bot/dialog`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
* `github.com/joho/godotenv`  
  
**External Data/Inputs:**  
  
* `TG_KEY` environment variable  
* `os` environment variables (not explicitly used, but potentially used by `godotenv.Load()`)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Initialization  
  
The `main` function initializes the bot by loading the environment variables and creating a new bot API instance. It also sets up the database and commander.  
  
### Bot Setup  
  
The bot is authorized and a channel is created to handle updates. The `dialog` package is used to handle incoming updates.  
  
### Main Loop  
  
The main function enters an infinite loop, waiting for updates and processing them through the `dialog` package.  
  
**  
  
