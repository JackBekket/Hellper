## Package: command

### Imports:

- log
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External data, input sources:

- updateMessage: A tgbotapi.Message object containing information about the incoming update.

### TODOs:

- check for registration
- IsAlreadyRegistred(session, chatID)

### Summary:

#### AddNewUserToMap function:

This function is responsible for adding a new user to the database and assigning them a "Dialog_status" of 3. It takes an updateMessage as input, which contains information about the incoming update. The function first extracts the chatID and username from the updateMessage. Then, it creates a new User object with the extracted information and the Dialog_status set to 3. The User object is then added to the database using the AddUser function from the database package.

After adding the user to the database, the function logs the user's ID and username. It then creates a new message using the "hello" template from the msgTemplates map and sends it to the user's chatID using the bot object.

The function also includes a commented-out section that checks if the user is already registered. This section is marked as TODO and needs to be implemented.

lib/bot/command/cases.go
## Package: command

### Imports:

```
context
fmt
log
strings

db "github.com/JackBekket/hellper/lib/database"
langchain "github.com/JackBekket/hellper/lib/langchain"
localai "github.com/JackBekket/hellper/lib/localai"
stt "github.com/JackBekket/hellper/lib/localai/audioRecognition"
imgrec "github.com/JackBekket/hellper/lib/localai/imageRecognition"
tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
godotenv "github.com/joho/godotenv"
```

### External Data, Input Sources:

- Database: `db.UsersMap` - A map of users, each with an `AiSession` struct containing information about their current AI session.
- Telegram Bot API: Used for sending messages and receiving updates from the Telegram bot.
- Environment variables: `AI_ENDPOINT` - The endpoint for the local AI node.

### TODOs:

1. Write down user choice when attaching a model.
2. Implement key validation.

### Summary:

The code defines a `Commander` struct that handles the main logic for the Telegram bot. It includes methods for:

1. Choosing a model:
   - `ChooseModel`: Handles user input for selecting a model and updates the user's `DialogStatus` to 4.
2. Handling model choice:
   - `HandleModelChoose`: Processes the user's choice of model and updates the user's `DialogStatus` to 5.
3. Attaching a model:
   - `attachModel`: Attaches the selected model to the user's profile and sends a confirmation message.
4. Attaching an API key:
   - `AttachKey`: Stores the API key in the user's profile.
5. Changing dialog status:
   - `ChangeDialogStatus`: Updates the user's `DialogStatus`.
6. Handling wrong responses:
   - `WrongResponse`: Sends a message to the user if they don't use the provided keyboard.
7. Connecting to AI with language:
   - `ConnectingToAiWithLanguage`: Sets up the AI sequence with the selected language and updates the user's `DialogStatus` to 6.
8. Dialog sequence:
   - `DialogSequence`: Handles the main loop of the bot, processing user input and sending responses.
9. Getting users from the database:
   - `GetUsersDb`: Returns the map of users from the database.
10. Getting a user by ID:
    - `GetUser`: Returns the user with the specified ID from the database.

The code also includes helper functions for handling voice messages, image recognition, and deleting files.

lib/bot/command/commandHandler.go
## Package: command

### Imports:

* log
* os
* github.com/JackBekket/hellper/lib/database
* github.com/JackBekket/hellper/lib/langchain
* tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:

* os.Getenv("AI_ENDPOINT") - Environment variable containing the AI endpoint URL.
* database.UsersMap - A map containing user data.

### TODOs:

* Implement the logic for the "image" command, including image generation and sending.
* Implement the logic for the "restart" command, including deleting user data from the database.
* Implement the logic for the "search_doc" command, including searching for documents and returning results.
* Implement the logic for the "instruct" command, including generating content based on user input and the selected AI model.
* Implement the logic for the "usage" command, including displaying usage information.
* Implement the logic for the "helper" command, including sending media helper information.
* Implement the logic for the "setContext" command, including setting the user's context.
* Implement the logic for the "clearContext" command, including clearing the user's context.

### Summary:

The code defines a function called `HandleCommands` that handles incoming commands from a Telegram bot. The function takes a message and a Commander object as input. It then uses a switch statement to handle different commands, such as "image", "restart", "help", "search_doc", "rag", "instruct", "usage", "helper", "setContext", and "clearContext". Each command has its own logic, which may involve interacting with the database, generating images, searching for documents, or sending messages to the user. The code also includes some TODO comments that indicate areas where more functionality needs to be implemented.

lib/bot/command/msgTemplates.go
## Package/Component: command

### Imports:

- map[string]string

### External Data, Input Sources:

- msgTemplates: A map containing various message templates for different commands and scenarios.

### TODOs:

- Implement the functionality for each command mentioned in the help_command message.

### Summary:

The code defines a package or component named "command" that appears to be responsible for handling various commands and providing corresponding messages. It utilizes a map called "msgTemplates" to store different message templates for various commands and scenarios. The code also includes a TODO comment indicating that the functionality for each command mentioned in the help_command message needs to be implemented.

The help_command message lists several commands, including:

- /help: Prints the help message.
- /restart: Restarts the session, potentially switching between local-ai and openai chatGPT.
- /search_doc: Searches for documents.
- /rag: Processes Retrival-Augmented Generation.
- /instruct: Uses a system prompt template instead of langchain (higher priority).
- /image: Generates an image.

The code snippet also mentions that all the functions are experimental and the bot can halt and catch fire, suggesting that the functionality is still under development and may not be fully functional.



lib/bot/command/newCommander.go
## Package: command

### Imports:
- context
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External data, input sources:
- usersDb: map[int64]database.User

### TODOs:
- Implement GetCommander() function

### Summary:
#### Commander struct:
The Commander struct is responsible for managing the bot's interactions with users and the database. It contains the following fields:
- bot: A pointer to the Telegram bot API instance.
- usersDb: A map of user IDs to database.User objects, representing the users the bot interacts with.
- ctx: A context object for managing the bot's lifecycle and operations.

#### NewCommander function:
This function creates a new Commander instance, taking the bot API, user database, and context as input. It initializes the Commander struct with the provided values and returns the newly created instance.

#### GetCommander function (TODO):
This function is not yet implemented and is marked as a TODO. It is expected to return a Commander instance, but the implementation details are not provided in the code.

lib/bot/command/ui.go
## Package: command

### Imports:
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:
- msgTemplates (not shown in the code, but mentioned in the RenderModelMenuLAI function)

### TODOs:
- None found

### Summary:
#### RenderModelMenuLAI:
This function renders a menu of LLaMA-based models with an inline keyboard. It creates a new message with the text from the msgTemplates["case1"] and a reply markup with an inline keyboard. The keyboard has three rows, each with a button that corresponds to a different model. When a user clicks on a button, it sends the model name as data to the bot.

#### RenderLanguage:
This function renders a menu for choosing a language with an inline keyboard. It creates a new message with the text "Choose a language or send 'Hello' in your desired language." and a reply markup with an inline keyboard. The keyboard has one row with two buttons, one for English and one for Russian. When a user clicks on a button, it sends the language name as data to the bot.



lib/bot/command/utils.go
## Package: command

### Imports:

```
fmt
io
log
math/rand
net/http
net/url
os
path
path/filepath
github.com/JackBekket/hellper/lib/database
github.com/JackBekket/hellper/lib/embeddings
github.com/JackBekket/hellper/lib/localai
github.com/go-telegram-bot-api/telegram-bot-api/v5
github.com/joho/godotenv
```

### External Data, Input Sources:

- Environment variables: PG_LINK, AI_BASEURL, OPENAI_API_KEY, IMAGE_GENERATION_MODEL, IMAGE_GENERATION_SUFFIX
- Database: UsersMap (from lib/database)
- Local AI models: stablediffusion (from lib/localai)

### TODOs:

- Implement error handling for all functions.
- Add more robust logging.
- Implement a mechanism for updating the AI session usage.

### Summary:

The code defines a `Commander` struct with several methods for handling commands and interactions with the Telegram bot.

#### HelpCommandMessage:

This method handles the "/help" command by sending a predefined message to the user.

#### SearchDocuments:

This method searches for documents based on a given prompt and returns the top results. It uses an external vector store to perform semantic search and retrieves the page content and score for each result.

#### GetUsage:

This method retrieves and displays the usage statistics for the current user, including prompt tokens, completion tokens, and total tokens.

#### SendMediaHelper:

This method sends a random video from the "media" directory to the user. It opens the video file, creates a new video message, and sends it to the user.

#### sendImage:

This method sends an image to the user. It first retrieves the image from the given URL using the provided authentication header. Then, it saves the image to a temporary file and sends it to the user.

#### getImage:

This method fetches an image from the given URL and saves it to a temporary file. It handles authentication and error handling.

#### DeleteFile:

This method deletes a file from the file system.

#### transformURL:

This method extracts the filename from a given URL.

#### GenerateNewImageLAI_SD:

This method generates a new image using the Stable Diffusion model. It takes a prompt, URL, chat ID, and bot as input. It then calls the local AI model to generate the image and sends it to the user.

