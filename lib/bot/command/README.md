## Package: command

### Imports:

- log
- db "github.com/JackBekket/hellper/lib/database"
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:

- adminKey (string)
- updateMessage (tgbotapi.Message)

### AddAdminToMap Function:

This function is responsible for adding an admin to the database and sending a confirmation message to the user. It takes two arguments: adminKey (string) and updateMessage (tgbotapi.Message).

1. It extracts the chatID from the updateMessage.
2. It creates a new User object with the chatID, username, dialog status, admin status, and AI session information. The AI session information includes the adminKey.
3. It adds the new User object to the UsersMap in the database.
4. It logs a message indicating that the admin has been authorized.
5. It sends a confirmation message to the admin using the bot.
6. It sends another message with a one-time reply keyboard containing a button for selecting the GPT-3.5 model.

lib/bot/command/addNewUsertoMap.go
## Package: command

### Imports:

- log
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External Data, Input Sources:

- updateMessage: A tgbotapi.Message object containing information about the incoming Telegram message.

### AddNewUserToMap Function:

This function is responsible for adding a new user to the database and assigning them a "Dialog_status" of 0. It takes an updateMessage as input, extracts the chatID and username from the message, and creates a new User object with these values. The function then calls the database.AddUser function to store the new user in the database.

After adding the user to the database, the function logs the user's ID and username. It then creates a new message using the msgTemplates["hello"] template and sends it to the user with a one-time reply keyboard containing a "Start!" button.

The function also includes a commented-out section that appears to be related to registration and checking if the user is already registered. This section is not currently being used.

lib/bot/command/cases.go
## Package: command

This package contains the core logic for handling user interactions and managing the AI session. It includes functions for handling user input, managing the dialog flow, and interacting with the AI model.

### Imports:

- context
- fmt
- log
- strings
- github.com/JackBekket/hellper/lib/database
- github.com/JackBekket/hellper/lib/langchain
- github.com/JackBekket/hellper/lib/localai
- github.com/JackBekket/hellper/lib/localai/audioRecognition
- github.com/JackBekket/hellper/lib/localai/imageRecognition
- github.com/go-telegram-bot-api/telegram-bot-api/v5
- github.com/joho/godotenv

### External Data and Input Sources:

- Database: The package uses a database (likely a key-value store) to store user data, including their AI session information and dialog status.
- Telegram Bot API: The package interacts with the Telegram Bot API to receive user messages and send responses.
- Local AI: The package utilizes local AI models for tasks like speech recognition and image recognition.

### Major Code Parts:

1. Dialog Flow Management:
   - The package manages the dialog flow by tracking the user's dialog status and providing appropriate responses based on the current status.
   - Functions like `InputYourAPIKey`, `ChooseModel`, `HandleModelChoose`, and `ConnectingToAiWithLanguage` handle different stages of the dialog flow.

2. AI Model Interaction:
   - The package allows users to choose an AI model from a list of available options.
   - The `attachModel` function attaches the selected model to the user's session.
   - The `DialogSequence` function handles the main loop for interacting with the AI model, processing user input and sending responses.

3. User Data Management:
   - The package provides functions to access and manage user data, such as `GetUsersDb` and `GetUser`.
   - User data is stored in the database and includes information about the user's AI session and dialog status.

4. Local AI Integration:
   - The package integrates with local AI models for tasks like speech recognition and image recognition.
   - Functions like `HandleVoiceMessage` and `RecognizeImage` handle these tasks and provide responses to the user.

5. Error Handling:
   - The package includes error handling mechanisms to handle potential issues, such as invalid API keys or failed AI model interactions.
   - Functions like `WrongResponse` provide appropriate responses to the user in case of errors.

6. Context Management:
   - The package uses the `context` package to manage the context of the AI session, including user data and other relevant information.

7. Environment Variables:
   - The package uses environment variables to configure the local AI endpoint and other settings.

8. File Handling:
   - The package handles file operations, such as deleting temporary files, to ensure proper cleanup.

9. Logging:
   - The package uses the `log` package to log important events and debug information.

10. Helper Functions:
   - The package includes helper functions for tasks like deleting files and managing dialog status.

By summarizing these major code parts, we can understand the overall functionality of the `command` package and its role in managing user interactions and AI session management.

lib/bot/command/checkAdmin.go
## Package: command

### Imports:

- fmt
- github.com/JackBekket/hellper/lib/bot/env
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External Data, Input Sources:

- adminData: A map of strings to env.AdminData, representing the admin data for each environment.
- updateMessage: A tgbotapi.Message, representing the incoming Telegram message.

### Summary:

#### CheckAdmin Function:

The CheckAdmin function is responsible for updating the "dialogStatus" in the database based on whether the user is an admin or not. It takes two arguments: adminData and updateMessage.

1. It first extracts the chatID from the updateMessage.
2. Then, it iterates through the adminData map, checking if the chatID matches any of the admin entries.
3. If a match is found, it checks if the admin's GPTKey is not empty. If it is, the function calls the AddAdminToMap function, passing the GPTKey and updateMessage as arguments.
4. If the GPTKey is empty, the function sends a message to the chatID, indicating that the environment is missing. It then calls the AddNewUserToMap function, passing the updateMessage as an argument.
5. If no match is found in the adminData map, the function calls the AddNewUserToMap function, passing the updateMessage as an argument.

The AddAdminToMap and AddNewUserToMap functions are not shown in the provided code, but they are likely responsible for updating the database with the appropriate dialogStatus based on whether the user is an admin or not.

lib/bot/command/msgTemplates.go
## Package/Component: command

### Imports:

```
map[string]string
```

### External Data, Input Sources:

- `msgTemplates`: A map containing various message templates used in the package.

### Summary:

The provided code snippet is part of a package or component named "command". It defines a map called `msgTemplates` that stores various message templates. These templates are likely used for generating responses or displaying information to the user. The code snippet also includes a comment that suggests the package or component might be related to a local AI node and provides a list of additional commands that can be used with the bot.

lib/bot/command/newCommander.go
## Package: command

### Imports:
- context
- github.com/JackBekket/hellper/lib/database
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External Data, Input Sources:
- `bot`: A pointer to a `tgbotapi.BotAPI` object, which is used to interact with the Telegram Bot API.
- `usersDb`: A map of user IDs to `database.User` objects, which is used to store and retrieve user data.
- `ctx`: A context object, which is used to manage the lifetime of the Commander instance.

### Commander Struct:
The `Commander` struct is responsible for managing the interaction with the Telegram Bot API and the database. It has the following fields:
- `bot`: A pointer to a `tgbotapi.BotAPI` object.
- `usersDb`: A map of user IDs to `database.User` objects.
- `ctx`: A context object.

### NewCommander Function:
The `NewCommander` function is used to create a new `Commander` instance. It takes the following arguments:
- `bot`: A pointer to a `tgbotapi.BotAPI` object.
- `usersDb`: A map of user IDs to `database.User` objects.
- `ctx`: A context object.

The function returns a pointer to a new `Commander` instance, which is initialized with the provided arguments.

### GetCommander Function:
The `GetCommander` function is not provided in the code snippet. It is assumed to be a function that returns a pointer to a `Commander` instance.

lib/bot/command/ui.go
## Package: command

### Imports:
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:
- msgTemplates: A map containing message templates for different scenarios.

### Code Summary:
#### RenderModelMenuLAI:
This function renders a menu of LLaMA-based models with an inline keyboard. It takes the chat ID as input and constructs a message with the appropriate model names as buttons. The message is then sent to the specified chat using the bot.

#### RenderLanguage:
This function renders a menu for choosing a language with an inline keyboard. It takes the chat ID as input and constructs a message with language options as buttons. The message is then sent to the specified chat using the bot.



lib/bot/command/utils.go