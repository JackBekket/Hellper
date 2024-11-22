# lib/localai/localai.go  
## localai package  
  
This package provides functions for interacting with various AI models, including chat completion, image generation, and transcription.  
  
### Imports  
  
The package imports the following packages:  
  
- bytes  
- encoding/json  
- fmt  
- io  
- io/ioutil  
- log  
- mime/multipart  
- net/http  
- os  
- path/filepath  
- strings  
  
### External Data, Input Sources  
  
The package uses the following external data and input sources:  
  
- API endpoints for various AI models (e.g., chat completion, image generation, transcription)  
- Environment variables (e.g., OPENAI_API_KEY)  
  
### Code Summary  
  
#### Chat Completion  
  
The `GenerateCompletion` function sends a chat completion request to a specified API endpoint. It takes the prompt, model name, and API URL as input. The function first creates a JSON payload containing the prompt, model name, and temperature. Then, it sends a POST request to the API endpoint with the JSON payload. The response is parsed as a ChatResponse object, which contains the assistant's response.  
  
#### Image Generation  
  
The `GenerateImageStableDiffusion` function generates an image using the Stable Diffusion model. It takes the prompt, image size, API URL, and model name as input. The function creates a JSON payload containing the prompt, model name, and image size. It then sends a POST request to the API endpoint with the JSON payload. The response is parsed as a GenerationResponse object, which contains the URL of the generated image.  
  
#### Transcription  
  
The `TranscribeWhisper` function transcribes an audio file using the Whisper model. It takes the API URL, model name, and path to the audio file as input. The function first opens the audio file and creates a multipart form data payload containing the model name and the audio file. It then sends a POST request to the API endpoint with the multipart form data payload. The response is parsed as a struct containing the transcribed text.  
  
#### Wrong Password Handling  
  
The `GenerateCompletionWithPWD` function checks if the provided user password matches the stored password. If they match, it calls the `GenerateCompletion` function to generate the chat completion. Otherwise, it returns an error indicating that the password is incorrect.  
  
#### Text Cleaning  
  
The `cleanText` function removes the "[BLANK_AUDIO]" string from the input text if it is present and the input text is empty after removing the "[BLANK_AUDIO]" string. Otherwise, it returns the input text with the "[BLANK_AUDIO]" string removed.  
  
  
  
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
  
### Code Summary:  
  
#### SetupSequenceWithKey Function:  
  
This function is responsible for setting up the AI sequence for a given user. It takes the following parameters:  
  
- `bot`: Telegram bot API instance.  
- `user`: Database record containing user information.  
- `language`: Language preference for the user.  
- `ctx`: Context for the operation.  
- `spwd`: Password for authentication with the AI model.  
- `ai_endpoint`: Endpoint for communication with the AI model.  
  
The function first acquires a mutex lock to ensure thread safety. It then retrieves the user's GPT key and model from the `user.AiSession` field. It also retrieves the user's ID and network from the `user` struct.  
  
The function then uses a switch statement to handle different language preferences. For each language, it calls the `tryLanguage` function to generate a response from the AI model. The response is then sent to the user via the Telegram bot.  
  
#### tryLanguage Function:  
  
This function takes the following parameters:  
  
- `user`: Database record containing user information.  
- `language`: Language preference for the user.  
- `languageCode`: Code representing the language (0 - default, 1 - Russian, 2 - English).  
- `ctx`: Context for the operation.  
- `ai_endpoint`: Endpoint for communication with the AI model.  
- `spwd`: Password for authentication with the AI model.  
- `upwd`: User's password for authentication with the AI model.  
  
The function first constructs a language prompt based on the `languageCode`. It then calls the `GenerateCompletionWithPWD` function to generate a response from the AI model. The response is then logged and returned as a string.  
  
#### GenerateCompletionWithPWD Function:  
  
This function is responsible for generating a response from the AI model using the provided password. It takes the following parameters:  
  
- `languagePromt`: Prompt to be sent to the AI model.  
- `model`: AI model to use for generating the response.  
- `ai_endpoint`: Endpoint for communication with the AI model.  
- `spwd`: Password for authentication with the AI model.  
- `upwd`: User's password for authentication with the AI model.  
  
The function constructs a request to the AI model using the provided parameters and sends it to the specified endpoint. The response is then returned.  
  
#### LogResponse Function:  
  
This function is responsible for logging the response from the AI model. It takes the response as input and logs it to the console.  
  
  
  
# lib/localai/startDialogSequence.go  
## Package: localai  
  
### Imports:  
  
* context  
* log  
* db "github.com/JackBekket/hellper/lib/database"  
* tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
### External Data, Input Sources:  
  
* `db.UsersMap`: A map containing user data from the database.  
* `ai_endpoint`: A string representing the endpoint for the AI model.  
  
### Code Summary:  
  
#### errorMessage Function:  
  
This function handles errors that occur during the process of creating a request. It logs the error, sends an error message to the user via Telegram, and removes the user from the database.  
  
#### StartDialogSequence Function:  
  
This function initiates a dialog sequence with the user. It retrieves the user's data from the database, logs the GPT model and prompt, and calls the GenerateCompletion function to generate a response from the AI model. If an error occurs during the process, it calls the errorMessage function. Otherwise, it logs the response, formats it for display, and sends it to the user via Telegram. Finally, it updates the user's dialog status in the database.  
  
#### LogResponse Function:  
  
This function logs the full response object, including its creation time, ID, model, object, choices, and usage information.  
  
  
  
