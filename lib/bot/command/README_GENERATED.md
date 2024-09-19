# lib/bot/command

This package contains the logic for handling commands in the bot. It includes functions for managing user accounts, handling different types of commands, and interacting with various AI models.

## File Structure

- lib/bot/command/addNewUsertoMap.go
- lib/bot/command/addAdminTomap.go
- lib/bot/command/cases.go
- lib/bot/command/checkAdmin.go
- lib/bot/command/msgTemplates.go
- lib/bot/command/newCommander.go
- lib/bot/command/ui.go
- lib/bot/command/utils.go

## Code Summary

1. **User Management:**
   - The package includes functions for adding new users to the database and managing their accounts.
   - `AddNewUserToMap` function creates a new user entry in the database and sends a welcome message to the user.
   - `AddAdminToMap` function handles adding an admin user to the database and sets their DialogStatus to 2.

2. **Command Handling:**
   - The `cases.go` file contains various functions for handling different types of commands, such as choosing a network, selecting a model, and handling user input.
   - `RenderModelMenuLAI`, `RenderModelMenuOAI`, and `RenderModelMenuVAI` functions display menus for users to choose their preferred AI model.
   - `ChooseNetwork` and `HandleNetworkChoose` functions handle user input for choosing between OpenAI and LocalAI.

3. **AI Model Interaction:**
   - The package includes functions for interacting with different AI models, such as OpenAI and LocalAI.
   - `AttachKey` and `ChangeDialogStatus` functions manage user authentication and DialogStatus.
   - `attachModel` function sets the user's preferred AI model.

4. **Additional Features:**
   - The `utils.go` file contains functions for sending media, searching documents, and performing RAG (Retrieval-Augmented Generation).
   - `SendMediaHelper` function sends a random media file to the user.
   - `SearchDocuments` and `RAG` functions perform semantic search and RAG, respectively.

5. **Environment Variables:**
   - The package uses environment variables for configuration, such as `PG_LINK` and `AI_BASEURL`.

6. **Edge Cases:**
   - The package handles edge cases such as missing environment variables and unregistered users.

7. **Dead Code:**
   - There is no apparent dead code in the package.

8. **Missing Information:**
   - The package does not provide information on how to launch the application or any specific configuration files.

9. **Overall Logic:**
   - The package provides a comprehensive set of functions for managing user accounts, handling commands, and interacting with AI models. It also includes additional features such as sending media, searching documents, and performing RAG. The package is well-structured and easy to understand.

