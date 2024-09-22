# dialog

This package provides a dialog management system for a Telegram bot. It handles user interactions, manages dialog states, and provides various commands for users to interact with the bot.

## File structure

- lib/bot/dialog/dialog.go

## Code summary

The `HandleUpdates` function is the main entry point for handling user interactions. It iterates over a channel of Telegram updates and processes each update. For each update, it extracts the chat ID, retrieves the user from a database, and handles the user's command or dialog state.

The package provides several commands for users to interact with the bot, including:

- `/image`: Generates an image based on a prompt.
- `/restart`: Restarts the user's session.
- `/help`: Displays help information.
- `/search_doc`: Searches for documents based on a prompt.
- `/rag`: Performs RAG based on a prompt.
- `/instruct`: Calls a local AI model to generate content based on a prompt.
- `/usage`: Displays usage information.
- `/helper`: Sends a media helper message.

The package also manages dialog states, which are used to guide the user through a sequence of interactions. The dialog states are:

- Status 0: Chooses a network.
- Status 1: Handles network choice.
- Status 2: Prompts the user to input their API key.
- Status 3: Chooses a model.
- Status 4: Handles model choice.
- Status 5: Connects to the AI with the chosen language.
- Status 6: Performs the dialog sequence.

The package uses a database to store user information and dialog states. It also uses a local AI model for content generation and a network for communication with the AI.

