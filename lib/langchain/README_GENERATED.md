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



lib/langchain/setupSequenceWithKey.go
## Package: langchain

### Imports:
- context
- log
- sync
- db (from github.com/JackBekket/hellper/lib/database)
- tgbotapi (from github.com/go-telegram-bot-api/telegram-bot-api/v5)

### External Data, Input Sources:
- Database: db.User, db.ChatSession, db.GetSessionUsage
- Telegram Bot API: tgbotapi.BotAPI, tgbotapi.NewMessage, bot.Send
- AI Endpoint: ai_endpoint

### Code Summary:
#### SetupSequenceWithKey Function:
- Takes a Telegram bot, user, language, context, and AI endpoint as input.
- Retrieves user's GPT key, model, and network from the session.
- Determines the language based on the provided language parameter.
- Calls the tryLanguage function to handle language-specific logic.
- Updates the user's dialog status, AI session, and usage.
- Stores the updated user in the db.UsersMap.

#### tryLanguage Function:
- Takes a user, language, language code, context, and AI endpoint as input.
- Constructs a language prompt based on the language code.
- Initializes a new chat session using the StartNewChat function.
- Returns the chat result, dialog thread, and any errors encountered.

#### StartNewChat Function:
- Takes context, GPT key, model, AI endpoint, and language prompt as input.
- Starts a new chat session using the provided parameters.
- Returns the chat result and dialog thread.

#### ContinueChatWithContextNoLimit Function:
- Takes a chat thread and language prompt as input.
- Continues the chat session using the provided parameters.
- Returns the chat result and any errors encountered.

#### GenerateContentLAI Function:
- Takes an AI endpoint, model, and language prompt as input.
- Generates content using the provided parameters.
- Returns the generated content and any errors encountered.

#### LogResponseContentChoice Function:
- Takes a response as input.
- Logs the content of the first choice in the response.

#### errorMessage Function:
- Takes an error, bot, and user as input.
- Sends an error message to the user via the Telegram bot.

lib/langchain/startDialogSequence.go
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

This function is called when an error occurs during the process of creating a request. It logs the error, sends an error message to the user, and then sends a helper video to the user. The helper video is selected randomly from a directory containing media files.

#### StartDialogSequence Function:

This function initiates a dialog sequence with an AI model. It takes the following arguments:

1. `bot`: A pointer to a Telegram bot API instance.
2. `chatID`: The ID of the chat where the dialog will take place.
3. `promt`: The prompt to be sent to the AI model.
4. `ctx`: A context object for managing the execution of the function.
5. `ai_endpoint`: The endpoint for the AI model.

The function first retrieves the user's session information from the database. Then, it uses the `ContinueChatWithContextNoLimit` function to continue the conversation with the AI model. If an error occurs, the `errorMessage` function is called. Otherwise, the AI response is sent to the user via Telegram, and the user's session information is updated in the database.

#### LogResponse Function:

This function is not used in the provided code but is commented out. It would have been used to log the full response object from the AI model, including information about the model, the response ID, and the usage statistics.



lib/langchain/handler.go
## Package: langchain

### Imports:
- context
- encoding/json
- log
- github.com/JackBekket/hellper/lib/database
- github.com/tmc/langchaingo/llms
- github.com/tmc/langchaingo/schema

### External Data, Input Sources:
- Database: The code uses a database (likely a relational database like PostgreSQL or MySQL) to store user information and session usage data. The database is accessed through the `db` package.

### Code Summary:
#### ChainCallbackHandler:
- This struct implements a callback handler for various events in a LangChain agent.
- It includes methods for handling agent actions, agent finishes, chain ends, chain errors, chain starts, LLM errors, LLM content generation starts, LLM starts, retriever ends, retriever starts, streaming functions, tool ends, tool errors, and tool starts.
- The `HandleLLMGenerateContentEnd` method is responsible for logging the content choice, stop reason, context, and generation information.
- The `LogResponseContentChoice` method is called within `HandleLLMGenerateContentEnd` to log the content choice, stop reason, context, and generation information.
- It also updates the user's session usage information based on the number of prompt tokens, completion tokens, and total tokens.

#### Other Methods:
- The code includes other methods like `HandleText` and `HandleLLMGenerateContentEnd` that can be implemented as needed.

#### Database Interaction:
- The code uses the `db` package to interact with the database.
- It retrieves user information from the context and updates the user's session usage information based on the number of prompt tokens, completion tokens, and total tokens.

#### Logging:
- The code uses the `log` package to log various events and information, such as content choices, stop reasons, generation information, and user information.

#### Usage Information:
- The code updates the user's session usage information based on the number of prompt tokens, completion tokens, and total tokens.
- This information is stored in a separate structure to avoid race conditions.

#### Summary:
- The code provides a callback handler for various events in a LangChain agent.
- It includes methods for handling agent actions, agent finishes, chain ends, chain errors, chain starts, LLM errors, LLM content generation starts, LLM starts, retriever ends, retriever starts, streaming functions, tool ends, tool errors, and tool starts.
- The code also interacts with a database to store user information and session usage data.
- It uses the `log` package to log various events and information.

