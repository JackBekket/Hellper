# lib/langchain/handler.go  
## langchain  
  
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
- `context.Context`: Provides context for the code execution.  
- `encoding/json`: Used for marshaling and unmarshaling JSON data.  
- `log`: Provides logging functionality.  
- `db "github.com/JackBekket/hellper/lib/database"`: Database library for interacting with the database.  
- `llms`: Library for interacting with large language models.  
- `schema`: Library for defining and working with schemas.  
  
### Code Summary  
The provided code defines a `ChainCallbackHandler` struct, which implements various callback functions for handling different events during the execution of a chain. These callbacks are used to handle actions, finishes, errors, and other events related to the chain.  
  
The `HandleLLMGenerateContentEnd` function is responsible for handling the end of an LLM content generation process. It logs the content, stop reason, context, and generation information. It also updates the user's usage information based on the generated content and saves it to the database.  
  
The `LogResponseContentChoice` function is called within `HandleLLMGenerateContentEnd` to log the content, stop reason, context, and generation information. It also logs specific fields from the generation information, such as prompt tokens, completion tokens, and total tokens.  
  
The code also includes a comment indicating that the user's usage information should be updated based on the generated content and saved to the database. This suggests that the code is part of a larger system that tracks and manages user usage.  
  
The code snippet demonstrates the use of callbacks to handle events during the execution of a chain, as well as the logging of relevant information for debugging and monitoring purposes. It also highlights the importance of updating user usage information and saving it to the database.  
  
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
  
3. **RunChain()**: This function runs a given prompt through the provided conversation, generating a response.  
  
4. **ContinueChatWithContextNoLimit()**: This function continues an existing chat session by running a given prompt through the conversation, generating a response.  
  
5. **GenerateContentInstruction()**: This function generates content from a single prompt without using memory or context. It takes a base URL, prompt, model name, API token, network, and options as input.  
  
6. **ChainCallbackHandler**: This struct is used as a callback handler for the language model.  
  
7. **db.ChatSession**: This struct represents a chat session, containing the conversation buffer and dialog thread.  
  
8. **memory.NewConversationBuffer()**: This function creates a new conversation buffer for storing chat history.  
  
9. **chains.NewConversation()**: This function creates a new conversation using the provided language model and memory buffer.  
  
10. **llms.GenerateFromSinglePrompt()**: This function generates a response from a single prompt using the provided language model and options.  
  
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
  
- `db.User`: User data from the database.  
- `ai_endpoint`: Endpoint for the AI model.  
- `language`: Language preference for the user.  
- `ctx`: Context for the request.  
  
### Code Summary:  
  
#### SetupSequenceWithKey Function:  
  
This function is responsible for setting up a sequence of interactions with the AI model based on the user's preferences and session data. It takes the following parameters:  
  
- `bot`: Telegram bot API instance.  
- `user`: User data from the database.  
- `language`: Language preference for the user.  
- `ctx`: Context for the request.  
- `ai_endpoint`: Endpoint for the AI model.  
  
The function first retrieves the user's GPT key and model from the session data. Then, it determines the appropriate language based on the user's preference and calls the `tryLanguage` function to initiate the conversation with the AI model.  
  
#### tryLanguage Function:  
  
This function is responsible for initiating a conversation with the AI model based on the user's language preference. It takes the following parameters:  
  
- `user`: User data from the database.  
- `language`: Language preference for the user.  
- `languageCode`: Code for the language (0 - default, 1 - Russian, 2 - English).  
- `ctx`: Context for the request.  
- `ai_endpoint`: Endpoint for the AI model.  
  
The function first constructs a language prompt based on the language code and then calls the `StartNewChat` function to initiate a new conversation with the AI model. The result of the conversation is returned along with the dialog thread.  
  
#### StartNewChat Function:  
  
This function is responsible for starting a new conversation with the AI model. It takes the following parameters:  
  
- `ctx`: Context for the request.  
- `gptKey`: GPT key for the user.  
- `model`: AI model to use.  
- `ai_endpoint`: Endpoint for the AI model.  
- `languagePromt`: Language prompt to use for the conversation.  
  
The function returns the result of the conversation and the dialog thread.  
  
#### ContinueChatWithContextNoLimit Function:  
  
This function is responsible for continuing a conversation with the AI model. It takes the following parameters:  
  
- `thread`: Dialog thread for the conversation.  
- `languagePromt`: Language prompt to use for the conversation.  
  
The function returns the result of the conversation.  
  
#### GenerateContentLAI Function:  
  
This function is responsible for generating content using the AI model. It takes the following parameters:  
  
- `ai_endpoint`: Endpoint for the AI model.  
- `model`: AI model to use.  
- `languagePromt`: Language prompt to use for the conversation.  
  
The function returns the result of the content generation.  
  
#### LogResponseContentChoice Function:  
  
This function is responsible for logging the content choice from the AI model's response. It takes the following parameter:  
  
- `resp`: Response from the AI model.  
  
The function logs the content choice from the response.  
  
  
  
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
  
### Code Summary:  
  
#### errorMessage Function:  
  
- This function is called when an error occurs during the process of creating a request.  
- It logs the error, sends an error message to the user, and then sends a helper video to the user.  
- The helper video is selected randomly from a list of files in the "media" directory.  
- Finally, it removes the user from the database (temporary solution).  
  
#### StartDialogSequence Function:  
  
- This function initiates a dialog sequence with an AI model.  
- It takes the Telegram bot API, chat ID, prompt, context, and AI endpoint as input.  
- It retrieves the user object from the database and initializes the AI session.  
- It calls the `ContinueChatWithContextNoLimit` function to continue the conversation with the AI model.  
- If there is an error, it calls the `errorMessage` function.  
- If the AI response is successful, it sends the response to the user, updates the user's dialog status, and stores the updated user object back in the database.  
  
#### LogResponse Function:  
  
- This function is commented out but appears to be intended for logging the full response object from the AI model.  
- It would log various attributes of the response, such as the creation time, ID, model, object, choices, and usage information.  
  
  
  
