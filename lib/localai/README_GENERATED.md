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



lib/localai/localai.go
## localai package

This package provides functions for interacting with a local AI model and generating images using Stable Diffusion.

### Imports

```
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/StarkBotsIndustries/telegraph"
)
```

### External Data, Input Sources

The package uses the following external data and input sources:

1. Local AI model: The package assumes that a local AI model is available at the specified URL.
2. Stable Diffusion model: The package uses the Stable Diffusion model for image generation.
3. Telegraph API: The package uses the Telegraph API for uploading images to a Telegram channel.

### Code Summary

1. Chat Completion: The `GenerateCompletion` function sends a chat request to the local AI model and returns the response. It takes the prompt, model name, and URL as input. The function first creates a JSON payload containing the prompt, model name, and temperature. Then, it sends a POST request to the specified URL with the JSON payload. The response is parsed as a JSON object, and the relevant information is extracted.

2. Image Generation: The `GenerateImageStableDissusion` function generates an image using the Stable Diffusion model. It takes the prompt and size as input. The function creates a JSON payload containing the prompt and size, and sends a POST request to the specified URL. The response is parsed as a JSON object, and the image URL is extracted.

3. Upload to Telegraph: The `UploadToTelegraph` function uploads the generated image to a Telegram channel using the Telegraph API. It takes the image file path as input. The function opens the file, uploads it to the Telegraph API, and returns the link to the uploaded image.

4. Delete from Temp: The `deleteFromTemp` function deletes the generated image from the temporary directory. It takes the image file name as input. The function constructs the absolute path to the image file and deletes it using the `os.Remove` function.

5. Wrong Password Error: The `WrongPwdError` struct represents an error that occurs when the provided password is incorrect. It has a single field, `message`, which contains the error message.

6. Main Function: The `main` function demonstrates how to use the package to generate a chat response and an image. It sets the prompt, model name, and URL, and calls the `GenerateCompletion` and `GenerateImageStableDissusion` functions to perform the respective tasks.

7. Generate Completion with Password: The `GenerateCompletionWithPWD` function is similar to `GenerateCompletion` but also takes a secret password as input. It checks if the provided password matches the secret password and returns an error if they don't match. Otherwise, it calls `GenerateCompletion` to generate the chat response.

8. Upload to Telegraph: The `UploadToTelegraph` function is responsible for uploading the generated image to a Telegram channel using the Telegraph API. It takes the image file path as input and returns the link to the uploaded image.

9. Delete from Temp: The `deleteFromTemp` function is responsible for deleting the generated image from the temporary directory. It takes the image file name as input and deletes the file using the `os.Remove` function.



lib/localai/setupSequenceWithKey.go
## Package: localai

### Imports:
- context
- log
- sync

- db "github.com/JackBekket/hellper/lib/database"
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:
- Database: `db.User` struct, which contains information about the user, including their AI session data.
- Telegram Bot API: `tgbotapi.BotAPI` struct, which is used to interact with the Telegram bot.
- AI Endpoint: `ai_endpoint` string, which specifies the URL of the AI service to be used.
- Password: `spwd` string, which is used for authentication with the AI service.
- User's GPT Key: `gptKey` string, which is stored in the user's AI session data.

### Code Summary:

#### SetupSequenceWithKey Function:

This function is responsible for setting up the sequence with the user's GPT key and language. It takes the following parameters:

- `bot`: A pointer to the Telegram bot API instance.
- `user`: A `db.User` struct containing the user's information.
- `language`: A string representing the user's preferred language.
- `ctx`: A context object for managing the request.
- `spwd`: A string representing the password for authentication with the AI service.
- `ai_endpoint`: A string representing the URL of the AI service.

The function first locks a mutex to ensure thread safety. Then, it retrieves the user's GPT key and model from their AI session data. It then calls the `tryLanguage` function to determine the appropriate language prompt based on the user's preferred language and language code. Finally, it sends the generated response to the user via the Telegram bot and updates the user's dialog status.

#### tryLanguage Function:

This function takes the user's information, language code, context, AI endpoint, password, and user's GPT key as input. It constructs a language prompt based on the language code and calls the `GenerateCompletionWithPWD` function to generate a response from the AI service. The response is then logged and returned as a string.

#### GenerateCompletionWithPWD Function:

This function is responsible for generating a completion from the AI service using the provided prompt, model, AI endpoint, password, and user's GPT key. It returns the generated response and any errors encountered during the process.

#### LogResponse Function:

This function logs the AI service's response to the console.

#### errorMessage Function:

This function handles any errors that occur during the process and sends an error message to the user via the Telegram bot.



