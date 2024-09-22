# Package: dialog

### Imports:
- log
- github.com/JackBekket/hellper/lib/bot/command
- github.com/JackBekket/hellper/lib/database
- github.com/JackBekket/hellper/lib/langchain
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

### External Data, Input Sources:
- Updates from Telegram bot API
- Command data from Telegram bot API
- Database for user information

### Summary:
The `HandleUpdates` function is responsible for handling incoming updates from the Telegram bot API and managing user interactions. It iterates through the updates and processes each message based on the command provided by the user.

#### Command Handling:
- `/image`: Generates an image based on the provided prompt or a default prompt if none is given.
- `/restart`: Restarts the user's session by deleting their data from the database.
- `/help`: Displays a help message with available commands.
- `/search_doc`: Searches for documents based on the provided prompt and returns the top 3 results.
- `/rag`: Performs a RAG (Retrieval Augmented Generation) task based on the provided prompt.
- `/instruct`: Calls a local AI model to generate content based on the provided prompt and user's AI session settings.
- `/usage`: Displays the usage statistics for the user.
- `/helper`: Sends a media helper message to the user.

#### User Interaction:
- The function checks if the user is already in the database and adds them if not.
- It then determines the user's dialog status and handles the corresponding interaction based on the status.
- The dialog status is updated as the user progresses through the interaction, and the function logs the user's status and other relevant information.

#### Network and Model Selection:
- The function allows the user to choose their preferred network and AI model.
- It handles the selection process and updates the user's AI session settings accordingly.

#### Connecting to AI:
- The function connects the user to the chosen AI model using the provided API key and base URL.

#### Dialog Sequence:
- The function manages the dialog sequence by calling the appropriate functions based on the user's dialog status and input.

This summary provides a comprehensive overview of the `HandleUpdates` function and its role in managing user interactions and dialog flow within the `dialog` package.

