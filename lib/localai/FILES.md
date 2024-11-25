# lib/localai/localai.go  
## localai package  
  
This package provides functions for interacting with various AI models, including chat models and image generation models. It also includes a function for transcribing audio using the Whisper model.  
  
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
  
- API endpoints for chat models and image generation models (e.g., "http://localhost:8080/v1/chat/completions")  
- Whisper model for audio transcription  
- Environment variable "OPENAI_API_KEY" for authentication with the API  
  
### Code Summary  
  
1. Chat Request and Response Structures: The package defines structures for chat requests and responses, including fields for model, messages, temperature, created, object, ID, model, choices, and usage statistics.  
  
2. Generation Response Structure: A structure for generation responses is also defined, including fields for created, ID, data, and usage.  
  
3. Wrong Password Error: A custom error type, WrongPwdError, is defined to handle incorrect passwords.  
  
4. Main Function: The main function demonstrates how to use the package by sending a chat request to a chat model and printing the assistant's response.  
  
5. GenerateCompletion Function: This function takes a prompt, model name, and API URL as input and returns a chat response. It creates a chat request, converts it to JSON, sends the request to the API, and parses the response.  
  
6. GenerateCompletionWithPWD Function: This function is similar to GenerateCompletion but also takes a secret password as input and returns an error if the password is incorrect.  
  
7. GenerateImageStableDiffusion Function: This function takes a prompt, size, API URL, and model as input and returns an image URL. It creates a payload with the prompt, size, and model, sends a POST request to the API, and parses the response to extract the image URL.  
  
8. TranscribeWhisper Function: This function takes a URL, model, and path to an audio file as input and returns the transcribed text. It opens the audio file, creates a multipart request body, sends the request to the API, and parses the response to extract the transcribed text.  
  
9. cleanText Function: This function removes "[BLANK_AUDIO]" from the output of the TranscribeWhisper function.  
  
### End of Output  
  
  
  
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
  
  
  
