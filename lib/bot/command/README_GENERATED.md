Here is a markdown summary of the provided package code:

**command**
================

The `command` package provides a set of functions and structs to manage a Telegram bot. It handles various commands and interactions with users.

**Environment Variables and Flags:**

* `PG_LINK`
* `AI_BASEURL`
* `OPENAI_API_KEY`
* `IMAGE_GENERATION_MODEL`
* `IMAGE_GENERATION_SUFFIX`

**Cmd/CLI Arguments:**

* None found

**Files and Paths:**

* `addNewUsertoMap.go`
* `cases.go`
* `commandHandler.go`
* `msgTemplates.go`
* `newCommander.go`
* `ui.go`
* `utils.go`

**Launch Options:**

* Run the `NewCommander` function to initialize a new `Commander` instance.
* Run the `HandleCommands` function to handle user input and update the dialog status.

**Edge Cases:**

* None found

**Summary:**

The `command` package provides a set of functions and structs to manage a Telegram bot. It handles various commands and interactions with users. The package includes the following components:

* `AddNewUserToMap` function: adds a new user to the database and assigns a "Dialog_status" of 0.
* `Commander` struct: manages the bot's interactions with users.
* `CommandHandler` function: handles various commands and updates the dialog status.
* `msgTemplates` map: provides pre-defined message templates for the bot.
* `NewCommander` function: initializes a new `Commander` instance.
* `ui` functions: render menus and send messages to users.
* `utils` functions: perform various utility tasks, such as image processing and file management.

**