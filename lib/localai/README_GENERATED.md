# localai

This package provides a set of functions to interact with a GPT model and generate text completions. It also includes functions for uploading images to Telegraph and managing user sessions.

## Project Package Structure

```
lib/localai/
├── localai.go
└── startDialogSequence.go
```

## Code Summary

### localai.go

This file contains functions for generating text completions using a GPT model and uploading images to Telegraph.

1. GenerateCompletion(prompt, modelName, url)
   - Takes a prompt, model name, and URL as input.
   - Creates a ChatRequest struct with the given prompt, model name, and temperature.
   - Converts the ChatRequest struct to JSON.
   - Sends a POST request to the given URL with the JSON data.
   - Reads the response body and parses it as a ChatResponse struct.
   - Returns the ChatResponse struct and any errors encountered.

2. GenerateCompletionWithPWD(prompt, modelName, url, s_pwd, u_pwd)
   - Takes a prompt, model name, URL, stored password, and user-provided password as input.
   - Checks if the user-provided password matches the stored password.
   - If the passwords match, calls GenerateCompletion with the given parameters and returns the result.
   - If the passwords don't match, returns an error.

3. GenerateImageStableDissusion(prompt, size)
   - Takes a prompt and size as input.
   - Creates a payload with the given prompt and size.
   - Converts the payload to JSON.
   - Sends a POST request to the given URL with the JSON data.
   - Reads the response body and parses it as a GenerationResponse struct.
   - Returns the image URL from the GenerationResponse struct and any errors encountered.

4. UploadToTelegraph(fileName)
   - Takes a file name as input.
   - Gets the absolute path to the file.
   - Opens the file using the absolute path.
   - Uploads the file to Telegraph using the telegraph.Upload function.
   - Returns the uploaded file link.

5. deleteFromTemp(fileName)
   - Takes a file name as input.
   - Gets the absolute path to the file.
   - Deletes the file from the temporary directory.

### startDialogSequence.go

This file contains functions for managing user sessions and handling dialog sequences.

1. errorMessage(err, bot, user)
   - Logs the error.
   - Sends an error message to the user via the bot.
   - Sends a message instructing the user to recreate the client and initialize a new session.
   - Removes the user from the database.

2. StartDialogSequence(bot, chatID, promt, ctx, ai_endpoint)
   - Acquires a lock.
   - Retrieves the user from the database.
   - Logs the GPT model and prompt.
   - Generates a completion using the prompt and GPT model.
   - Handles errors by calling errorMessage.
   - Logs the response.
   - Sends the response to the user via the bot.
   - Updates the user's dialog status.
   - Releases the lock.

3. LogResponse(resp)
   - Logs various details about the response, including the full response object, created timestamp, ID, model, object, Choices[0], and usage statistics.

### setupSequenceWithKey.go

This file contains functions for setting up user sessions and handling dialog sequences with a key.

1. SetupSequenceWithKey(bot, user, language, ctx, spwd, ai_endpoint)
   - Acquires a lock on mu.
   - Assigns chatID to user.ID.
   - Assigns gptKey to user.AiSession.GptKey.
   - Logs user GPT key from session.
   - Logs user model from session.
   - Logs upwd.
   - Checks if language is "English", "Russian" or default.
   - Calls tryLanguage function based on language.
   - Updates user.DialogStatus to 4.
   - Updates db.UsersMap with user.

2. tryLanguage(user, language, languageCode, ctx, ai_endpoint, spwd, upwd)
   - Assigns languagePromt based on languageCode.
   - Logs languagePromt.
   - Assigns model to user.AiSession.GptModel.
   - Calls GenerateCompletionWithPWD function.
   - Logs response.
   - Returns answer from response.

3. GenerateCompletionWithPWD(languagePromt, model, ai_endpoint, spwd, upwd)
   - TODO: implement GenerateCompletionWithPWD.

4. errorMessage(err, bot, user)
   - TODO: implement errorMessage.

