# localai

This package provides a set of functions to manage and interact with a local AI model. It includes functions for setting up a sequence with a key, starting a dialog sequence, and generating completions and images.

## File Structure

- lib/localai/setupSequenceWithKey.go
- lib/localai/startDialogSequence.go
- lib/localai/localai.go

## Code Summary

### SetupSequenceWithKey

This function takes a bot API, user, language, context, and passwords as input. It retrieves the user's GPT key and model from the session, and then calls the `tryLanguage` function to generate a response based on the provided language. The response is then sent to the user, and the user's dialog status is updated.

### StartDialogSequence

This function takes a bot API, chat ID, prompt, context, and AI endpoint as input. It retrieves the user from the database, logs the GPT model and prompt, and generates a completion using the prompt and GPT model. If there is an error, it calls the `errorMessage` function. If there is no error, it logs the response, sends the response to the user, and updates the user's dialog status.

### LocalAI

This package provides a set of functions to manage and interact with a local AI model. It includes functions for setting up a sequence with a key, starting a dialog sequence, and generating completions and images.

### GenerateCompletion

This function takes a prompt, model name, and URL as input. It creates a ChatRequest with the model name, prompt, and temperature, converts the ChatRequest to JSON, sends a POST request to the URL with the JSON data, reads the response body, parses the JSON response into a ChatResponse, and returns the ChatResponse.

### GenerateCompletionWithPWD

This function takes a prompt, model name, URL, and passwords as input. It checks if the passwords match, calls GenerateCompletion with the prompt, model name, and URL, and returns the result.

### GenerateImageStableDissusion

This function takes a prompt and size as input. It creates a payload with the prompt and size, converts the payload to JSON, sends a POST request to the URL with the JSON data, reads the response body, parses the JSON response into a GenerationResponse, and returns the image URL.

### UploadToTelegraph

This function takes a file name as input. It gets the absolute path to the file, opens the file, uploads the file to telegraph, and returns the link.

### DeleteFromTemp

This function takes a file name as input. It gets the absolute path to the file and deletes the file from the local machine.

## Edge Cases

- If the passwords do not match, the `GenerateCompletionWithPWD` function will return an error.
- If the user is not found in the database, the `StartDialogSequence` function will return an error.
- If there is an error during the generation of the completion or image, the corresponding functions will return an error.

## Conclusion

The `localai` package provides a set of functions to manage and interact with a local AI model. It includes functions for setting up a sequence with a key, starting a dialog sequence, and generating completions and images. The package also includes functions for uploading files to telegraph and deleting files from the local machine.