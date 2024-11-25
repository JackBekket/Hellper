## Package: localai

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