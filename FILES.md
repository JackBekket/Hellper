# lib/bot/command/addNewUsertoMap.go  
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
  
# lib/bot/command/cases.go  
  //TODO: find why it cannot procced with this file
  
# lib/bot/dialog/dialog.go  
## Package: dialog  
  
### Imports:  
- log  
- github.com/JackBekket/hellper/lib/bot/command  
- github.com/JackBekket/hellper/lib/database  
- github.com/JackBekket/hellper/lib/langchain  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data, Input Sources:  
- Updates from Telegram bot API  
- Command data from Telegram bot API  
- Database for user information  
  
### Summary:  
The `HandleUpdates` function is responsible for handling incoming updates from the Telegram bot API and managing user interactions. It iterates through the updates and processes each message based on the command provided by the user.  
  
#### Command Handling:  
- `/image`: Generates an image based on the provided prompt or a default prompt if none is given.  
- `/restart`: Restarts the user's session by deleting their data from the database.  
- `/help`: Displays a help message with available commands.  
- `/search_doc`: Searches for documents based on the provided prompt and returns the top 3 results.  
- `/rag`: Performs a RAG (Retrieval Augmented Generation) task based on the provided prompt.  
- `/instruct`: Calls a local AI model to generate content based on the provided prompt and user's AI session settings.  
- `/usage`: Displays the usage statistics for the user.  
- `/helper`: Sends a media helper message to the user.  
  
#### User Interaction:  
- The function checks if the user is already in the database and adds them if not.  
- It then determines the user's dialog status and handles the corresponding interaction based on the status.  
- The dialog status is updated as the user progresses through the interaction, and the function logs the user's status and other relevant information.  
  
#### Network and Model Selection:  
- The function allows the user to choose their preferred network and AI model.  
- It handles the selection process and updates the user's AI session settings accordingly.  
  
#### Connecting to AI:  
- The function connects the user to the chosen AI model using the provided API key and base URL.  
  
#### Dialog Sequence:  
- The function manages the dialog sequence by calling the appropriate functions based on the user's dialog status and input.  
  
This summary provides a comprehensive overview of the `HandleUpdates` function and its role in managing user interactions and dialog flow within the `dialog` package.  
  
# lib/embeddings/common.go  
## Package: embeddings  
  
### Imports:  
  
```  
context  
fmt  
log  
  
github.com/tmc/langchaingo/embeddings  
github.com/tmc/langchaingo/llms/openai  
github.com/tmc/langchaingo/vectorstores  
github.com/tmc/langchaingo/vectorstores/pgvector  
  
github.com/jackc/pgx/v5/pgxpool  
```  
  
### External Data, Input Sources:  
  
1. Environment variables:  
    - `AI_BASEURL`: Base URL for the AI service (e.g., localhost, OpenAI, Docker).  
    - `OPENAI_API_KEY`: API key for the OpenAI service.  
    - `PG_HOST`: Hostname for the PostgreSQL database.  
    - `PG_USER`: Username for the PostgreSQL database.  
    - `PG_PASSWORD`: Password for the PostgreSQL database.  
    - `PG_DB`: Database name for the PostgreSQL database.  
2. Database connection string: `db_link` - A string containing the connection details for the PostgreSQL database.  
  
### Code Summary:  
  
#### LoadEnv() function:  
  
- This function is not implemented in the provided code.  
  
#### GetVectorStore() function:  
  
1. Retrieves the base URL for the AI service from the `ai_url` parameter.  
2. Retrieves the API token for the AI service from the `api_token` parameter.  
3. Retrieves the database connection string from the `db_link` parameter.  
4. Creates a PostgreSQL connection pool using the provided database connection string.  
5. Creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and API token.  
6. Creates an embeddings embedder using the OpenAI embeddings client.  
7. Creates a vector store using the PostgreSQL connection pool and the embeddings embedder.  
8. Returns the created vector store and any errors encountered during the process.  
  
#### GetVectorStoreWithOptions() function:  
  
1. Retrieves the base URL for the AI service from the `ai_url` parameter.  
2. Retrieves the API token for the AI service from the `api_token` parameter.  
3. Retrieves the database connection string from the `db_link` parameter.  
4. Retrieves the name of the collection for the vector store from the `name` parameter.  
5. Creates a PostgreSQL connection pool using the provided database connection string.  
6. Creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and API token.  
7. Creates an embeddings embedder using the OpenAI embeddings client.  
8. Creates a vector store using the PostgreSQL connection pool, embeddings embedder, and the specified collection name.  
9. Returns the created vector store and any errors encountered during the process.  
  
  
  
# lib/localai/localai.go  
## Package: localai  
  
### Imports:  
  
```  
bytes  
encoding/json  
fmt  
io/ioutil  
log  
net/http  
os  
path/filepath  
github.com/StarkBotsIndustries/telegraph  
```  
  
### External Data, Input Sources:  
  
1. The package uses a local API endpoint at `http://localhost:8080/v1/chat/completions` for generating text completions.  
2. It also uses a local API endpoint at `http://localhost:8080/v1/images/generations` for generating images using Stable Diffusion.  
3. The package uses the `telegraph` library for uploading images to Telegram.  
  
### Code Summary:  
  
#### Chat Completion:  
  
The package provides a function `GenerateCompletion` that takes a prompt, model name, and API URL as input. It constructs a JSON request body with the prompt, model name, and temperature, and sends a POST request to the specified API endpoint. The response is parsed as a `ChatResponse` object, which contains the generated text completion.  
  
#### Image Generation:  
  
The package provides a function `GenerateImageStableDissusion` that takes a prompt and image size as input. It constructs a JSON request body with the prompt and size, and sends a POST request to the specified API endpoint. The response is parsed as a `GenerationResponse` object, which contains the generated image URL.  
  
#### Image Upload:  
  
The package provides a function `UploadToTelegraph` that takes a file path as input. It opens the file, uploads it to Telegram using the `telegraph` library, and returns the uploaded image URL.  
  
#### Wrong Password Handling:  
  
The package provides a function `GenerateCompletionWithPWD` that takes a prompt, model name, API URL, and two passwords as input. It checks if the passwords match, and if they do, it calls the `GenerateCompletion` function to generate the text completion. If the passwords don't match, it returns an error.  
  
#### File Deletion:  
  
The package provides a function `deleteFromTemp` that takes a file name as input. It deletes the file from the temporary directory.  
  
  
  
# lib/bot/command/addAdminTomap.go  
## Package: command  
  
### Imports:  
- log  
- db "github.com/JackBekket/hellper/lib/database"  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data:  
- msgTemplates (not shown in the code, but mentioned in the code)  
  
### Code Summary:  
#### AddAdminToMap function:  
This function is responsible for adding a new admin to the system. It takes two arguments: adminKey (the API key for the admin's GPT model) and updateMessage (a telegram message object containing information about the user).  
  
1. It extracts the chatID from the updateMessage and creates a new User object with the chatID, username, dialog status, admin status, and AI session information. The AI session contains the adminKey.  
  
2. It stores the new User object in the UsersMap, which is a map of chatID to User objects.  
  
3. It logs a message indicating that the admin has been authorized.  
  
4. It sends a message to the admin confirming their authorization.  
  
5. It sends another message to the admin with a one-time reply keyboard containing a button for selecting the GPT-3.5 model.  
  
# lib/bot/command/msgTemplates.go  
Package: command  
  
Imports:  
  
External data, input sources:  
  
- msgTemplates: A map containing various message templates used in the package.  
  
Summary:  
  
The provided code defines a map called `msgTemplates` that stores various message templates used within the package. These templates are used to generate responses and messages for different scenarios, such as greetings, instructions, and error messages. The map contains key-value pairs, where the key is a string representing the template name, and the value is the corresponding message template.  
  
The templates include:  
  
- "hello": A greeting message indicating that the bot is working with a local AI node.  
- "case0": Instructions for users to input their OpenAI API key or password for local AI authentication.  
- "await": A message indicating that the bot is awaiting a response or input.  
- "case1": A prompt for users to choose a model to use.  
- "ch_network": A prompt for users to choose a network to work with, either OpenAI or local AI.  
- "help_command": A list of available commands for users, including /help, /restart, /search_doc, /rag, /instruct, and /image.  
  
The code snippet provides a concise and organized way to manage and access various message templates within the package, ensuring consistency and clarity in the bot's communication with users.  
  
  
  
# lib/bot/command/ui.go  
Package: command  
  
Imports:  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
External data, input sources:  
- msgTemplates: A map of message templates used in the code.  
  
Summary:  
### RenderModelMenuOAI  
This function renders a menu for choosing an OAI model. It creates a new message with a predefined template and a one-time reply keyboard containing buttons for "gpt-3.5" and "gpt-4". The message is sent to the specified chat ID using the bot.  
  
### RenderModelMenuLAI  
This function renders a menu for choosing an LAI model. It creates a new message with a predefined template and a one-time reply keyboard containing buttons for "wizard-uncensored-13b", "wizard-uncensored-30b", and "tiger-gemma-9b-v1-i1". The message is sent to the specified chat ID using the bot.  
  
### RenderModelMenuVAI  
This function renders a menu for choosing a VAI model. It creates a new message with a predefined template and a one-time reply keyboard containing buttons for "deepseek-coder-6b-instruct", "wizard-uncensored-code-34b", "tiger-gemma-9b-v1-i1", and "big-tiger-gemma-27b-v1". The message is sent to the specified chat ID using the bot.  
  
### RenderLanguage  
This function renders a menu for choosing a language. It creates a new message with a predefined template and a one-time reply keyboard containing buttons for "English" and "Russian". The message is sent to the specified chat ID using the bot.  
  
  
  
# lib/bot/command/utils.go  
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
  
  
  
# lib/database/newAiSessionDataBase.go  
Package: database  
  
Imports:  
- gogpt  
  
External data, input sources:  
- AiSessionMap: A map that stores AiSession objects, where the key is an int64 and the value is an AiSession object.  
  
Summary:  
The code defines a map called AiSessionMap, which stores AiSession objects. Each AiSession object has three fields: GptKey, GptClient, and GptModel. The GptKey field is a string, the GptClient field is a pointer to a gogpt.Client object, and the GptModel field is a string. The AiSessionMap is initialized as an empty map.  
  
  
  
# lib/langchain/handler.go  
## Package: langchain  
  
### Imports:  
  
```  
context  
encoding/json  
log  
  
github.com/JackBekket/hellper/lib/database  
github.com/tmc/langchaingo/llms  
github.com/tmc/langchaingo/schema  
```  
  
### External Data, Input Sources:  
  
- Database: The code uses a database (likely a relational database like PostgreSQL or MySQL) to store user information and session usage data. The database package used is `github.com/JackBekket/hellper/lib/database`.  
- LLM: The code interacts with a large language model (LLM) through the `llms` package, which is likely a wrapper around a popular LLM like OpenAI's GPT-3 or HuggingFace's Transformers.  
  
### Code Summary:  
  
#### ChainCallbackHandler:  
  
This struct implements a callback handler for various events during the execution of a chain of actions. It includes methods for handling agent actions, agent finishes, chain ends, chain errors, chain starts, LLM errors, LLM content generation starts, LLM starts, retriever ends, retriever starts, streaming functions, tool ends, tool errors, and tool starts.  
  
#### HandleLLMGenerateContentEnd:  
  
This method is called when the LLM has finished generating content. It logs the content, stop reason, context, and generation information. It also updates the user's session usage based on the number of prompt tokens, completion tokens, and total tokens used.  
  
#### LogResponseContentChoice:  
  
This helper function logs the content, stop reason, context, and generation information for a given LLM response. It also updates the user's session usage based on the number of prompt tokens, completion tokens, and total tokens used.  
  
#### Other Methods:  
  
The remaining methods in the `ChainCallbackHandler` struct are not implemented in the provided code. They are likely placeholders for future functionality related to handling various events during the execution of a chain of actions.  
  
  
  
# main.go  
## Package: hellper/lib/bot  
  
### Imports:  
  
- context  
- log  
- os  
- strconv  
- github.com/JackBekket/hellper/lib/bot/command  
- github.com/JackBekket/hellper/lib/bot/dialog  
- github.com/JackBekket/hellper/lib/bot/env  
- github.com/JackBekket/hellper/lib/database  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
- github.com/joho/godotenv  
  
### External Data, Input Sources:  
  
- OPENAI_API_KEY (local key for localai)  
- PG_LINK (not used in the code)  
- TG_KEY (telegram bot token)  
- ADMIN_ID (admin user ID)  
- AI_ENDPOINT (local AI endpoint)  
  
### Summary:  
  
#### Initialization:  
  
1. Loads environment variables using godotenv.Load() and retrieves values for OPENAI_API_KEY, TG_KEY, ADMIN_ID, and AI_ENDPOINT.  
2. Creates a new instance of the Telegram bot using the retrieved TG_KEY.  
3. Initializes a map called adminData to store admin information, including the admin ID and GPT key.  
  
#### Database and Commander:  
  
1. Initializes a database for storing user information using the database.UsersMap.  
2. Creates a new instance of the command.Commander, which handles bot commands and user interactions.  
  
#### Update Handling:  
  
1. Sets up a channel for handling incoming updates from the Telegram bot.  
2. Starts a goroutine to handle updates using the dialog.HandleUpdates function.  
3. Iterates through incoming updates and checks if the user is new. If so, the user is added to the database.  
  
#### Main Loop:  
  
1. Continuously listens for new updates from the Telegram bot.  
2. For each update, checks if the user is new and adds them to the database if necessary.  
3. Sends the update to the update handler goroutine for processing.  
  
# lib/bot/command/newCommander.go  
## Package: command  
  
### Imports:  
- `context`  
- `github.com/JackBekket/hellper/lib/database`  
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
### External Data and Input Sources:  
- `bot`: A Telegram bot API instance.  
- `usersDb`: A map of user IDs to database.User objects.  
- `ctx`: A context object.  
  
### Commander struct:  
- The `Commander` struct holds the bot, user database, and context.  
  
### NewCommander function:  
- This function creates a new Commander instance, taking the bot, user database, and context as input.  
  
### GetCommander function:  
- This function is not fully implemented in the provided code.  
  
# lib/database/newUserDataBase.go  
## Package: database  
  
### Imports:  
- github.com/tmc/langchaingo/chains  
- github.com/tmc/langchaingo/memory  
  
### External Data, Input Sources:  
- UsersMap: A map that stores user data, where the key is the Telegram user ID (int64) and the value is a User struct.  
- UsageMap: A map that stores session usage data, where the key is the Telegram user ID (int64) and the value is a SessionUsage struct.  
  
### Code Summary:  
  
#### User Struct:  
- Represents a user in the database.  
- Contains fields for user ID, username, dialog status, admin status, AI session, and network.  
  
#### SessionUsage Struct:  
- Represents a session usage record.  
- Contains fields for session ID and usage data.  
  
#### AiSession Struct:  
- Represents an AI session.  
- Contains fields for GPT key, GPT model, AI type, dialog thread, base URL, and usage data.  
  
#### ChatSession Struct:  
- Represents a chat session.  
- Contains fields for conversation buffer and dialog thread.  
  
#### AddUser Function:  
- Adds a new user to the UsersMap.  
  
#### UpdateUserUsage Function:  
- Updates the usage data for a specific user in the UsersMap.  
  
#### UpdateSessionUsage Function:  
- Updates the session usage data for a specific user in the UsageMap.  
  
#### GetSessionUsage Function:  
- Retrieves the session usage data for a specific user from the UsageMap.  
  
# lib/embeddings/load.go  
## Package: embeddings  
  
### Imports:  
  
```  
context  
fmt  
net/http  
log  
github.com/tmc/langchaingo/documentloaders  
github.com/tmc/langchaingo/schema  
github.com/tmc/langchaingo/textsplitter  
github.com/tmc/langchaingo/vectorstores  
```  
  
### External Data and Input Sources:  
  
1. The code uses the `http.Get` function to fetch data from a given URL.  
2. It utilizes the `documentloaders` package to load and split the fetched data into documents.  
3. The `textsplitter` package is used to split the documents into smaller units.  
4. The code also uses the `vectorstores` package to store the processed documents in a vector store.  
  
### Code Summary:  
  
#### LoadDocsToStore Function:  
  
This function takes a list of documents and a vector store as input. It first prints a message indicating that data is being loaded from the given source. Then, it iterates through the list of documents and adds them to the vector store using the `AddDocuments` method. If any errors occur during the process, the function logs the error and panics. Finally, it prints a message confirming that the data has been successfully loaded into the vector store.  
  
#### getDocs Function:  
  
This function takes a URL as input and returns a list of documents and an error. It first uses the `http.Get` function to fetch the data from the given URL. Then, it uses the `documentloaders` package to load and split the fetched data into documents. Finally, it returns the list of documents and any errors that occurred during the process.  
  
#### Other Code:  
  
The code also includes a commented-out function called `GetTextDocs` which is not used in the current implementation. This function was likely intended to load documents from a local file or other data source, but it is not currently being used.  
  
# lib/localai/setupSequenceWithKey.go  
## Package: localai  
  
### Imports:  
  
```  
import (  
	"context"  
	"log"  
	"sync"  
  
	db "github.com/JackBekket/hellper/lib/database"  
	//"github.com/sashabaranov/go-openai"  
  
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
)  
```  
  
### External Data, Input Sources:  
  
- Database: `db.User` struct, which contains user information, including their AI session data.  
- Telegram Bot API: `tgbotapi.BotAPI` for interacting with the Telegram bot.  
- AI Endpoint: `ai_endpoint` for communication with the AI model.  
- Password: `spwd` for authentication with the AI model.  
- User's GPT Key: `gptKey` from the user's AI session.  
  
### Code Summary:  
  
#### SetupSequenceWithKey Function:  
  
This function is responsible for setting up the sequence with the user's GPT key and language. It takes the following parameters:  
  
- `bot`: Telegram bot API instance.  
- `user`: Database user object containing user information.  
- `language`: User's preferred language.  
- `ctx`: Context for the operation.  
- `spwd`: Password for authentication with the AI model.  
- `ai_endpoint`: URL of the AI model endpoint.  
  
The function first locks a mutex to ensure thread safety. Then, it retrieves the user's GPT key and model from the `user.AiSession` field. It also retrieves the user's ID and network information.  
  
The function then uses a switch statement to handle different language cases: English, Russian, and default. For each case, it calls the `tryLanguage` function to generate a response based on the user's language preference. If an error occurs during the process, it calls the `errorMessage` function to handle the error. Otherwise, it sends the generated response to the user via the Telegram bot and updates the user's dialog status.  
  
#### tryLanguage Function:  
  
This function takes the user's language preference, language code, context, AI endpoint, password, and user's GPT key as input. It constructs a language prompt based on the language code and calls the `GenerateCompletionWithPWD` function to generate a response from the AI model. The function then logs the response and returns the generated answer.  
  
#### GenerateCompletionWithPWD Function:  
  
This function is responsible for generating a response from the AI model using the provided prompt, model, AI endpoint, password, and user's GPT key. It returns the generated response and any errors that may have occurred during the process.  
  
#### LogResponse Function:  
  
This function logs the generated response from the AI model.  
  
#### errorMessage Function:  
  
This function handles errors that may occur during the process and sends an error message to the user via the Telegram bot.  
  
  
  
# lib/bot/command/checkAdmin.go  
## Package: command  
  
### Imports:  
- fmt  
- github.com/JackBekket/hellper/lib/bot/env  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
### External data, input sources:  
- adminData: map[string]env.AdminData  
- updateMessage: *tgbotapi.Message  
  
### CheckAdmin function:  
This function checks if the user is an admin and updates the "dialogStatus" in the database accordingly. It iterates through the adminData map and compares the chatID of the updateMessage with the ID of each admin. If a match is found, it checks if the admin has a valid GPTKey. If the GPTKey is not empty, the function adds the admin to the map and returns. Otherwise, it sends a message to the user informing them that the environment variable is missing and directs to the case 0, which is adding the user to the map as a regular user. If no match is found in the adminData map, the function adds the user to the map as a regular user.  
  
# lib/bot/env/newEvn.go  
## Package: env  
  
### Imports:  
  
```  
errors  
log  
strconv  
github.com/joho/godotenv  
```  
  
### External Data, Input Sources:  
  
- `.env` file: Contains environment variables that are loaded and used by the package.  
  
### Code Summary:  
  
#### Load() function:  
  
- Loads environment variables from the `.env` file using the `godotenv` package.  
- Returns an error if there is an issue loading the environment variables.  
  
#### LoadAdminData() function:  
  
- Creates a map of AdminData structs, where each key is an admin identifier (e.g., "ADMIN_ID", "MINTY_ID") and the value is an AdminData struct containing the admin ID and GPT key.  
- Iterates through the loaded environment variables and parses the values for each admin identifier.  
- Parses the admin ID as an integer and the GPT key as a string.  
- Stores the parsed data in the AdminData struct and adds it to the map.  
  
#### LoadTGToken() function:  
  
- Retrieves the Telegram token from the environment variables.  
- Returns an error if the Telegram token is not found in the `.env` file.  
  
#### LoadLocalPD() function:  
  
- Retrieves the local password from the environment variables.  
  
#### LoadLocalAI_Endpoint() function:  
  
- Retrieves the local AI endpoint from the environment variables.  
  
#### GetAdminToken() function:  
  
- Retrieves the admin token from the environment variables.  
  
# lib/embeddings/query.go  
## Package: embeddings  
  
### Imports:  
  
```  
"context"  
"fmt"  
"github.com/tmc/langchaingo/chains"  
"github.com/tmc/langchaingo/llms/openai"  
"github.com/tmc/langchaingo/schema"  
"github.com/tmc/langchaingo/vectorstores"  
```  
  
### External Data, Input Sources:  
  
1. `ai_url`: URL of the AI service (e.g., OpenAI API).  
2. `api_token`: API token for authentication with the AI service.  
3. `question`: The query to be answered or searched for.  
4. `numOfResults`: The number of results to return.  
5. `store`: A vector store to store and retrieve embeddings.  
6. `searchQuery`: The query for semantic search.  
7. `maxResults`: The maximum number of results to return for semantic search.  
8. `options`: Additional options for the vector store.  
  
### Code Summary:  
  
#### Rag Function:  
  
The `Rag` function performs a retrieval-augmented generation (RAG) task using a language model (LLM) and a vector store. It takes the AI service URL, API token, question, number of results, and vector store as input. The function first creates an embeddings client using the provided AI service URL, API token, and model names. Then, it runs a retrieval-augmented generation chain using the LLM and vector store to generate a response to the question. Finally, it returns the generated response and any errors encountered during the process.  
  
#### SemanticSearch Function:  
  
The `SemanticSearch` function performs a semantic search using a vector store. It takes the search query, maximum number of results, vector store, and additional options as input. The function first retrieves the vector store (if not provided) and then performs a similarity search using the provided search query and maximum number of results. Finally, it returns the search results and any errors encountered during the process.  
  
  
  
# lib/langchain/langchain.go  
## Package: langchain_controller  
  
### Imports:  
  
```  
import (  
	"context"  
	"fmt"  
	"log"  
  
	db "github.com/JackBekket/hellper/lib/database"  
  
	"github.com/tmc/langchaingo/chains"  
	"github.com/tmc/langchaingo/llms"  
	"github.com/tmc/langchaingo/memory"  
  
	//"github.com/tmc/langchaingo/llms/options"  
	"github.com/tmc/langchaingo/llms/openai"  
)  
```  
  
### External Data, Input Sources:  
  
- API token for OpenAI or local AI model  
- Model name for OpenAI or local AI model  
- Base URL for local AI model (if applicable)  
- User initial prompt  
  
### Code Summary:  
  
#### InitializeNewChatWithContextNoLimit:  
  
This function initializes a new chat session with a given API token, model name, base URL (for local AI), and user initial prompt. It creates a new conversation using the specified LLM (OpenAI or local AI) and a memory buffer to store the conversation history. The function returns a new chat session object and an error if any.  
  
#### StartNewChat:  
  
This function starts a new conversation by calling InitializeNewChatWithContextNoLimit and then running the conversation using the RunChain function. It returns the result of the conversation, the updated chat session object, and an error if any.  
  
#### RunChain:  
  
This function runs a given prompt through the provided chat session and returns the result and the updated chat session object.  
  
#### ContinueChatWithContextNoLimit:  
  
This function continues an existing conversation by running a given prompt through the provided chat session and returns the result and the updated chat session object.  
  
#### GenerateContentInstruction:  
  
This function generates content from a single prompt without using memory or context. It takes the base URL, prompt, model name, API token, network (local or openai), and additional options as input. It creates an LLM instance based on the network and model name, and then uses the llms.GenerateFromSinglePrompt function to generate the content. The function returns the generated content and an error if any.  
  
  
  
# lib/langchain/setupSequenceWithKey.go  
## Package: langchain  
  
### Imports:  
  
```  
context  
log  
sync  
db "github.com/JackBekket/hellper/lib/database"  
tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
```  
  
### External Data, Input Sources:  
  
1. Database: The code uses a database (likely a local database) to store user information, including their AI session data. The database is accessed through the `db` package.  
2. Telegram Bot API: The code uses the Telegram Bot API (provided by the `tgbotapi` package) to interact with a Telegram bot. This suggests that the code is part of a Telegram bot application.  
3. AI Endpoint: The code uses an AI endpoint (specified by the `ai_endpoint` parameter) to interact with an AI model. This endpoint could be a local API or a remote service.  
  
### Code Summary:  
  
#### SetupSequenceWithKey Function:  
  
This function is responsible for setting up a sequence of interactions with an AI model for a given user. It takes the following parameters:  
  
1. `bot`: A pointer to a Telegram bot API instance.  
2. `user`: A `db.User` struct containing user information, including their AI session data.  
3. `language`: A string representing the language in which the user wants to interact with the AI model.  
4. `ctx`: A context object for managing the execution of the function.  
5. `ai_endpoint`: A string representing the URL or address of the AI endpoint.  
  
The function first locks a mutex to ensure thread safety. Then, it extracts the user's GPT key, model, and other relevant information from the `user` struct. Based on the provided language, it calls the `tryLanguage` function to initiate a conversation with the AI model. The `tryLanguage` function returns a response, a chat session thread, and an error. If there is an error, the function calls an `errorMessage` function to handle the error. Otherwise, it sends the response to the user via the Telegram bot, updates the user's dialog status, and stores the updated user information in the database.  
  
#### tryLanguage Function:  
  
This function is responsible for initiating a conversation with the AI model based on the provided language. It takes the following parameters:  
  
1. `user`: A `db.User` struct containing user information.  
2. `language`: A string representing the language in which the user wants to interact with the AI model.  
3. `languageCode`: An integer representing the language code (0 - default, 1 - Russian, 2 - English).  
4. `ctx`: A context object for managing the execution of the function.  
5. `ai_endpoint`: A string representing the URL or address of the AI endpoint.  
  
The function first constructs a language prompt based on the provided language code. Then, it calls the `StartNewChat` function to initiate a new chat session with the AI model. The `StartNewChat` function takes the context, GPT key, model, AI endpoint, and language prompt as parameters. It returns a response, a chat session thread, and an error. If there is an error, the function returns an empty string, nil, and the error. Otherwise, it returns the response and the chat session thread.  
  
#### Other Functions:  
  
The code also includes other functions, such as `errorMessage`, `ContinueChatWithContextNoLimit`, and `GenerateContentLAI`, which are not described in detail in the provided code. These functions likely handle error messages, continue existing chat sessions, and generate content using the AI model, respectively.  
  
  
  
# lib/langchain/startDialogSequence.go  
## Package: langchain  
  
### Imports:  
  
```  
context  
io/ioutil  
log  
math/rand  
os  
path/filepath  
github.com/JackBekket/hellper/lib/database  
github.com/go-telegram-bot-api/telegram-bot-api/v5  
```  
  
### External Data, Input Sources:  
  
1. Database: The code interacts with a database (likely a key-value store) to store user data and session information. The database is accessed through the `db` package.  
2. Telegram Bot API: The code uses the `tgbotapi` package to interact with the Telegram Bot API. This allows the code to send messages and receive updates from users via Telegram.  
  
### Code Summary:  
  
#### errorMessage Function:  
  
This function is called when an error occurs during the creation of a request. It logs the error, sends an error message to the user, and removes the user from the database. Additionally, it sends a random video from the "media" directory as a helper video.  
  
#### StartDialogSequence Function:  
  
This function initiates a dialog sequence with an AI model. It takes the following arguments:  
  
1. `bot`: A pointer to a Telegram bot API instance.  
2. `chatID`: The ID of the chat where the dialog will take place.  
3. `promt`: The prompt to be sent to the AI model.  
4. `ctx`: A context object for managing the request.  
5. `ai_endpoint`: The endpoint for the AI model.  
  
The function first retrieves the user's session information from the database. Then, it continues the chat using the provided prompt and the user's existing dialog thread. The AI response is sent to the user via Telegram, and the user's dialog status and session usage are updated in the database.  
  
#### LogResponse Function:  
  
This function is not used in the provided code but is commented out. It would have logged various details about the AI response, such as the model, object, and usage information.  
  
#### Other Code:  
  
The code also includes a mutex (`mu`) to protect shared resources and a variable `userDb` to store user data.  
  
  
  
# lib/localai/startDialogSequence.go  
## Package: localai  
  
### Imports:  
- context  
- log  
- db "github.com/JackBekket/hellper/lib/database"  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data, Input Sources:  
- db.UsersMap: A map containing user data from the database.  
- chatID: The ID of the chat to which the message should be sent.  
- promt: The prompt to be sent to the AI model.  
- ctx: A context object for managing the request.  
- ai_endpoint: The endpoint for the AI model.  
  
### Code Summary:  
  
#### errorMessage Function:  
This function handles errors that occur during the process of creating a request. It logs the error, sends an error message to the user, and removes the user from the database.  
  
#### StartDialogSequence Function:  
This function initiates a dialog sequence with the AI model. It retrieves the user's AI session data, logs the GPT model and prompt, and calls the GenerateCompletion function to get the AI's response. If an error occurs, it calls the errorMessage function. Otherwise, it logs the response, formats the response text, and sends it to the user. Finally, it updates the user's dialog status and saves the changes to the database.  
  
#### LogResponse Function:  
This function logs the full response object, including the created timestamp, response ID, model, object, choices, and usage information.  
  
  
  
