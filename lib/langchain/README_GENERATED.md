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



lib/langchain/handler.go
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



