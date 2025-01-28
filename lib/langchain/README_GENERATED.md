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

lib/langchain/langchain.go
## langchain

### Imports
```
import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"

	//"github.com/tmc/langchaingo/llms/options"
	"github.com/tmc/langchaingo/llms/openai"
)
```

### External Data, Input Sources
- API token for OpenAI API
- Model name for OpenAI API
- Base URL for OpenAI API (optional)
- User initial prompt

### Major Code Parts
#### InitializeNewChatWithContextNoLimit
This function initializes a new chat session with the specified API token, model name, base URL, and user initial prompt. It creates a new conversation using the OpenAI LLM and a memory buffer to store the conversation history.

#### StartNewChat
This function starts a new conversation by calling InitializeNewChatWithContextNoLimit and then running the user's initial prompt through the conversation chain. It returns the result of the conversation and the updated chat session.

#### RunChain
This function runs a prompt through the provided chat session and returns the result and the updated chat session.

#### ContinueChatWithContextNoLimit
This function continues a conversation by running a new prompt through the provided chat session and returning the result and the updated chat session.

#### GenerateContentInstruction
This function generates content from a single prompt without using memory or context. It takes the base URL, prompt, model name, API token, network, and options as input and returns the generated content and any errors encountered.



lib/langchain/langgraph.go
## Package: langchain

### Imports:
- github.com/JackBekket/hellper/lib/agent
- github.com/JackBekket/hellper/lib/database
- github.com/tmc/langchaingo/llms/openai

### External Data, Input Sources:
- api_token (string)
- model_name (string)
- base_url (string)
- user_promt (string)
- state (db.ChatSessionGraph)

### RunNewAgent Function:
This function initializes a new agent with the provided API token, model name, and base URL. It creates a new OpenAI LLM instance using the provided parameters and runs a thread using the agent.RunThread function. The function returns a new ChatSessionGraph object containing the conversation buffer and the output text.

### ContinueAgent Function:
This function continues an existing agent with the provided API token, model name, base URL, and user prompt. It creates a new OpenAI LLM instance using the provided parameters and runs a thread using the agent.RunThread function, passing the existing conversation buffer as input. The function returns a new ChatSessionGraph object containing the updated conversation buffer and the output text.



lib/langchain/setupSequenceWithKey.go
## Package: langchain

### Imports:

- context
- log
- sync
- db "github.com/JackBekket/hellper/lib/database"
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:

- Database: db.User, db.ChatSessionGraph, db.GetSessionUsage
- Telegram Bot API: tgbotapi.BotAPI, tgbotapi.NewMessage, bot.Send
- AI Endpoint: ai_endpoint

### Code Summary:

#### SetupSequenceWithKey Function:

This function is responsible for setting up a sequence of interactions with an AI model based on user preferences and session data. It takes a Telegram bot instance, a user object, a language, a context, and an AI endpoint as input.

1. It first acquires a lock to ensure thread safety.
2. It retrieves the user's ID, GPT key, and model from the user object.
3. Based on the provided language, it calls the tryLanguage function to initiate a conversation with the AI model.
4. If successful, it updates the user's dialog status, AI session data, and usage information.
5. Finally, it stores the updated user object in the db.UsersMap.

#### tryLanguage Function:

This function is responsible for initiating a conversation with the AI model based on the provided language and language code. It takes a user object, a language string, a language code, and an AI endpoint as input.

1. It constructs a language prompt based on the language code.
2. It retrieves the user's GPT key, model, and chat ID from the user object.
3. It calls the RunNewAgent function to start a new agent and initiate the conversation with the AI model.
4. If successful, it returns the conversation thread and the AI model's response.

#### RunNewAgent Function:

This function is responsible for starting a new agent and initiating a conversation with the AI model. It takes the GPT key, model, AI endpoint, and language prompt as input.

1. It uses the provided GPT key, model, and AI endpoint to connect to the AI model.
2. It sends the language prompt to the AI model and waits for a response.
3. It returns the conversation thread and the AI model's response.



lib/langchain/startDialogSequence.go
## Package: langchain

### Imports:

```
context
log
math/rand
os
path/filepath
io/fs
github.com/JackBekket/hellper/lib/database
github.com/go-telegram-bot-api/telegram-bot-api/v5
```

### External Data, Input Sources:

- Database: `db.UsersMap`
- Telegram Bot API: `tgbotapi.BotAPI`
- Media directory: `../../media/`

### Code Summary:

#### errorMessage Function:

This function handles errors that occur during the process of creating a request. It logs the error, sends an error message to the user, and then sends a helper video to the user. The helper video is selected randomly from the media directory.

#### StartDialogSequence Function:

This function initiates a dialog sequence with an AI model. It takes the following parameters:

- `bot`: A Telegram bot API instance.
- `chatID`: The ID of the chat to send the message to.
- `promt`: The prompt to send to the AI model.
- `ctx`: A context object.
- `ai_endpoint`: The endpoint for the AI model.

The function first retrieves the user's AI session data from the database. Then, it uses the provided parameters to continue the agent's dialog thread. If an error occurs, the `errorMessage` function is called. Otherwise, the AI response is sent to the user, and the user's dialog status and usage are updated in the database.

#### LogResponse Function:

This function is commented out but appears to be intended for logging the full response object from an AI model. It would log various attributes of the response, such as the model, object, choices, and usage information.



