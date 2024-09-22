# Package: localai

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



lib/localai/setupSequenceWithKey.go
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



lib/localai/startDialogSequence.go
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



