# command

## Summary

This code package defines a function called AddAdminToMap that handles the authorization of new administrators for a Telegram bot. It takes an adminKey and an updateMessage as input. The function first extracts the chatID and username from the updateMessage. Then, it creates a new user entry in the UsersMap, which is a map of users and their corresponding information. The new user entry includes the chatID, username, dialog status, admin status, and an AiSession object containing the adminKey.

The function then logs the authorized username and sends a confirmation message to the user. Finally, it sends another message with a one-time reply keyboard containing a button for selecting the desired GPT model (in this case, GPT-3.5).



## Summary

This code package defines a function called `AddNewUserToMap` that adds a new user to the database and assigns them a "Dialog_status" of 0. The function takes an update message as input, extracts the user's ID and username, and creates a new user object with these values. The user object is then added to the database using the `database.AddUser` function.

The function also logs the new user's ID and username to the console. Additionally, it sends a welcome message to the new user, including a "Start!" button in the reply markup.

The code package also includes imports for the `log` package and the `database` package, as well as the `tgbotapi` package for interacting with the Telegram Bot API.



## Summary

This code package provides a command-line interface for interacting with an AI model. It includes functions for handling user input, managing dialog flow, and connecting to various AI services.

The package starts by defining a `Commander` struct that manages the interaction with the user. It includes methods for handling different user inputs, such as choosing a network, selecting a model, and providing an API key. The package also includes functions for connecting to different AI services, such as OpenAI, LocalAI, and VastAI.

The `InputYourAPIKey` method prompts the user to enter their API key, which is then stored in the user's profile. The `ChooseNetwork` method allows the user to select the AI service they want to use, and the `HandleNetworkChoose` method handles the user's choice by setting the appropriate network and model selection parameters.

The `ChooseModel` method prompts the user to select a model from the available options, and the `HandleModelChoose` method handles the user's choice by setting the model name and API endpoint. The `AttachKey` method is used to store the user's API key in their profile.

The package also includes functions for generating images using Stable Diffusion, uploading images to Telegram, and handling dialog flow. The `DialogSequence` method is responsible for managing the conversation between the user and the AI model, while the `GenerateNewImageLAI_SD` method generates an image using Stable Diffusion and uploads it to Telegram.

Overall, this code package provides a comprehensive set of tools for interacting with AI models, including user input handling, dialog flow management, and integration with various AI services.

## Summary

This code package defines a function called `CheckAdmin` that manages user permissions and loads keys from the environment into the database. It takes an `adminData` map containing admin information and an `updateMessage` as input.

The function first checks if the user is an admin by comparing the user's ID with the IDs in the `adminData` map. If the user is an admin, it checks if the corresponding GPT key is present in the environment. If the key is found, it adds the admin to a map and returns. If the key is missing, it sends a message to the user and directs the function to the case for non-admin users.

If the user is not an admin, the function adds the user to a map and returns. This ensures that only admins with valid GPT keys can access certain features or resources.

In summary, this code package provides a mechanism for managing user permissions and loading keys from the environment into the database. It ensures that only admins with valid GPT keys can access certain features or resources.



## Summary

This code package defines a set of message templates for a bot that interacts with either an OpenAI API or a local AI node. The templates cover various aspects of the bot's functionality, including greetings, authorization prompts, model selection, network selection, and help commands. The package also includes templates for experimental features such as document searching, Retrival-Augmented Generation, and image generation.

The message templates are stored in a map called `msgTemplates`, where each key represents a specific message and its corresponding value is the message content. The package provides a comprehensive set of templates to guide the bot's interactions with users and manage its various functionalities.



## Summary

This code package defines a Commander struct that manages interactions with a Telegram bot and a database of users. The Commander struct has three fields: a pointer to a Telegram bot API object, a map of user IDs to database.User objects, and a context object.

The package provides a constructor function, NewCommander, which takes a Telegram bot API object, a map of user IDs to database.User objects, and a context object as input and returns a new Commander instance.

The package also includes a function, GetCommander, which is not fully implemented in the provided code.



## Summary

This code package provides a set of functions for rendering menus and handling user interactions within a Telegram bot. It utilizes the `tgbotapi` library to interact with the Telegram API.

The package includes functions for rendering model menus for different types of AI models, such as GPT-3.5, GPT-4, and various language models. These functions create and send messages to the user, along with a one-time reply keyboard containing buttons for selecting the desired model.

Additionally, the package includes a function for rendering a language menu, which allows the user to choose a language for the bot to understand and respond in. This function sends a message to the user with a one-time reply keyboard containing buttons for selecting the desired language.

In summary, this code package provides a framework for creating a Telegram bot that can handle user interactions related to selecting AI models and languages. It uses the `tgbotapi` library to interact with the Telegram API and provides functions for rendering menus and handling user input.



## Summary

This code package provides a set of functions for a Telegram bot that can perform various tasks, such as providing help messages, searching for documents, and retrieving usage statistics. The package also includes a function for sending media files, such as videos, to users.

The package starts by defining a Commander struct, which contains methods for handling different commands. The HelpCommandMessage method is responsible for sending a help message to the user when they request it. The SearchDocuments method allows users to search for documents based on a given prompt and returns the results along with their scores. The RAG method implements a Retrieval-Augmented Generation approach, which combines the results of a semantic search with a language model to generate a more comprehensive response.

The GetUsage method retrieves and displays the usage statistics for a particular user, including the number of prompt tokens, completion tokens, and total tokens used. The SendMediaHelper method is responsible for sending a random video file from a specified media directory to the user.

Overall, this code package provides a comprehensive set of functions for a Telegram bot that can assist users with various tasks, including document search, usage tracking, and media sharing.



