# Package: command

### Imports:
- `context`
- `github.com/JackBekket/hellper/lib/database`
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`

### External Data and Input Sources:
- `bot`: A Telegram bot API instance.
- `usersDb`: A map of user IDs to database.User objects.
- `ctx`: A context object.

### Code Summary:
The provided code defines a set of functions that handle different cases or scenarios related to user interactions with the bot. These functions are designed to manage user authentication, model selection, and other aspects of the bot's functionality.

#### Case 0:
This function handles the case where a user needs to provide their OpenAI API key or password for local AI authentication. It prompts the user to enter their credentials and stores them in the database.

#### Case 1:
This function handles the case where a user needs to choose a model to use. It presents a menu of available models and allows the user to select one.

#### Case 2:
This function handles the case where a user needs to choose a network to work with, either OpenAI or local AI. It presents a menu of available networks and allows the user to select one.

#### Case 3:
This function handles the case where a user needs to perform a semantic search on a given prompt. It takes the chat ID, prompt, and maximum number of results as input and performs the search using the embeddings library.

#### Case 4:
This function handles the case where a user needs to perform Retrieval-Augmented Generation (RAG). It takes the chat ID, prompt, and maximum number of results as input and performs the RAG using the embeddings library.

#### Case 5:
This function handles the case where a user needs to retrieve and display their usage statistics. It takes the chat ID as input and retrieves the user object from the database. It then extracts the prompt tokens, completion tokens, and total tokens from the user object and sends them to the user.

#### Case 6:
This function handles the case where a user needs to send a random video from the media directory. It first reads the files in the media directory and selects a random file. Then, it opens the video file, creates a new video message, and sends it to the user.

The provided code provides a comprehensive set of functions to handle various user interactions and scenarios, ensuring that the bot can effectively manage user authentication, model selection, and other aspects of its functionality.

