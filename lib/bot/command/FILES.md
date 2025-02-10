# lib/bot/command/addNewUsertoMap.go  
**Package Name:** command  
  
**Imports:**  
  
* `log`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
**External Data/Inputs:**  
  
* `updateMessage` (of type `*tgbotapi.Message`)  
* `session` (not used in this code snippet)  
  
**TODO Comments:**  
  
* None found in this code snippet  
  
**Summary:**  
  
### AddNewUserToMap Function  
  
The `AddNewUserToMap` function is part of the `Commander` struct and is responsible for adding a new user to the database and assigning a "Dialog_status" of 0. It takes an `updateMessage` of type `*tgbotapi.Message` as input.  
  
The function first creates a new `database.User` struct and adds it to the database using the `database.AddUser` method. It then logs a message to the console indicating the addition of a new user.  
  
Next, it sends a message to the user using the `c.bot.Send` method, with a message template from `msgTemplates["hello"]`.  
  
The function also checks if the user is already registered, but this part is commented out.  
  
  
  
# lib/bot/command/cases.go  
**Package/Component Name:** command  
  
**Imports:**  
* `context`  
* `fmt`  
* `log`  
* `strings`  
* `github.com/JackBekket/hellper/lib/database` (db)  
* `github.com/JackBekket/hellper/lib/langchain`  
* `github.com/JackBekket/hellper/lib/localai`  
* `github.com/JackBekket/hellper/lib/localai/audioRecognition` (stt)  
* `github.com/JackBekket/hellper/lib/localai/imageRecognition` (imgrec)  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
* `github.com/joho/godotenv` (godotenv)  
  
**External Data/Inputs:**  
  
* `updateMessage` (of type `*tgbotapi.Message` or `*tgbotapi.CallbackQuery`)  
* `chatID` (int64)  
* `ai_endpoint` (string)  
* `id` (int64)  
  
**TODO Comments:**  
  
* `TODO: Write down user choice` in the `attachModel` function  
  
**Summary:**  
  
### Commander Functions  
  
The `Commander` type provides several functions to interact with a Telegram bot. These functions are responsible for handling user input and updating the dialog status.  
  
### Dialog Flow  
  
The code implements a dialog flow, where the bot interacts with the user to choose a model and then starts a dialog sequence. The dialog flow is controlled by the `DialogStatus` field, which is updated based on user input.  
  
### Model Selection  
  
The `ChooseModel` function is responsible for handling the initial model selection. It updates the `DialogStatus` to 4.  
  
### Model Attachment  
  
The `attachModel` function is used to attach a model to a user's profile. It updates the `DialogStatus` to 5.  
  
### Dialog Sequence  
  
The `DialogSequence` function is the main loop of the dialog flow. It handles user input and updates the `DialogStatus` accordingly. It also interacts with the `langchain` and `localai` libraries to perform tasks such as image and voice recognition.  
  
### User Management  
  
The `GetUsersDb` and `GetUser` functions are used to interact with the database and retrieve user information.  
  
### Error Handling  
  
The code includes basic error handling using the `log` package.  
  
**  
  
# lib/bot/command/commandHandler.go  
**Package Name:** command  
  
**Imports:**  
  
* `log`  
* `os`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/JackBekket/hellper/lib/langchain`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
**External Data/Inputs:**  
  
* `os` environment variable `AI_ENDPOINT`  
* `database.UsersMap` (a map of users)  
* `user.AiSession.GptModel` and `user.AiSession.GptKey` (AI session parameters)  
* `user.Network` (network information)  
* `chatID` and `message` (Telegram bot API message and chat ID)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Main Functionality  
  
The `HandleCommands` function is responsible for handling various commands sent to the Telegram bot. It takes in a `message` and a `comm` object as inputs.  
  
### Command Handling  
  
The function switches on the `message.Command()` value to determine which command to execute. The following commands are implemented:  
  
* `image`: Generates an image link based on the provided prompt. If no prompt is provided, it generates a new image using the `GenerateNewImageLAI_SD` function.  
* `restart`: Deletes the user's session data from the `database.UsersMap`.  
* `help`: Sends a help message to the user.  
* `search_doc`: Searches for documents based on the provided prompt.  
* `instruct`: Calls the `GenerateContentInstruction` function from `langchain` to generate content based on the provided prompt.  
* `usage`: Retrieves usage information.  
* `helper`: Sends a media helper message to the user.  
* `setContext`: Sets a context for the user based on the provided name.  
* `clearContext`: Clears the user's context.  
* `default`: Handles unknown commands.  
  
### Notes  
  
The code uses various external dependencies, including the `database` and `langchain` packages, as well as the Telegram bot API. It also interacts with the `os` environment variable to retrieve the `AI_ENDPOINT` value.  
  
**  
  
# lib/bot/command/msgTemplates.go  
**Package Name:** command  
  
**Imports:**  
  
* None (no imports found)  
  
**External Data and Input Sources:**  
  
* None (no external data or input sources found)  
  
**TODO Comments:**  
  
* None (no TODO comments found)  
  
**Summary:**  
  
### Message Templates  
  
The code defines a map of message templates, `msgTemplates`, which contains key-value pairs. Each key corresponds to a specific message, and the value is the actual message content. The templates include:  
  
* "hello": a welcome message  
* "await": a waiting message  
* "case1": a prompt for choosing a model  
* "help_command": a help message with instructions for additional commands  
  
This code provides a set of pre-defined messages that can be used by the bot to interact with users.  
  
**  
  
# lib/bot/command/newCommander.go  
**Package Name:** command  
  
**Imports:**  
  
* `context`  
* `github.com/JackBekket/hellper/lib/database`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
**External Data/Inputs:**  
  
* `bot *tgbotapi.BotAPI` (Telegram bot API)  
* `usersDb map[int64]database.User` (database of users)  
* `ctx context.Context` (context for the bot)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Commander Structure  
  
The `Commander` struct is defined, which holds references to a Telegram bot API, a map of users, and a context. This struct is used to manage the bot's interactions with users.  
  
### NewCommander Function  
  
The `NewCommander` function is implemented, which initializes a new `Commander` instance. It takes three inputs: a `bot` instance, a `usersDb` map, and a `ctx` context. The function returns a pointer to a newly created `Commander` instance.  
  
### Missing Function  
  
The `GetCommander` function is not implemented in this file.  
  
**  
  
# lib/bot/command/ui.go  
**Package Name:** command  
  
**Imports:**  
  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
  
**External Data/Inputs:**  
  
* `chatID` (int64)  
* `msgTemplates` (map[string]string) (not shown in the provided code snippet)  
  
**TODO Comments:**  
  
* None found in the provided code snippet  
  
**Summary:**  
  
### Overview  
  
The provided code defines two functions within the `command` package: `RenderModelMenuLAI` and `RenderLanguage`. These functions are part of a Telegram bot API and are used to render menus with inline keyboards.  
  
### RenderModelMenuLAI  
  
This function renders a menu with inline keyboard buttons for selecting LLaMA-based models. It sends a message with a predefined template and inline keyboard markup to the specified `chatID`.  
  
### RenderLanguage  
  
This function renders a language selection menu with inline keyboard buttons. It sends a message with a predefined template and inline keyboard markup to the specified `chatID`. The menu allows users to select English or Russian as their preferred language.  
  
**  
  
# lib/bot/command/utils.go  
**Package/Component Name:** command  
  
**Imports:**  
  
* `fmt`  
* `io`  
* `log`  
* `math/rand`  
* `net/http`  
* `net/url`  
* `os`  
* `path/filepath`  
* `github.com/JackBekket/hellper/lib/database` (db)  
* `github.com/JackBekket/hellper/lib/embeddings` (embeddings)  
* `github.com/JackBekket/hellper/lib/localai` (localai)  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
* `joho/godotenv` (godotenv)  
  
**External Data/Inputs:**  
  
* Environment variables:  
	+ `PG_LINK`  
	+ `AI_BASEURL`  
	+ `OPENAI_API_KEY`  
	+ `IMAGE_GENERATION_MODEL`  
	+ `IMAGE_GENERATION_SUFFIX`  
* Database data:  
	+ `db.UsersMap` (map of user IDs to user data)  
* Telegram API:  
	+ `chatID` (chat ID)  
	+ `updateMessage` (update message)  
	+ `bot` (Telegram bot API)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Help Command  
  
The `HelpCommandMessage` function sends a help message to a user based on their ID.  
  
### Search Documents  
  
The `SearchDocuments` function performs a semantic search on a prompt and returns a list of results. It also sends a message to the user with the results.  
  
### Get Usage  
  
The `GetUsage` function sends a message to the user with their usage statistics.  
  
### Generate New Image (Stable Diffusion)  
  
The `GenerateNewImageLAI_SD` function generates a new image using the stable diffusion model and sends it to the user.  
  
### Send Media Helper  
  
The `SendMediaHelper` function sends a random media file to the user.  
  
### Image Processing  
  
The `getImage` function downloads an image from a URL and saves it to a temporary file. The `DeleteFile` function deletes a file. The `transformURL` function transforms a URL into a file name.  
  
**  
  
