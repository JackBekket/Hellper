# lib/bot/command/addAdminTomap.go  
## Package: command  
  
### Imports:  
- log  
- db "github.com/JackBekket/hellper/lib/database"  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data:  
- msgTemplates (not shown in the code, but mentioned in the code)  
  
### Summary:  
#### AddAdminToMap function:  
This function is responsible for adding a new admin to the system. It takes two arguments: adminKey (the API key for the admin's GPT model) and updateMessage (a telegram message object containing information about the user).  
  
1. It extracts the chatID from the updateMessage and creates a new User object with the chatID, username, dialog status, admin status, and AI session information (including the adminKey).  
  
2. It stores the new User object in the UsersMap (which is part of the database package).  
  
3. It logs a message indicating that the admin has been authorized.  
  
4. It sends a message to the admin confirming their authorization.  
  
5. It sends another message to the admin with a one-time reply keyboard containing a button for selecting the GPT-3.5 model.  
  
# lib/bot/command/addNewUsertoMap.go  
## Package: command  
  
### Imports:  
  
- log  
- github.com/JackBekket/hellper/lib/database  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
### External Data, Input Sources:  
  
- updateMessage: A tgbotapi.Message object containing information about the incoming Telegram message.  
  
### Summary:  
  
#### AddNewUserToMap function:  
  
This function is responsible for adding a new user to the database and assigning them a "Dialog_status" of 0. It takes an updateMessage as input, which contains information about the incoming Telegram message. The function first extracts the chatID and username from the updateMessage. Then, it creates a new User object with the extracted information and the Dialog_status set to 0. The User object is then added to the database using the AddUser function from the database package.  
  
After adding the user to the database, the function logs the user's ID and username. It then creates a new message using the "hello" template from the msgTemplates map and sends it to the user with a one-time reply keyboard containing a "Start!" button.  
  
The function also includes a commented-out section that checks if the user is already registered and updates the user's Dialog_status accordingly. However, this section is not currently being used.  
  
# lib/bot/command/cases.go  
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
   - The package includes error handling mechanisms to handle potential issues during the dialog flow and AI model interaction.  
   - Functions like `WrongResponse` and `ChangeDialogStatus` help manage errors and maintain the dialog flow.  
  
6. Environment Variables:  
   - The package uses environment variables to configure the local AI endpoint.  
   - The `AttachKey` function handles the process of attaching the user's API key to their session.  
  
7. Context Management:  
   - The package uses a context to store user data and pass it between functions.  
   - The `contextKey` type is used to define the key for storing user data in the context.  
  
8. Helper Functions:  
   - The package includes helper functions like `DeleteFile` and `GetEnvsForSST` to support the main functionality.  
  
By summarizing these major code parts, we can understand the overall structure and functionality of the command package. It provides a comprehensive framework for managing user interactions, handling dialog flow, and integrating with AI models for various tasks.  
  
# lib/bot/command/checkAdmin.go  
## Package: command  
  
### Imports:  
- fmt  
- github.com/JackBekket/hellper/lib/bot/env  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
### External Data, Input Sources:  
- adminData: A map of strings to env.AdminData, representing the admin data for each environment.  
- updateMessage: A tgbotapi.Message, representing the incoming Telegram message.  
  
### CheckAdmin Function:  
This function is responsible for checking if the user is an admin and updating the dialog status in the database accordingly. It takes two arguments: adminData and updateMessage.  
  
1. It first extracts the chatID from the updateMessage.  
2. Then, it iterates through the adminData map, looking for a matching chatID.  
3. If a match is found, it checks if the admin has a valid GPTKey.  
    - If the GPTKey is not empty, it adds the admin to the c.adminMap and returns.  
    - If the GPTKey is empty, it sends a message to the user informing them that the environment variable is missing and then adds the user to the c.userMap.  
4. If no match is found in the adminData map, it adds the user to the c.userMap.  
  
# lib/bot/command/msgTemplates.go  
package command  
  
imports:  
- fmt  
- strings  
  
external data:  
- msgTemplates: map[string]string  
  
summary:  
The code defines a package called "command" that provides a set of message templates for various commands and functionalities. The package includes a map called "msgTemplates" that stores string values for different commands, such as "hello," "case0," "await," "case1," and "help_command." These templates can be used to generate messages for users or other components within the system.  
  
The "help_command" template provides a list of available commands and their descriptions, including:  
- /help: Print this message.  
- /restart: Restart the session to switch between local-ai and openai chatGPT.  
- /search_doc: Search for documents.  
- /rag: Process Retrival-Augmented Generation.  
- /instruct: Use system prompt template instead of langchain (higher priority, see examples).  
- /image: Generate images.  
  
The code also includes a comment indicating that all functions are experimental and may cause the bot to halt or catch fire.  
  
  
  
# lib/bot/command/newCommander.go  
## Package: command  
  
### Imports:  
  
- `context`  
- `github.com/JackBekket/hellper/lib/database`  
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
### External Data and Input Sources:  
  
- `bot`: A pointer to a `tgbotapi.BotAPI` object, which is used to interact with the Telegram Bot API.  
- `usersDb`: A map of user IDs to `database.User` objects, which is used to store and retrieve user data from the database.  
- `ctx`: A context object, which is used to manage the lifetime of the Commander instance and its associated resources.  
  
### Summary of Major Code Parts:  
  
#### Commander Struct:  
  
The `Commander` struct represents the core component of the package, responsible for managing the interaction with the Telegram Bot API and the database. It contains the following fields:  
  
- `bot`: A pointer to the Telegram Bot API object.  
- `usersDb`: A map of user IDs to `database.User` objects.  
- `ctx`: A context object.  
  
#### NewCommander Function:  
  
The `NewCommander` function is a constructor for the `Commander` struct. It takes the Telegram Bot API object, the user database, and a context object as input and returns a new `Commander` instance.  
  
#### GetCommander Function:  
  
The `GetCommander` function is not provided in the code snippet, but it is mentioned in the comments. It is likely a function that returns a new `Commander` instance, similar to the `NewCommander` function.  
  
# lib/bot/command/ui.go  
## Package: command  
  
### Imports:  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data, Input Sources:  
- msgTemplates: A map containing message templates for different scenarios.  
  
### Code Summary:  
#### RenderModelMenuLAI:  
This function renders a menu of LLaMA-based models with an inline keyboard. It takes the chat ID as input and constructs a message with the appropriate model names as buttons. The message is then sent to the specified chat using the bot instance.  
  
#### RenderLanguage:  
This function renders a menu for choosing a language with an inline keyboard. It takes the chat ID as input and constructs a message with language options as buttons. The message is then sent to the specified chat using the bot instance.  
  
# lib/bot/command/utils.go  
  
  
