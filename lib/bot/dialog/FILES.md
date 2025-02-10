# lib/bot/dialog/dialog.go  
# Package Name and Imports  
The package name is **dialog**. The imports are:  
* "log"  
* "os"  
* "regexp"  
* "strings"  
* "github.com/JackBekket/hellper/lib/bot/command"  
* "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
## External Data and Input Sources  
The external data and input sources are:  
* **tgbotapi.Update**: updates from the Telegram Bot API  
* **tgbotapi.BotAPI**: the Telegram Bot API object  
* **command.Commander**: the commander object for handling commands  
* **os.Getenv("AI_ENDPOINT")**: the AI endpoint URL from the environment variable  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### HandleUpdates Function  
The **HandleUpdates** function handles updates from the Telegram Bot API. It checks if the update is a message or a callback query. If it's a message, it checks if the message is from a group or a private chat. If it's a group chat, it checks if the message is addressed to the bot. If it's not addressed to the bot, it ignores the message.  
  
### Message Handling  
If the message is addressed to the bot, it removes the bot's username from the message text and checks the user's dialog status. Based on the dialog status, it performs different actions:  
* If the dialog status is 3, it chooses a model.  
* If the dialog status is 4 or 5, it handles a wrong response.  
* If the dialog status is 6, it starts a dialog sequence with the AI endpoint.  
  
### Callback Query Handling  
If the update is a callback query, it checks the user's dialog status and performs different actions:  
* If the dialog status is 4, it handles the model choose callback.  
* If the dialog status is 5, it connects to the AI with the chosen language.  
  
  
  
