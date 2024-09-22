# command

This package provides the logic for handling commands in the Telegram bot. It includes functions for managing user interactions, handling different types of commands, and managing the bot's state.

## Functions

1. `HelpCommandMessage(updateMessage *tgbotapi.Message)`:
   - Gets the chatID from the updateMessage.
   - Gets the user from the database using the chatID.
   - Creates a new message using the msgTemplates["help_command"] template.
   - Sends the message to the user using the bot.

2. `SearchDocuments(chatID int64, promt string, maxResults int)`:
   - Loads environment variables using godotenv.Load().
   - Gets the PG_LINK and AI_BASEURL from the environment variables.
   - Gets the user from the database using the chatID.
   - Gets the api_token from the user's AiSession.GptKey.
   - Gets the vector store using embeddings.GetVectorStore.
   - Performs semantic search using embeddings.SemanticSearch.
   - Sends the results to the user using the bot.

3. `RAG(chatID int64, promt string, maxResults int)`:
   - Gets the user from the database using the chatID.
   - Loads environment variables using godotenv.Load().
   - Gets the PG_LINK and AI_BASEURL from the environment variables.
   - Gets the api_token from the user's AiSession.GptKey.
   - Gets the vector store using embeddings.GetVectorStore.
   - Performs RAG using embeddings.Rag.
   - Sends the result to the user using the bot.

4. `GetUsage(chatID int64)`:
   - Gets the user from the database using the chatID.
   - Gets the promt_tokens, completion_tokens, and total_tokens from the user's AiSession.Usage.
   - Sends the usage information to the user using the bot.

5. `SendMediaHelper(chatID int64)`:
   - Gets a list of files in the media directory.
   - Selects a random file.
   - Opens the video file.
   - Creates a new video message.
   - Sends the video message to the user using the bot.

6. `InputYourAPIKey(updateMessage *tgbotapi.Message)`:
   - Gets the chatID and user from the updateMessage.
   - Sets the user's DialogStatus to 3.
   - Sends a message to the user asking for their API key.

7. `ChooseNetwork(updateMessage *tgbotapi.Message)`:
   - Gets the chatID and user from the updateMessage.
   - Sets the user's DialogStatus to 1.
   - Sends a message to the user asking them to choose a network (openai, localai, vastai).

8. `HandleNetworkChoose(updateMessage *tgbotapi.Message)`:
   - Gets the chatID, network, and user from the updateMessage.
   - Sets the user's network and AiSession.AI_Type based on the chosen network.
   - Sets the user's DialogStatus to 2.

9. `ChooseModel(updateMessage *tgbotapi.Message)`:
   - Gets the chatID, model_name, and user from the updateMessage.
   - Sets the user's AiSession.GptModel based on the chosen model.
   - Sets the user's DialogStatus to 4.

10. `HandleModelChoose(updateMessage *tgbotapi.Message)`:
   - Gets the chatID, model_name, and user from the updateMessage.
   - Sets the user's AiSession.GptModel based on the chosen model.
   - Sets the user's DialogStatus to 5.

11. `attachModel(model_name string, chatID int64)`:
   - Gets the user from the database using the chatID.
   - Sets the user's AiSession.GptModel to the given model_name.
   - Sends a message to the user confirming the chosen model.

12. `AttachKey(gpt_key string, chatID int64)`:
   - Gets the user from the database using the chatID.
   - Sets the user's AiSession.GptKey to the given gpt_key.

13. `ChangeDialogStatus(chatID int64, ds int8)`:
   - Gets the user from the database using the chatID.
   - Updates the user's DialogStatus to the given ds value.

14. `WrongModel(updateMessage *tgbotapi.Message)`:
   - Gets the chatID and user from the updateMessage.
   - Sends a message to the user asking them to choose a valid model.

15. `WrongNetwork(updateMessage *tgbotapi.Message)`:
   - Gets the chatID and user from the updateMessage.
   - Sends a message to the user asking them to choose a valid network.

16. `ConnectingToAiWithLanguage(updateMessage *tgbotapi.Message, ai_endpoint string)`:
   - Gets the chatID, language, and user from the updateMessage.
   - Sets up the AI sequence based on the chosen network and language.

17. `DialogSequence(updateMessage *tgbotapi.Message, ai_endpoint string)`:
   - Gets the chatID, promt, and user from the updateMessage.
   - Starts the dialog sequence using the given promt and ai_endpoint.

18. `GenerateNewImageLAI_SD(promt string, chatID int64, bot *tgbotapi.BotAPI)`:
   - Generates a new image using the given promt and chatID.
   - Sends the image to the user using the bot.

19. `GetUsersDb() (map[int64]db.User)`:
   - Returns the current users database.

20. `GetUser(id int64) (db.User)`:
   - Gets the user from the database using the given ID.

21. `RenderModelMenuOAI(chatID int64)`:
   - Sends a message to the user with a one-time reply keyboard containing the available GPT models.

22. `RenderModelMenuLAI(chatID int64)`:
   - Sends a message to the user with a one-time reply keyboard containing the available local AI models.

23. `RenderModelMenuVAI(chatID int64)`:
   - Sends a message to the user with a one-time reply keyboard containing the available Vast AI models.

24. `RenderLanguage(chat_id int64)`:
   - Sends a message to the user asking them to choose a language or send "Hello" in their desired language.

25. `AddNewUserToMap(updateMessage *tgbotapi.Message)`:
   - Adds a new user to the database and sends a welcome message.

26. `CheckAdmin(adminData map[string]env.AdminData, updateMessage *tgbotapi.Message)`:
   - Checks if the user is an admin and adds them to the database if they are.

27. `NewCommander(bot *tgbotapi.BotAPI, usersDb map[int64]database.User, ctx context.Context) *Commander`:
   - Creates a new Commander instance with the given bot, usersDb, and ctx.

28. `AddAdminToMap(adminKey string, updateMessage *tgbotapi.Message)`:
   - Adds an admin to the database and sends a confirmation message.


