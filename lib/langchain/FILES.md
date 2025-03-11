# lib/langchain/handler.go  
# Package Name and Imports  
The package name is **langchain**. The imports are:  
- `context`  
- `encoding/json`  
- `log`  
- `db` from `github.com/JackBekket/hellper/lib/database`  
- `github.com/tmc/langchaingo/llms`  
- `github.com/tmc/langchaingo/schema`  
  
## External Data and Input Sources  
The external data and input sources include:  
- `db.User` from the database  
- `llms.ContentResponse` and `llms.MessageContent` from the `llms` package  
- `schema.AgentAction`, `schema.AgentFinish`, and other schema types from the `schema` package  
  
## TODO Comments  
The list of TODO comments is:  
- Implement the `HandleAgentAction` method  
- Implement the `HandleAgentFinish` method  
- Implement the `HandleChainEnd` method  
- Implement the `HandleChainError` method  
- Implement the `HandleChainStart` method  
- Implement the `HandleLLMError` method  
- Implement the `HandleLLMGenerateContentStart` method  
- Implement the `HandleLLMStart` method  
- Implement the `HandleRetrieverEnd` method  
- Implement the `HandleRetrieverStart` method  
- Implement the `HandleStreamingFunc` method  
- Implement the `HandleToolEnd` method  
- Implement the `HandleToolError` method  
- Implement the `HandleToolStart` method  
  
## Major Code Parts  
### ChainCallbackHandler Struct  
The `ChainCallbackHandler` struct has several methods that implement the `callbacks.Handler` interface. These methods include:  
- `HandleAgentAction`  
- `HandleAgentFinish`  
- `HandleChainEnd`  
- `HandleChainError`  
- `HandleChainStart`  
- `HandleLLMError`  
- `HandleLLMGenerateContentStart`  
- `HandleLLMStart`  
- `HandleRetrieverEnd`  
- `HandleRetrieverStart`  
- `HandleStreamingFunc`  
- `HandleToolEnd`  
- `HandleToolError`  
- `HandleToolStart`  
  
### LogResponseContentChoice Function  
The `LogResponseContentChoice` function logs the content and other information from the `llms.ContentResponse` object. It also updates the user's usage information in the database.  
  
### HandleLLMGenerateContentEnd Method  
The `HandleLLMGenerateContentEnd` method calls the `LogResponseContentChoice` function to log the response content.  
  
  
  
# lib/langchain/langchain.go  
# Package/Component Name  
The package/component name is **langchain**.  
  
## Imports  
The following imports are used in the code:  
* `context`  
* `fmt`  
* `log`  
* `github.com/tmc/langchaingo/llms`  
* `github.com/tmc/langchaingo/llms/openai`  
  
## External Data/Input Sources  
The external data/input sources are:  
* `base_url`: a string representing the base URL for the API  
* `promt`: a string representing the prompt for the language model  
* `model_name`: a string representing the name of the language model  
* `localAIToken`: a string representing the API token for authentication  
* `network`: a string representing the network to use (either "local" or "openai")  
* `options`: a variable number of `llms.CallOption` arguments  
  
## TODO Comments  
The following TODO comments are found in the code:  
* None explicitly marked as TODO, but there is an obsolete comment indicating that `LANGGRAPH.GO` should be used instead.  
  
## Code Summary  
### Function GenerateContentInstruction  
The `GenerateContentInstruction` function generates content based on a single prompt without memory and context. It takes in several parameters, including `base_url`, `promt`, `model_name`, `localAIToken`, `network`, and `options`. The function returns a string result and an error.  
  
### Language Model Initialization  
The function initializes a language model using the `openai` package, either with a local or openai network. It sets up the model with the provided `localAIToken`, `base_url`, and `model_name`.  
  
### Content Generation  
The function generates content using the `llms.GenerateFromSinglePrompt` function, passing in the context, language model, prompt, and options. The result is then printed to the console and returned as a string.  
  
  
  
# lib/langchain/langgraph.go  
# Package/Component Name  
The package/component name is **langchain**.  
  
## Imports  
The following imports are used in the code:  
* `github.com/JackBekket/hellper/lib/agent`  
* `github.com/JackBekket/hellper/lib/database` (aliased as **db**)  
* `github.com/tmc/langchaingo/llms/openai`  
  
## External Data/Input Sources  
The external data/input sources are:  
* **localAIToken**: a string representing the API token  
* **model_name**: a string representing the model name  
* **base_url**: a string representing the base URL  
* **user_promt**: a string representing the user prompt  
* **state**: a pointer to a **db.ChatSessionGraph** object  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Function RunNewAgent  
The **RunNewAgent** function creates a new agent and runs a thread with the given user prompt. It returns a **db.ChatSessionGraph** object, an output text string, and an error.  
  
### Function ContinueAgent  
The **ContinueAgent** function continues an existing agent and runs a thread with the given user prompt and conversation buffer. It returns a **db.ChatSessionGraph** object, an output text string, and an error.  
  
  
  
# lib/langchain/setupSequenceWithKey.go  
# Package Name and Imports  
The package name is **langchain**. The imports are:  
* `context`  
* `log`  
* `sync`  
* `db` from `github.com/JackBekket/hellper/lib/database`  
* `tgbotapi` from `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `db.User` struct  
* `tgbotapi.BotAPI` struct  
* `ai_endpoint` string  
* `language` string  
* `ctx` context.Context  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### SetupSequenceWithKey Function  
The `SetupSequenceWithKey` function sets up a sequence with a key for a given user and language. It takes in several parameters, including `bot`, `user`, `language`, `ctx`, and `ai_endpoint`. The function uses a mutex to lock and unlock the sequence setup process. It then checks the language and calls the `tryLanguage` function accordingly.  
  
### tryLanguage Function  
The `tryLanguage` function tries to determine the language of the user. It takes in several parameters, including `user`, `language`, `languageCode`, and `ai_endpoint`. The function uses a switch statement to determine the language prompt based on the language code. It then calls the `RunNewAgent` function to get the result and thread.  
  
### RunNewAgent Function  
The `RunNewAgent` function is not defined in the provided code, but it is called by the `tryLanguage` function. It takes in several parameters, including `gptKey`, `model`, `ai_endpoint`, and `languagePromt`.  
  
### Error Handling  
The code uses error handling to catch any errors that may occur during the sequence setup process. If an error occurs, the `errorMessage` function is called to send an error message to the user.  
  
  
  
# lib/langchain/startDialogSequence.go  
# Package Name and Imports  
The package name is `langchain`. The imports are:  
* `context`  
* `fmt`  
* `log`  
* `math/rand`  
* `os`  
* `path/filepath`  
* `io/fs`  
* `github.com/JackBekket/hellper/lib/agent`  
* `github.com/JackBekket/hellper/lib/database` (as `db`)  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (as `tgbotapi`)  
  
## External Data and Input Sources  
The external data and input sources are:  
* `../../media/` directory, which contains video files  
* `db.UsersMap`, which is a map of user IDs to user objects  
* `ai_endpoint`, which is a string representing the URL of the AI endpoint  
* `ctx`, which is a context object  
  
## TODO Comments  
The TODO comments are:  
* Investigate why meme videos with helper are not sent by the `errorMessage` function  
  
## Code Summary  
### Error Handling  
The `errorMessage` function handles errors by sending an error message to the user, sending a helper video, and removing the user from the database.  
  
### Dialog Sequence  
The `StartDialogSequence` function starts a dialog sequence with the user by:  
1. Locking the mutex to ensure thread safety  
2. Retrieving the user object from the `db.UsersMap`  
3. Logging the GPT model and prompt  
4. Continuing the agent session using the `ContinueAgent` function  
5. Sending the AI response to the user  
6. Updating the user's dialog status and session usage  
7. Saving the user conversation to the database  
  
### Database Updates  
The `ds.UpdateHistory` function is used to update the history of the conversation in the database.  
  
  
  
