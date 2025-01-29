# main.go  
**Package/Component Name:** main  
  
**Imports:**  
  
* `context`  
* `log`  
* `os`  
* `strconv`  
* `github.com/JackBekket/hellper/lib/bot/command`  
* `github.com/JackBekket/hellper/lib/bot/dialog`  
* `github.com/JackBekket/hellper/lib/bot/env`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
* `github.com/joho/godotenv`  
  
**External Data/Inputs:**  
  
* `OPENAI_API_KEY` (environment variable)  
* `PG_LINK` (environment variable)  
* `TG_KEY` (environment variable)  
* `ADMIN_ID` (environment variable)  
* `AI_ENDPOINT` (environment variable)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Initialization  
  
The `main` function initializes the bot API with the provided `TG_KEY` and sets up the database and commander.  
  
### Database and Commander Setup  
  
The code sets up a map of users (`usersDatabase`) and initializes the commander (`comm`) with the bot API and the database.  
  
### Update Handling  
  
The code sets up an update handler to process incoming updates from the Telegram API. It checks for new users and creates an entry in the database if necessary.  
  
### Logic  
  
The code checks for callback queries and updates the database accordingly. It also handles inline keyboards and checks for new users.  
  
**  
  
