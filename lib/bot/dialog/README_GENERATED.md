# dialog

This package provides a dialog handler for a Telegram bot. It handles incoming updates from the Telegram API and routes them to the appropriate command handler. The package also manages user data, such as chat IDs and AI session information.

## Environment variables

- None specified

## Flags

- None specified

## Cmdline arguments

- None specified

## Files and their paths

- lib/bot/dialog/dialog.go

## Edge cases

- None specified

## Project package structure

- lib/bot/dialog/dialog.go

## Code summary

The `HandleUpdates` function is the main entry point for the dialog handler. It iterates over incoming updates from the Telegram API and processes each update. For each update, it extracts the chat ID and retrieves the corresponding user data from the command handler's user database. If the user is not found in the database, a new user entry is created.

The function then checks the command of the update and routes it to the appropriate command handler. The available commands are:

- `/image`: Generates an image based on the provided prompt or a default prompt if no prompt is given.
- `/restart`: Restarts the user's session by deleting their data from the user database.
- `/help`: Displays a help message with available commands.
- `/search_doc`: Searches for documents based on the provided prompt and returns the top 3 results.
- `/rag`: Performs a RAG (Retrieval Augmented Generation) task based on the provided prompt.
- `/instruct`: Generates content based on the provided prompt using the user's selected AI model and API key.
- `/usage`: Displays the usage statistics for the user.
- `/helper`: Sends a media helper message to the user.

The `HandleUpdates` function also handles the logic for generating images using the LAI_SD model and searching for documents using a search engine. It also manages the user's AI session information, such as the selected AI model, API key, and network.

The package also includes helper functions for generating help messages, searching for documents, performing RAG tasks, and sending media helper messages.

