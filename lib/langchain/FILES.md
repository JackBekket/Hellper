# lib/langchain/handler.go  
## langchain  
  
This package provides a ChainCallbackHandler struct and various methods to handle different events during the execution of a chain.  
  
### Imports  
  
```  
import (  
	"context"  
	"encoding/json"  
	"log"  
  
	db "github.com/JackBekket/hellper/lib/database"  
	"github.com/tmc/langchaingo/llms"  
	"github.com/tmc/langchaingo/schema"  
	//""  
)  
```  
  
### External Data, Input Sources  
  
The package uses the following external data and input sources:  
  
1. `db.User`: A struct representing a user, likely from a database.  
2. `llms.ContentResponse`: A struct containing information about the generated content, including choices and generation info.  
3. `schema.AgentAction`, `schema.AgentFinish`, `schema.Document`, etc.: Structs representing various events and data structures related to the chain execution.  
  
### Code Summary  
  
1. `ChainCallbackHandler` struct: This struct is responsible for handling various events during the chain execution. It has methods for handling agent actions, agent finishes, chain ends, chain errors, chain starts, LLM errors, LLM generate content starts, LLM starts, retriever ends, retriever starts, streaming functions, tool ends, tool errors, and tool starts.  
  
2. `HandleLLMGenerateContentEnd`: This method is called when the LLM has finished generating content. It logs the content, stop reason, context, and generation info. It also updates the user's usage information based on the generated content and saves it to the database.  
  
3. `LogResponseContentChoice`: This helper function logs the content, stop reason, context, and generation info of the chosen content. It also logs the prompt tokens, completion tokens, and total tokens from the generation info.  
  
4. `HandleText`: This method is intended to be implemented if needed to handle text input.  
  
5. Other methods: The remaining methods in the `ChainCallbackHandler` struct are placeholders for handling various events during the chain execution. They are currently unimplemented but can be filled in as needed.  
  
6. Database interaction: The package interacts with a database to store user usage information. It uses the `db.UpdateSessionUsage` function to update the user's usage based on the generated content.  
  
7. Logging: The package uses the `log` package to log various events and information during the chain execution. This includes logging the content, stop reason, context, generation info, and other relevant data.  
  
8. Error handling: The package includes error handling mechanisms, such as panic statements and error checking, to ensure that the code can handle unexpected situations gracefully.  
  
9. Type assertions: The package uses type assertions to ensure that the data being accessed is of the expected type. This helps to prevent runtime errors and ensure that the code is working as intended.  
  
# lib/langchain/langchain.go  
## langchain_controller  
  
This package provides functions for interacting with language models, specifically OpenAI's API and a local AI model. It also includes a database component for managing chat sessions.  
  
### Imports  
  
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
  
### External Data and Input Sources  
  
The package uses the following external data and input sources:  
  
1. OpenAI API: Used for interacting with OpenAI's language models.  
2. Local AI model: A local AI model that can be used as an alternative to OpenAI's API.  
3. Database: Used for storing and retrieving chat session data.  
  
### Code Summary  
  
1. **InitializeNewChatWithContextNoLimit()**: This function initializes a new chat session with a given API token, model name, base URL, and user initial prompt. It creates a new conversation using the specified language model and memory buffer.  
  
2. **StartNewChat()**: This function starts a new conversation by calling InitializeNewChatWithContextNoLimit() and then running the conversation using the RunChain() function.  
  
3. **RunChain()**: This function runs a given prompt through the specified conversation, generating a response.  
  
4. **ContinueChatWithContextNoLimit()**: This function continues an existing chat session by running a given prompt through the conversation, generating a response.  
  
5. **GenerateContentInstruction()**: This function generates content from a single prompt without using memory or context. It takes a base URL, prompt, model name, API token, network, and options as input.  
  
6. **ChainCallbackHandler**: This struct is used as a callback handler for the language model.  
  
7. **db.ChatSession**: This struct represents a chat session, containing the conversation buffer and dialog thread.  
  
8. **memory.NewConversationBuffer()**: This function creates a new conversation buffer for storing chat history.  
  
9. **chains.NewConversation()**: This function creates a new conversation using the specified language model and memory buffer.  
  
10. **llms.GenerateFromSinglePrompt()**: This function generates a response from a single prompt using the specified language model and options.  
  
11. **openai.New()**: This function creates a new OpenAI client using the provided API token, model name, and base URL.  
  
12. **openai.WithToken()**: This function sets the API token for the OpenAI client.  
  
13. **openai.WithModel()**: This function sets the model name for the OpenAI client.  
  
14. **openai.WithBaseURL()**: This function sets the base URL for the OpenAI client.  
  
15. **openai.WithAPIVersion()**: This function sets the API version for the OpenAI client.  
  
16. **openai.WithCallback()**: This function sets the callback handler for the OpenAI client.  
  
  
  
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
  
1. Database: The code uses a database to store user information, including their AI session data. The database is accessed through the `db` package.  
2. Telegram Bot API: The code uses the Telegram Bot API to interact with a Telegram bot. The API is accessed through the `tgbotapi` package.  
  
### Code Summary:  
  
#### SetupSequenceWithKey Function:  
  
This function is responsible for setting up a sequence of interactions with an AI model, based on the user's language preference and other session data. It takes the following parameters:  
  
1. `bot`: A pointer to a Telegram bot instance.  
2. `user`: A `db.User` struct containing user information, including their AI session data.  
3. `language`: A string representing the user's preferred language.  
4. `ctx`: A context object for managing the execution of the function.  
5. `ai_endpoint`: A string representing the endpoint for the AI model.  
  
The function first locks a mutex to ensure thread safety. Then, it retrieves the user's GPT key, model, and other session data from the `user` struct. Based on the `language` parameter, it calls the `tryLanguage` function to initiate a conversation with the AI model. The `tryLanguage` function returns a response, a chat session object, and an error. If there is an error, the function calls the `errorMessage` function to handle the error. Otherwise, it updates the user's dialog status, AI session data, and stores the updated user information in the `db.UsersMap`.  
  
#### tryLanguage Function:  
  
This function is responsible for initiating a conversation with the AI model based on the user's language preference. It takes the following parameters:  
  
1. `user`: A `db.User` struct containing user information, including their AI session data.  
2. `language`: A string representing the user's preferred language.  
3. `languageCode`: An integer representing the language code (0 - default, 1 - Russian, 2 - English).  
4. `ctx`: A context object for managing the execution of the function.  
5. `ai_endpoint`: A string representing the endpoint for the AI model.  
  
The function first constructs a language prompt based on the `languageCode` parameter. Then, it calls the `StartNewChat` function to initiate a new chat session with the AI model. The `StartNewChat` function returns a response, a chat session object, and an error. If there is an error, the function returns an empty string, nil, and the error. Otherwise, it returns the response, chat session object, and nil.  
  
#### StartNewChat Function:  
  
This function is responsible for starting a new chat session with the AI model. It takes the following parameters:  
  
1. `ctx`: A context object for managing the execution of the function.  
2. `gptKey`: A string representing the user's GPT key.  
3. `model`: A string representing the AI model to use.  
4. `ai_endpoint`: A string representing the endpoint for the AI model.  
5. `languagePromt`: A string representing the language prompt to use.  
  
The function returns a response, a chat session object, and an error. If there is an error, the function returns an empty string, nil, and the error. Otherwise, it returns the response, chat session object, and nil.  
  
#### ContinueChatWithContextNoLimit Function:  
  
This function is responsible for continuing a chat session with the AI model. It takes the following parameters:  
  
1. `thread`: A chat session object representing the current chat thread.  
2. `languagePromt`: A string representing the language prompt to use.  
  
The function returns a response and an error. If there is an error, the function returns an empty string and the error. Otherwise, it returns the response and nil.  
  
#### GenerateContentLAI Function:  
  
This function is responsible for generating content using the AI model. It takes the following parameters:  
  
1. `ai_endpoint`: A string representing the endpoint for the AI model.  
2. `model`: A string representing the AI model to use.  
3. `languagePromt`: A string representing the language prompt to use.  
  
The function returns a response and an error. If there is an error, the function returns an empty string and the error. Otherwise, it returns the response and nil.  
  
#### LogResponseContentChoice Function:  
  
This function is responsible for logging the content of the AI model's response. It takes the following parameter:  
  
1. `resp`: A response object containing the AI model's response.  
  
The function logs the content of the response's first choice.  
  
#### errorMessage Function:  
  
This function is responsible for handling errors during the interaction with the AI model. It takes the following parameters:  
  
1. `err`: An error object representing the error that occurred.  
2. `bot`: A pointer to a Telegram bot instance.  
3. `user`: A `db.User` struct containing user information.  
  
The function sends an error message to the user through the Telegram bot.  
  
  
  
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
  
- Database: `db.UsersMap` (likely a map of user IDs to user objects)  
- Telegram Bot API: `tgbotapi.BotAPI` (for sending messages and handling interactions with the Telegram bot)  
- Media directory: `../../media/` (for retrieving random video files to send as error messages)  
  
### Code Summary:  
  
#### `errorMessage` function:  
  
- This function is called when an error occurs during the process of creating a request.  
- It logs the error, sends an error message to the user, and then sends a helper video as an error message.  
- It also removes the user from the database (temporary solution).  
  
#### `StartDialogSequence` function:  
  
- This function initiates a dialog sequence with an AI model.  
- It takes the Telegram bot API, chat ID, prompt, context, and AI endpoint as input.  
- It retrieves the user's AI session and dialog thread from the database.  
- It calls the `ContinueChatWithContextNoLimit` function to continue the conversation with the AI model.  
- If there is an error, it calls the `errorMessage` function.  
- If the AI response is successful, it sends the response to the user, updates the user's dialog status, and stores the updated dialog thread in the database.  
  
#### Other code:  
  
- There is a commented-out function `LogResponse` that seems to be intended for logging the full response object from the AI model.  
  
  
  
