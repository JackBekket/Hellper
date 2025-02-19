# lib/bot/command/addNewUsertoMap.go  
# Package Name and Imports  
The package name is **command**. The imports are:  
* "log"  
* "github.com/JackBekket/hellper/lib/database"  
* "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
## External Data and Input Sources  
The external data and input sources are:  
* `updateMessage` of type `*tgbotapi.Message`  
* `base_url` of type `string`  
* `database` which is used to store and retrieve user data  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Function `AddNewUserToMap`  
This function adds a new user to the database and assigns a "Dialog_status" of 3. It takes two parameters: `updateMessage` and `base_url`. The function:  
* Extracts the chat ID from the `updateMessage`  
* Creates a new `database.User` object with the extracted chat ID, username, and other default values  
* Sets the `Base_url` of the user's AI session to the provided `base_url`  
* Adds the new user to the database using `database.AddUser`  
* Logs a message indicating that a new user has been added to the database  
* Sends a "hello" message to the user using the `c.bot.Send` method  
  
  
  
# lib/bot/command/cases.go  
# Package/Component Name  
The package/component name is **command**.  
  
## Imports  
The following imports are used in the code:  
* `context`  
* `fmt`  
* `log`  
* `strings`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/JackBekket/hellper/lib/langchain`  
* `github.com/JackBekket/hellper/lib/localai`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
## External Data/Input Sources  
The external data/input sources used in the code are:  
* `tgbotapi.Message`  
* `tgbotapi.CallbackQuery`  
* `database.Service`  
* `localai.audioRecognition`  
* `localai.imageRecognition`  
  
## TODO Comments  
The following TODO comments are found in the code:  
* `TODO: Write down user choice`  
* `TODO: same for hardcode`  
  
## Code Summary  
### User Interaction  
The code handles user interactions through the `ChooseModel`, `HandleModelChoose`, and `DialogSequence` functions. These functions update the user's dialog status and render models or language options.  
  
### Model Selection  
The `attachModel` function attaches a model to a user's profile, and the `RenderModels` function renders models for a user.  
  
### Language Selection  
The `ConnectingToAiWithLanguage` function connects a user to an AI node with a selected language.  
  
### Dialog Loop  
The `DialogSequence` function generates an image or text response based on user input and updates the dialog status.  
  
### Database Functions  
The `GetUsersDb`, `GetUser`, and `RenderModelsForRegisteredUser` functions interact with the in-memory database.  
  
### Error Handling  
The `WrongResponse` function handles incorrect user responses.  
  
### Recovery  
The `RecoverUserAfterDrop` function recovers a user's session after a drop.  
  
  
  
# lib/bot/command/commandHandler.go  
# Package Name and Imports  
The package name is **command**. The imports are:  
* "log"  
* "os"  
* "github.com/JackBekket/hellper/lib/database"  
* "github.com/JackBekket/hellper/lib/langchain"  
* "tgbotapi": "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
## External Data and Input Sources  
The external data and input sources are:  
* Environment variables: "AI_ENDPOINT"  
* Database: "database.Service" and "database.UsersMap"  
* Telegram bot API: "tgbotapi"  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### HandleCommands Function  
The `HandleCommands` function handles incoming messages from the Telegram bot. It uses a switch statement to determine which command to execute based on the message's command.  
  
### Commands  
The available commands are:  
* "image": generates an image link  
* "reload": reloads the session  
* "clear": clears the dialog thread from the database  
* "purge": deletes all user data from the database and restarts the session  
* "drop": drops the session  
* "help": sends a help message  
* "search_doc": searches for documents  
* "instruct": generates content using the local AI  
* "usage": gets the usage  
* "helper": sends a media helper  
* "setContext": sets the context  
* "clearContext": clears the context  
  
### Functionality  
The function uses the `Commander` and `database.Service` objects to interact with the database and the Telegram bot API. It also uses the `langchain` package to generate content.  
  
  
  
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
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Commander Struct  
The package defines a `Commander` struct with the following fields:  
* `bot`: a pointer to a `tgbotapi.BotAPI` instance  
* `usersDb`: a map of `int64` to `database.User` instances  
* `ctx`: a `context.Context` instance  
  
### NewCommander Function  
The package defines a `NewCommander` function that returns a new `Commander` instance. The function takes the following parameters:  
* `bot`: a pointer to a `tgbotapi.BotAPI` instance  
* `usersDb`: a map of `int64` to `database.User` instances  
* `ctx`: a `context.Context` instance  
  
  
  
# lib/bot/command/ui.go  
# Package/Component Name  
The package/component name is **command**.  
  
## Imports  
The following imports are used in this package:  
* `fmt`  
* `github.com/JackBekket/hellper/lib/database`  
* `tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"`  
* `github.com/go-telegram/bot/models`  
  
## External Data/Input Sources  
The external data/input sources used in this package are:  
* `database` service, which provides access to the database  
* `user` data, which contains information about the user, including their AI session details  
* `chatID`, which is the ID of the chat where the messages are being sent  
* `models_list`, which is a list of models retrieved from the database  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Render Models  
The `RenderModels` function retrieves a list of models from the database and renders them as a message with an inline keyboard. If an error occurs, it sends an error message to the chat.  
  
### Get Models  
The `getModels` function retrieves a list of models from the database using the provided `db_service` and `user` data.  
  
### Render Language Menu  
The `RenderLanguage` function renders a language menu with an inline keyboard, allowing the user to choose a language.  
  
### Create Models Markup  
The `CreateModelsMarkup` function creates an inline keyboard markup for the provided list of models.  
  
  
  
# lib/bot/command/utils.go  
# Package/Component Name  
The package/component name is **command**.  
  
## Imports  
The following imports are used in the code:  
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
  
## External Data/Input Sources  
The following external data/input sources are used in the code:  
* Environment variables:  
	+ `PG_LINK`  
	+ `AI_BASEURL`  
	+ `OPENAI_API_KEY`  
	+ `IMAGE_GENERATION_MODEL`  
	+ `IMAGE_GENERATION_SUFFIX`  
* Database: `db`  
* File system: `media` directory, `tmp` directory  
* Network: `http` requests to external APIs  
  
## TODO Comments  
The following TODO comments are found in the code:  
* None  
  
## Code Summary  
### Functions  
The following functions are defined in the code:  
#### HelpCommandMessage  
Sends a help message to the user.  
#### SearchDocuments  
Searches for documents based on a prompt and returns the results to the user.  
#### GetUsage  
Gets the usage for a user and sends it to the user.  
#### SendMediaHelper  
Sends a random media file to the user.  
#### sendImage  
Sends an image to the user.  
#### getImage  
Gets an image from a URL and saves it to a file.  
#### DeleteFile  
Deletes a file.  
#### transformURL  
Transforms a URL to a file name.  
#### GenerateNewImageLAI_SD  
Generates a new image using stable diffusion and sends it to the user.  
#### SendMessage  
Sends a message to the user.  
  
### Variables  
The following variables are defined in the code:  
* `chatID`: the ID of the chat  
* `user`: the user object  
* `promt`: the prompt for the search  
* `maxResults`: the maximum number of results to return  
* `api_token`: the API token for the user  
* `store`: the vector store for the search  
* `results`: the results of the search  
* `score`: the score of the result  
* `score_string`: the score as a string  
* `msg`: the message to be sent to the user  
* `bot`: the bot object  
  
  
  
