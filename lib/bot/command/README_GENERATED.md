## Package: command

### Imports:
- log
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External data, input sources:
- updateMessage: A tgbotapi.Message object containing information about the incoming update.

### Summary:
#### AddNewUserToMap function:
This function is responsible for adding a new user to the database and assigning them a "Dialog_status" of 0. It takes an updateMessage as input, which contains information about the incoming update. The function first extracts the chatID and username from the updateMessage. Then, it creates a new User object with the extracted information and the Dialog_status set to 0. The User object is then added to the database using the AddUser function from the database package.

After adding the user to the database, the function logs the user's ID and username. It then creates a new message using the "hello" template from the msgTemplates map and sends it to the user with a one-time reply keyboard containing a "Start!" button.

The function also includes a commented-out section that checks if the user is already registered and updates the user's Dialog_status accordingly. However, this section is not currently being used.

lib/bot/command/msgTemplates.go
Package: command

Imports:

External data, input sources:

- msgTemplates: A map containing various message templates used in the package.

Summary:

The provided code defines a map called `msgTemplates` that stores various message templates used within the package. These templates are used to generate responses and messages for different scenarios, such as greetings, instructions, and error messages. The map contains key-value pairs, where the key is a string representing the template name, and the value is the corresponding message template.

The templates include:

- "hello": A greeting message indicating that the bot is working with a local AI node.
- "case0": Instructions for users to provide their OpenAI API key or password for local AI authentication.
- "await": A message indicating that the bot is awaiting a response or input.
- "case1": A prompt for users to choose a model to use.
- "ch_network": A prompt for users to choose a network to work with, either OpenAI or local AI.
- "help_command": A list of available commands for users, including /help, /restart, /search_doc, /rag, /instruct, and /image.

The code snippet provides a concise and organized way to manage and access various message templates within the package, ensuring consistency and clarity in the bot's communication with users.



lib/bot/command/ui.go
Package: command

Imports:
- github.com/go-telegram-bot-api/telegram-bot-api/v5

External data, input sources:
- msgTemplates: A map of message templates used in the code.

Summary:
### Render OpenAI Model Menu with Inline Keyboard
This function renders a menu of OpenAI models with an inline keyboard. It creates a new message with the specified chatID and the message template "case1". The message also includes an inline keyboard with two buttons: "gpt-3.5" and "gpt-4". The message is then sent to the chat using the bot.

### Render LLaMA-based Model Menu with Inline Keyboard
This function renders a menu of LLaMA-based models with an inline keyboard. It creates a new message with the specified chatID and the message template "case1". The message also includes an inline keyboard with three rows of buttons: "wizard-uncensored-13b", "wizard-uncensored-30b", and "tiger-gemma-9b-v1-i1". The message is then sent to the chat using the bot.

### Render Various AI Models Menu with Inline Keyboard
This function renders a menu of various AI models with an inline keyboard. It creates a new message with the specified chatID and the message template "case1". The message also includes an inline keyboard with four rows of buttons: "deepseek-coder-6b-instruct", "wizard-uncensored-code-34b", "tiger-gemma-9b-v1-i1", and "big-tiger-gemma-27b-v1". The message is then sent to the chat using the bot.

### Render Language Menu with Inline Keyboard
This function renders a menu of languages with an inline keyboard. It creates a new message with the specified chatID and a message asking the user to choose a language or send "Hello" in their desired language. The message also includes an inline keyboard with two buttons: "English" and "Russian". The message is then sent to the chat using the bot.



lib/bot/command/utils.go
## Package: command

### Imports:

```
fmt
log
math/rand
os
path/filepath
github.com/JackBekket/hellper/lib/database
github.com/JackBekket/hellper/lib/embeddings
github.com/go-telegram-bot-api/telegram-bot-api/v5
github.com/joho/godotenv
```

### External Data, Input Sources:

- Environment variables: PG_LINK, AI_BASEURL
- Database: UsersMap (presumably a map of chat IDs to user objects)

### Code Summary:

#### HelpCommandMessage:

This function handles the help command message. It takes a `tgbotapi.Message` object as input and sends a help message to the user.

#### SearchDocuments:

This function performs semantic search on a given prompt. It takes the chat ID, prompt, and maximum number of results as input. It first loads environment variables and retrieves the vector store from the database. Then, it performs the semantic search using the embeddings library and sends the results to the user.

#### RAG:

This function performs Retrieval-Augmented Generation (RAG). It takes the chat ID, prompt, and maximum number of results as input. It loads environment variables, retrieves the vector store, and calls the RAG function from the embeddings library. Finally, it sends the result to the user.

#### GetUsage:

This function retrieves and displays the usage statistics for a user. It takes the chat ID as input and retrieves the user object from the database. It then extracts the prompt tokens, completion tokens, and total tokens from the user object and sends them to the user.

#### SendMediaHelper:

This function sends a random video from the media directory to the user. It first reads the files in the media directory and selects a random file. Then, it opens the video file, creates a new video message, and sends it to the user.



lib/bot/command/addAdminTomap.go
## Package: command

### Imports:
- log
- db "github.com/JackBekket/hellper/lib/database"
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data:
- msgTemplates (not shown in the code, but mentioned in the code)

### Code Summary:
#### AddAdminToMap function:
This function is responsible for adding a new admin to the database and sending a confirmation message to the user. It takes two arguments: adminKey (the API key for the admin's GPT model) and updateMessage (a telegram message object containing information about the user).

1. It extracts the chatID from the updateMessage and creates a new User object with the chatID, username, dialog status, admin status, and AI session information (including the adminKey).

2. It adds the new User object to the UsersMap (a database or in-memory storage for users).

3. It logs a message indicating that the admin has been authorized.

4. It sends a confirmation message to the admin using the bot.

5. It sends another message to the admin with a one-time reply keyboard containing a button for selecting the GPT model (in this case, only GPT-3.5 is available).

lib/bot/command/cases.go


lib/bot/command/checkAdmin.go
## Package: command

### Imports:
- fmt
- github.com/JackBekket/hellper/lib/bot/env
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External data, input sources:
- adminData: map[string]env.AdminData
- updateMessage: *tgbotapi.Message

### CheckAdmin function:
This function checks if the user is an admin and updates the "dialogStatus" in the database accordingly. It iterates through the adminData map and compares the chatID of the updateMessage with the ID of each admin. If a match is found, it checks if the admin has a valid GPTKey. If the GPTKey is not empty, the admin is added to the admin map and the function returns. Otherwise, a message is sent to the user informing them that the environment variable is missing, and the user is added to the user map. If no match is found in the adminData map, the user is added to the user map.

lib/bot/command/newCommander.go
## Package: command

### Imports:
- `context`
- `github.com/JackBekket/hellper/lib/database`
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`

### External Data, Input Sources:
- `bot`: Telegram bot API instance
- `usersDb`: Map of user IDs to database.User objects
- `ctx`: Context for the command execution

### Commander struct:
- `bot`: Telegram bot API instance
- `usersDb`: Map of user IDs to database.User objects
- `ctx`: Context for the command execution

### NewCommander function:
- Creates a new Commander instance with the provided bot, users database, and context.

### GetCommander function:
- (Incomplete)



