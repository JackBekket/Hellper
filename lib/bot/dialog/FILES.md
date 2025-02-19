# lib/bot/dialog/dialog.go  
# Package Name and Imports  
The package name is `dialog`. The imports are:  
* `log`  
* `os`  
* `regexp`  
* `strings`  
* `github.com/JackBekket/hellper/lib/bot/command`  
* `github.com/JackBekket/hellper/lib/database`  
* `tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `updates` channel of type `tgbotapi.Update`  
* `bot` of type `*tgbotapi.BotAPI`  
* `comm` of type `command.Commander`  
* `db_service` of type `*database.Service`  
* Environment variables:  
	+ `AI_ENDPOINT`  
  
## TODO Comments  
The TODO comments are:  
1. `// TODO: should not be here?` (referring to the `ai_endpoint` variable)  
2. `//TODO: when we do the endpoints part, remove this hardcode` (referring to the `isRegistered` variable)  
  
## Code Summary  
### HandleUpdates Function  
The `HandleUpdates` function handles updates from the Telegram bot API. It iterates over the `updates` channel and processes each update. If the update is a message, it checks if the message is from a group or a private chat. If it's a group chat, it checks if the message is addressed to the bot. If not, it ignores the message.  
  
### User Dialog Status  
The function checks the user's dialog status and performs different actions based on the status. The possible statuses are:  
* 3: Choose model  
* 4: Wrong response  
* 5: Connecting to AI with language  
* 6: Dialog sequence  
  
### Callback Query Logic  
The function also handles callback queries for inline buttons. It checks the user's dialog status and performs different actions based on the status.  
  
### Database Interaction  
The function interacts with the database to retrieve and update user data. It uses the `db_service` to check if a user exists in the database and to retrieve the user's session data.  
  
  
  
