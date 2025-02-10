# lib/bot/dialog/dialog.go  
**Package/Component Name:** dialog  
  
**Imports:**  
  
* `log`  
* `os`  
* `regexp`  
* `strings`  
* `github.com/JackBekket/hellper/lib/bot/command`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
**External Data/Inputs:**  
  
* `updates` channel from unknown source  
* `bot` object from unknown source  
* `comm` object from unknown source  
* `AI_ENDPOINT` environment variable  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Main Function  
  
The `HandleUpdates` function is the main entry point for handling updates from a Telegram bot. It iterates over a channel of updates and handles each update accordingly.  
  
### Handling Updates  
  
The function checks if the update is a message or a callback query. If it's a message, it checks if the message is a group message and if it's not, it skips the message. If it's a group message, it processes the message based on the user's dialog status.  
  
If the user's dialog status is 3, it calls the `ChooseModel` function from the `comm` object. If the status is 4 or 5, it calls the `WrongResponse` or `DialogSequence` functions, respectively. If the status is 6, it calls the `DialogSequence` function with the `ai_endpoint` environment variable.  
  
If the update is a callback query, it handles the query based on the user's dialog status. If the status is 4, it calls the `HandleModelChoose` function. If the status is 5, it calls the `ConnectingToAiWithLanguage` function with the `ai_endpoint` environment variable.  
  
**  
  
