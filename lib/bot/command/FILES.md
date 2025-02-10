# lib/bot/command/addNewUsertoMap.go  
# Package Name and Imports  
The package name is **command**. The imports are:  
* "log"  
* "github.com/JackBekket/hellper/lib/database"  
* "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
## External Data and Input Sources  
The external data and input sources are:  
* `tgbotapi.Message` object, which is used to get the chat ID and username of the user  
* `database` package, which is used to add a new user to the database  
* `msgTemplates` map, which is used to get the hello message template  
  
## TODO Comments  
The list of TODO comments is:  
* None explicitly marked as TODO, but there are some commented out code blocks that may need to be revisited:  
	+ `// check for registration`  
	+ `//	registred := IsAlreadyRegistred(session, chatID)`  
	+ `/*  
		if registred {  
			c.usersDb[chatID] = db.User{updateMessage.Chat.ID, updateMessage.Chat.UserName, 1}  
		}  
	*/`  
  
## Code Summary  
### Function `AddNewUserToMap`  
This function adds a new user to the database and assigns a "Dialog_status" of 3. It takes a `tgbotapi.Message` object as input and uses it to get the chat ID and username of the user. It then creates a new `database.User` object and adds it to the database. Finally, it sends a hello message to the user using the `tgbotapi` package.  
  
### Database Interaction  
The code interacts with the database by adding a new user to it. The `database.AddUser` function is used to add the user to the database.  
  
### Bot Interaction  
The code interacts with the bot by sending a hello message to the user. The `tgbotapi.NewMessage` function is used to create a new message, and the `c.bot.Send` function is used to send the message.  
  
  
  
# lib/bot/command/cases.go  
# Package/Component Name  
The package/component name is **command**.  
  
## Imports  
The following imports are used in this package:  
* `context`  
* `fmt`  
* `log`  
* `strings`  
* `db` from `github.com/JackBekket/hellper/lib/database`  
* `langchain` from `github.com/JackBekket/hellper/lib/langchain`  
* `localai` from `github.com/JackBekket/hellper/lib/localai`  
* `stt` from `github.com/JackBekket/hellper/lib/localai/audioRecognition`  
* `imgrec` from `github.com/JackBekket/hellper/lib/localai/imageRecognition`  
* `tgbotapi` from `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
* `godotenv` from `github.com/joho/godotenv`  
  
## External Data/Input Sources  
The external data/input sources used in this package are:  
* `db.UsersMap`: a map of user data  
* `updateMessage`: a message object from the Telegram Bot API  
* `ai_endpoint`: a string representing the AI endpoint URL  
* `language`: a string representing the language selected by the user  
* `ctx`: a context object with a value for the "user" key  
  
## TODO Comments  
The following TODO comments are found in the code:  
* `TODO: Write down user choice`  
  
## Summary of Major Code Parts  
### User Interaction  
The code handles user interactions through the `ChooseModel`, `HandleModelChoose`, and `DialogSequence` functions. These functions update the user's dialog status, attach a model to the user's profile, and handle the main dialog sequence.  
  
### Model Attachment  
The `attachModel` function attaches a model to a user's profile, and the `AttachKey` function attaches an API key to a user's profile.  
  
### Dialog Status Updates  
The `ChangeDialogStatus` function updates a user's dialog status, and the `WrongResponse` function sends a message to the user when they provide an incorrect response.  
  
### AI Connection  
The `ConnectingToAiWithLanguage` function connects the user to an AI node with a selected language, and the `DialogSequence` function generates an image or text response based on the user's input.  
  
### Database Functions  
The `GetUsersDb` function returns a map of user data, and the `GetUser` function returns a user object by ID.  
  
  
  
# lib/bot/command/commandHandler.go  
# Package/Component Name  
The package/component name is **command**.  
  
## Imports  
The following imports are used in this package:  
* `log`  
* `os`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/JackBekket/hellper/lib/langchain`  
* `tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"`  
  
## External Data/Input Sources  
The external data/input sources used in this package are:  
* Environment variables (e.g. `AI_ENDPOINT`)  
* User input from Telegram bot messages  
* Database (e.g. `database.UsersMap`)  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### HandleCommands Function  
The `HandleCommands` function handles incoming Telegram bot messages and executes corresponding commands.  
#### Supported Commands  
The following commands are supported:  
* `/image`: generates an image based on the provided prompt  
* `/restart`: restarts the user session  
* `/help`: sends a help message  
* `/search_doc`: searches for documents based on the provided prompt  
* `/instruct`: generates content using a local AI model  
* `/usage`: sends usage information  
* `/helper`: sends a media helper message  
* `/setContext`: sets the user context  
* `/clearContext`: clears the user context  
  
### Code Structure  
The code is structured around the `HandleCommands` function, which uses a switch statement to determine the command to execute based on the incoming message.  
  
  
  
# lib/bot/command/msgTemplates.go  
# Package: command  
## Imports  
The provided code does not contain any explicit imports. However, it is likely that this package is part of a larger Go program and may import other packages or modules.  
  
## External Data and Input Sources  
The code defines a map of message templates (`msgTemplates`) that contains various messages and prompts for a chatbot. These templates seem to be the primary external data source for this package.  
  
## TODO Comments  
There are no TODO comments in the provided code snippet.  
  
## Code Summary  
### Message Templates  
The `msgTemplates` map contains several key-value pairs, where each key corresponds to a specific message or prompt, and the value is the actual message text. These templates appear to be used for various chatbot interactions, such as greetings, help messages, and prompts for user input.  
  
### Package Overview  
The `command` package seems to be responsible for handling chatbot interactions, including message templates and potentially other command-related functionality. However, the provided code snippet only shows the definition of the `msgTemplates` map.  
  
  
  
# lib/bot/command/newCommander.go  
# Package: command  
## Imports  
The package imports the following modules:  
* `context`  
* `github.com/JackBekket/hellper/lib/database`  
* `tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"`  
  
## External Data and Input Sources  
The package uses the following external data and input sources:  
* `database.User` data from the `github.com/JackBekket/hellper/lib/database` package  
* `tgbotapi.BotAPI` instance from the `github.com/go-telegram-bot-api/telegram-bot-api/v5` package  
* `context.Context` instance  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Commander Struct  
The package defines a `Commander` struct with the following fields:  
* `bot`: a pointer to a `tgbotapi.BotAPI` instance  
* `usersDb`: a map of `int64` to `database.User` instances  
* `ctx`: a `context.Context` instance  
  
### NewCommander Function  
The package defines a `NewCommander` function that returns a new `Commander` instance.  
The function takes the following parameters:  
* `bot`: a pointer to a `tgbotapi.BotAPI` instance  
* `usersDb`: a map of `int64` to `database.User` instances  
* `ctx`: a `context.Context` instance  
  
  
  
# lib/bot/command/ui.go  
# Package Name and Imports  
The package name is **command**. The imports are:  
* `tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `chatID`: an integer representing the chat ID  
* `msgTemplates`: a map of message templates, where "case1" is a key  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Render Model Menu  
The `RenderModelMenuLAI` function renders a LLaMA-based model menu with an inline keyboard. It sends a message with a reply markup to the chat.  
  
### Render Language Menu  
The `RenderLanguage` function renders a language menu with an inline keyboard. It sends a message with a reply markup to the chat, allowing the user to choose a language.  
  
  
  
# lib/bot/command/utils.go  
# Package Name and Imports  
The package name is `command`. The imports are:  
* `fmt`  
* `io`  
* `log`  
* `math/rand`  
* `net/http`  
* `net/url`  
* `os`  
* `path`  
* `path/filepath`  
* `db` from `github.com/JackBekket/hellper/lib/database`  
* `github.com/JackBekket/hellper/lib/embeddings`  
* `github.com/JackBekket/hellper/lib/localai`  
* `tgbotapi` from `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
* `github.com/joho/godotenv`  
  
## External Data and Input Sources  
The external data and input sources are:  
* Environment variables:  
	+ `PG_LINK`  
	+ `AI_BASEURL`  
	+ `OPENAI_API_KEY`  
	+ `IMAGE_GENERATION_MODEL`  
	+ `IMAGE_GENERATION_SUFFIX`  
* Database: `db`  
* File system: `media` directory  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Help Command Message  
The `HelpCommandMessage` function sends a help message to the user.  
  
### Search Documents  
The `SearchDocuments` function searches for documents based on a prompt and returns the results to the user.  
  
### Get Usage  
The `GetUsage` function retrieves the usage for a user and sends it to the user.  
  
### Send Media Helper  
The `SendMediaHelper` function sends a random media file to the user.  
  
### Generate New Image LAI SD  
The `GenerateNewImageLAI_SD` function generates a new image using stable diffusion and sends it to the user.  
  
  
  
