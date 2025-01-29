Here is a markdown summary of the provided package code:

**command**
================

### Functionality

The `command` package provides a set of functions for managing a Telegram bot. It includes functionality for adding admins to a map, handling user input, and sending messages to users.

### Configuration

* Environment variables:
	+ `adminKey` (string)
* Flags/CommandLine Arguments:
	+ None
* Files/Paths:
	+ None

### Launching the Application

The application can be launched by running the `main` function in the `newCommander` package.

### Edge Cases

* None

### Package Structure

```
lib/
bot/
command/
addAdminTomap.go
addNewUsertoMap.go
cases.go
checkAdmin.go
msgTemplates.go
newCommander.go
ui.go
utils.go
```

### Relations

The package uses the `db` package to interact with a database, and the `tgbotapi` package to send messages to a Telegram bot.

**Note:** The package structure and functionality suggest that it is part of a larger AI-powered chatbot system.

**command/newCommander.go**

The `NewCommander` function initializes a new `Commander` instance, which is used to manage the bot's interactions with users.

**command/ui.go**

The `RenderModelMenuLAI` and `RenderLanguage` functions render menus with inline keyboard buttons for selecting LLaMA-based models and language selection, respectively.

**command/utils.go**

The package provides various utility functions, including `HelpCommandMessage`, `SearchDocuments`, `RAG` (marked as OBSOLETE), `GetUsage`, `GenerateNewImageLAI_SD`, `SendMediaHelper`, and `DeleteFile`.

**command/cases.go**

The package provides a set of functions for handling different cases, including `InputYourAPIKey`, `ChooseModel`, `HandleModelChoose`, `DialogSequence`, `ChangeDialogStatus`, and `GetUsersDb` and `GetUser`.

**command/checkAdmin.go**

The `CheckAdmin` function updates the "dialogStatus" in the database based on the provided `adminData` and `updateMessage` inputs.

**command/msgTemplates.go**

The `msgTemplates` map is defined, which contains a set of pre-defined message templates for interacting with the user.

