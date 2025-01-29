# lib/bot/command/addAdminTomap.go  
**Package Name:** command  
  
**Imports:**  
  
* `log`  
* `github.com/JackBekket/hellper/lib/database` (db)  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
  
**External Data/Inputs:**  
  
* `adminKey` (string)  
* `updateMessage` (*tgbotapi.Message)  
* `msgTemplates` (map[string]string) (not shown in the provided code snippet)  
  
**TODO Comments:** None  
  
**Summary:**  
  
### Functionality  
  
The provided code defines a function `AddAdminToMap` within the `Commander` struct. This function is responsible for adding an admin to a map and sending a message to the admin.  
  
### Flow  
  
The function takes two inputs: `adminKey` and `updateMessage`. It extracts the `chatID` from the `updateMessage` and uses it to create a new `db.User` object, which is then added to the `db.UsersMap`. The function also sets the `DialogStatus` to 2 and marks the user as an admin.  
  
The function then logs a message to the console indicating that the user has been authorized. It then sends a message to the admin with a confirmation message and a one-time reply keyboard with a single button labeled "GPT-3.5".  
  
### Notes  
  
The code uses the `db` package to interact with a database, and the `tgbotapi` package to send messages to a Telegram bot.  
  
**  
  
# lib/bot/command/addNewUsertoMap.go  
**Package/Component Name:** command  
  
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
  
Next, it creates a new `tgbotapi.Message` object and sends it to the user using the `c.bot.Send` method. The message includes a one-time reply keyboard with a single button labeled "Start!".  
  
The function also checks if the user is already registered, but this functionality is currently commented out.  
  
  
  
# lib/bot/command/cases.go  
**Package/Component Name:** `command`  
  
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
* `joho/godotenv`  
  
**External Data/Inputs:**  
  
* None  
  
**TODO Comments:**  
  
* `TODO: Write down user choice` in the `attachModel` function  
  
**Summary:**  
  
### Overview  
  
The `command` package provides a set of functions for handling Telegram bot commands. It appears to be part of a larger AI-powered chatbot system.  
  
### Functions  
  
#### InputYourAPIKey  
  
This function is responsible for handling the input of an OpenAI API key. It updates the user's dialog status to 3 and stores the key in memory.  
  
#### ChooseModel  
  
This function handles the selection of a model by the user. It updates the user's dialog status to 4 and stores the selected model in memory.  
  
#### HandleModelChoose  
  
This function is called when the user chooses a model. It updates the user's dialog status to 5 and starts a dialog sequence.  
  
#### DialogSequence  
  
This function is the main loop of the chatbot. It handles different types of updates (text, voice, or image) and performs corresponding actions.  
  
#### AttachKey  
  
This function is used to attach an API key to a user's profile.  
  
#### ChangeDialogStatus  
  
This function changes the dialog status of a user.  
  
#### GetUsersDb and GetUser  
  
These functions are used to retrieve and manipulate user data.  
  
### Notes  
  
* The code uses a database to store user data and dialog status.  
* It interacts with various AI-powered services, including OpenAI and local AI models.  
* There are TODO comments indicating areas that require further development.  
  
**  
  
# lib/bot/command/checkAdmin.go  
**Package Name:** command  
  
**Imports:**  
  
* `fmt`  
* `github.com/JackBekket/hellper/lib/bot/env`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
**External Data/Inputs:**  
  
* `adminData` (map of `env.AdminData` objects)  
* `updateMessage` (a `tgbotapi.Message` object)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Overview  
  
The `CheckAdmin` function is part of the `command` package and is responsible for updating the "dialogStatus" in the database based on the provided `adminData` and `updateMessage` inputs.  
  
### Functionality  
  
The function iterates over the `adminData` map and checks if the `ID` of the current admin matches the `Chat.ID` of the `updateMessage`. If a match is found, it checks if the `GPTKey` is not empty. If it is not empty, it adds the admin to a map using the `AddAdminToMap` method and returns. If the `GPTKey` is empty, it sends a message to the chat indicating that the environment variable is missing, and then adds the user to a map using the `AddNewUserToMap` method.  
  
### Logic  
  
The function appears to be designed to handle two cases: adding an admin to a map if the `GPTKey` is present, and adding a new user to a map if the `GPTKey` is missing. The logic is straightforward and easy to follow.  
  
**  
  
# lib/bot/command/msgTemplates.go  
**Package Name:** command  
  
**Imports:**  
  
* None  
  
**External Data/Inputs:**  
  
* None  
  
**TODO Comments:**  
  
* None  
  
**Summary:**  
  
### Message Templates  
  
The `msgTemplates` map is defined, which contains a set of pre-defined message templates. These templates are used to generate responses to user input. The templates are categorized into different cases, such as "hello", "case0", "await", "case1", and "help_command". Each template provides a specific message or instruction to the user.  
  
### Template Descriptions  
  
The templates are used to provide information to the user about the bot's capabilities and limitations. For example, the "help_command" template provides a list of available commands and their functions. The "case0" and "case1" templates are used to guide the user through the process of selecting a model to use.  
  
### Conclusion  
  
The `msgTemplates` map is a critical component of the command package, as it provides a set of pre-defined messages that can be used to interact with the user. The templates are designed to be flexible and adaptable, allowing the bot to respond to a wide range of user inputs and scenarios.  
  
  
  
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
* `github.com/JackBekket/hellper/lib/agent`  
* `github.com/JackBekket/hellper/lib/database` (db)  
* `github.com/JackBekket/hellper/lib/embeddings`  
* `github.com/JackBekket/hellper/lib/localai`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
* `joho/godotenv`  
  
**External Data/Inputs:**  
  
* None  
  
**TODO Comments:**  
  
* OBSOLETE in `RAG` function  
* Refactor to better readability in `RAG` function  
* Superagents in `RAG` function  
* None in other functions  
  
**Summary:**  
  
### HelpCommandMessage  
The `HelpCommandMessage` function sends a help message to a user based on the `msgTemplates` map.  
  
### SearchDocuments  
This function performs a semantic search for a given prompt and returns the results. It also sends the results to the user.  
  
### RAG  
This function is marked as OBSOLETE and is not used. It is intended to perform a retrieval-augmented generation task, but it is not implemented.  
  
### GetUsage  
This function sends the usage information to the user.  
  
### GenerateNewImageLAI_SD  
This function generates a new image using the stable diffusion model and sends it to the user.  
  
### SendMediaHelper  
This function sends a random media file from the media directory to the user.  
  
### DeleteFile  
This function deletes a file.  
  
### getImage  
This function downloads an image from a URL and saves it to a temporary directory.  
  
**  
  
