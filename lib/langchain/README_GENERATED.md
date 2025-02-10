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
- `context.Context`
- `schema.AgentAction`
- `schema.AgentFinish`
- `schema.Document`
- `llms.MessageContent`
- `llms.ContentResponse`
- `llms.ContentChoice`
- `db.User`
- `db.UpdateSessionUsage`

### TODOs
- Implement all the methods in the `ChainCallbackHandler` struct.

### Summary
The code defines a `ChainCallbackHandler` struct that implements various callback methods for handling different events in a chain of operations. These events include agent actions, agent finishes, chain ends, chain errors, chain starts, LLM errors, LLM generate content starts, LLM starts, retriever ends, retriever starts, streaming functions, tool ends, tool errors, and tool starts.

The `HandleLLMGenerateContentEnd` method is responsible for logging the content choice, stop reason, context, and generation information. It also updates the user's usage information based on the prompt tokens, completion tokens, and total tokens.

The code also includes a `LogResponseContentChoice` function that logs the content choice, stop reason, context, and generation information. It also updates the user's usage information based on the prompt tokens, completion tokens, and total tokens.

The code is designed to be used as part of a larger system that handles chains of operations involving LLMs, agents, and other components. The callback methods provide a way to handle events and update the system state accordingly.

lib/langchain/langchain.go
## Package: langchain

### Imports:
- context
- fmt
- log
- github.com/tmc/langchaingo/llms
- github.com/tmc/langchaingo/llms/openai

### External Data, Input Sources:
- base_url (string)
- promt (string)
- model_name (string)
- api_token (string)
- network (string)
- options (llms.CallOption)

### TODOs:
- None

### Summary:
The code defines a function called `GenerateContentInstruction` that takes several parameters, including a base URL, prompt, model name, API token, and network. It uses the OpenAI API to generate content from a single prompt. The function first checks the network type and initializes the OpenAI client accordingly. Then, it calls the `GenerateFromSinglePrompt` function from the `llms` package to generate the content. Finally, it returns the generated content as a string.

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

### TODOs:
- TODO: Implement ChainCallbackHandler

### RunNewAgent Function:
This function initializes a new agent with the given API token, model name, base URL, and user prompt. It creates a new OpenAI model using the provided parameters and runs a thread using the agent.RunThread function. The function returns a new ChatSessionGraph object containing the conversation buffer and the output text.

### ContinueAgent Function:
This function continues an existing agent with the given API token, model name, base URL, user prompt, and state. It creates a new OpenAI model using the provided parameters and runs a thread using the agent.RunThread function with the existing state. The function returns a new ChatSessionGraph object containing the conversation buffer and the output text.



lib/langchain/setupSequenceWithKey.go
## Package: langchain

### Imports:

- context
- log
- sync
- db "github.com/JackBekket/hellper/lib/database"
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:

- The code uses a database (presumably a local database) to store user information, including their AI session data.
- The code also uses a Telegram bot API to interact with users.

### TODOs:

- None found in the provided code.

### Summary:

The code defines a function called `SetupSequenceWithKey` that sets up a sequence for a user based on their language preference and other session data. The function takes a Telegram bot API instance, a user object, a language string, a context, and an AI endpoint as input.

The function first locks a mutex to ensure thread safety. Then, it retrieves the user's GPT key, model, and other session data from the user object. Based on the language preference, the function calls the `tryLanguage` function to determine the appropriate language prompt and model to use.

The `tryLanguage` function takes the user object, language string, language code, and AI endpoint as input. It constructs a language prompt based on the language code and calls the `RunNewAgent` function to execute the prompt using the user's GPT key and model. The function returns the agent's response and the chat session graph.

The `SetupSequenceWithKey` function then sends the agent's response to the user via the Telegram bot API and updates the user's session data, including their dialog status and usage. Finally, it stores the updated user object in a map.

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

### TODOs:

- Investigate why meme videos with helper are not sent by the `errorMessage` function.

### Code Summary:

#### `errorMessage` function:

This function is called when an error occurs during the process of creating a request. It first logs the error and then sends a message to the user informing them of the error and instructing them to recreate the client and initialize a new session.

The function then attempts to send a random video from the media directory as a helper video. It first reads the contents of the media directory, selects a random file, opens the video file, and creates a new video message. Finally, it sends the video message to the user.

#### `StartDialogSequence` function:

This function is responsible for starting a dialog sequence with an AI model. It first retrieves the user's AI session information from the database and then calls the `ContinueAgent` function to continue the conversation.

The `ContinueAgent` function takes the API key, GPT model, base URL, prompt, and dialog thread as input and returns the updated dialog thread, response, and any errors that occurred. If an error occurs, the `errorMessage` function is called. Otherwise, the response is sent to the user, and the dialog thread is updated.

The function also logs the total number of turns in the conversation and iterates over each message in the history, printing the content of each message. Finally, the updated dialog thread is stored in the database.

