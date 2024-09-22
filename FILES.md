# lib/langchain/setupSequenceWithKey.go  
package: langchain  
imports:  
- context  
- log  
- sync  
- db "github.com/JackBekket/hellper/lib/database"  
- tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
func SetupSequenceWithKey(  
	bot *tgbotapi.BotAPI,  
	user db.User,  
	language string,  
	ctx context.Context,  
	ai_endpoint string,  
) {  
	- Acquires a lock on the mutex  
	- Retrieves the chat ID, GPT key, and GPT model from the user's session  
	- Logs the retrieved GPT key and GPT model  
	- Determines the language based on the provided language string  
	- Calls the tryLanguage function to handle the language-specific logic  
	- Updates the user's dialog status and session usage  
	- Stores the updated user in the db.UsersMap  
	- Releases the lock on the mutex  
  
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context, ai_endpoint string) (string,*db.ChatSession, error) {  
	- Initializes a language prompt based on the language code  
	- Retrieves the GPT key, GPT model, and chat ID from the user's session  
	- Calls the StartNewChat function to initiate a new chat session  
	- Returns the chat result, thread, and any errors encountered  
  
  
  
# lib/langchain/startDialogSequence.go  
package: langchain  
imports: context, io/ioutil, log, math/rand, os, path/filepath, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User):  
	- logs the error  
	- sends an error message to the user  
	- sends a message instructing the user to recreate the client and initialize a new session  
	- gets a list of files in the media directory  
	- selects a random file  
	- opens the video file  
	- creates a new video message  
	- sends the video message  
	- removes the user from the database  
  
func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string):  
	- locks a mutex  
	- gets the user from the database  
	- logs the GPT model and prompt  
	- gets the current dialog thread  
	- continues the chat with the prompt  
	- if there's an error, calls errorMessage  
	- if successful, logs the AI response  
	- creates a new message with the response  
	- sends the message to the user  
	- updates the user's dialog status  
	- gets the session usage  
	- updates the user's session usage  
	- updates the user's dialog thread  
	- updates the user in the database  
	- unlocks the mutex  
  
  
  
# main.go  
package: main  
imports: context, log, os, strconv, github.com/JackBekket/hellper/lib/bot/command, github.com/JackBekket/hellper/lib/bot/dialog, github.com/JackBekket/hellper/lib/bot/env, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5, github.com/joho/godotenv  
  
cmdline arguments and flags:  
- OPENAI_API_KEY  
- PG_LINK  
- TG_KEY  
- ADMIN_ID  
- AI_ENDPOINT  
  
func main():  
	- loads environment variables using godotenv.Load()  
	- retrieves OPENAI_API_KEY from environment variable  
	- retrieves TG_KEY from environment variable  
	- retrieves ADMIN_ID from environment variable and parses it as an integer  
	- retrieves AI_ENDPOINT from environment variable  
	- creates a new instance of the Telegram bot API using the retrieved TG_KEY  
	- creates a map of admin data with ADMIN_ID as the key and env.AdminData as the value  
	- initializes the database and commander  
	- logs the authorized account  
	- creates a new update object with a timeout of 60 seconds  
	- creates a channel for handling updates  
	- starts a goroutine to handle updates using the dialog.HandleUpdates function  
	- iterates over incoming updates and checks if the user ID is in the database  
	- if the user ID is not in the database, it is added to the database and the update is sent to the update channel  
	- if the user ID is already in the database, the update is sent to the update channel  
  
# lib/bot/command/msgTemplates.go  
package: command  
imports:  
  
func main():  
	- prints "Hello world!"  
  
  
# lib/bot/dialog/dialog.go  
package: dialog  
imports:  
- log  
- github.com/JackBekket/hellper/lib/bot/command  
- github.com/JackBekket/hellper/lib/database  
- github.com/JackBekket/hellper/lib/langchain  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
external data:  
- usersDatabase: a map of users, where the key is the user's ID and the value is a struct containing user information.  
  
input sources:  
- updates: a channel of Telegram updates.  
  
cmdline arguments and flags:  
- None  
  
func HandleUpdates(updates <-chan tgbotapi.Update, bot *tgbotapi.BotAPI, comm command.Commander):  
	- Iterates over the updates channel.  
	- For each update:  
		- Extracts the chat ID from the update.  
		- Retrieves the user from the usersDatabase based on the chat ID.  
		- If the user is not found, adds a new user to the database.  
		- If the user is found, checks the command and handles it accordingly:  
			- /image: Generates an image based on the provided prompt or a default prompt.  
			- /restart: Restarts the user's session.  
			- /help: Displays help information.  
			- /search_doc: Searches for documents based on the provided prompt.  
			- /rag: Performs RAG (Retrieval Augmented Generation) based on the provided prompt.  
			- /instruct: Calls a local AI model to generate content based on the provided prompt.  
			- /usage: Displays usage information.  
			- /helper: Sends a media helper message.  
		- If the update message is not nil, it checks the user's dialog status and handles it accordingly:  
			- Status 0: Chooses a network.  
			- Status 1: Handles network choice.  
			- Status 2: Prompts the user to input their API key.  
			- Status 3: Chooses a model.  
			- Status 4: Handles model choice.  
			- Status 5: Connects to the AI with the chosen language.  
			- Status 6: Performs the dialog sequence.  
  
  
  
# lib/embeddings/query.go  
package: embeddings  
imports: context, fmt, github.com/tmc/langchaingo/chains, github.com/tmc/langchaingo/llms/openai, github.com/tmc/langchaingo/schema, github.com/tmc/langchaingo/vectorstores  
  
cmdline arguments and flags:  
- ai_url: URL of the AI service  
- api_token: API token for the AI service  
- question: The question to be answered  
- numOfResults: Number of results to return  
- store: Vector store to use for retrieval  
  
func Rag(ai_url string,api_token string,question string, numOfResults int,store vectorstores.VectorStore) (result string,err error):  
	- Sets base_url to ai_url  
	- Creates an embeddings client using the provided parameters  
	- Runs a retrieval QA chain using the embeddings client and the provided store  
	- Prints the final answer  
	- Returns the result and any error  
  
func SemanticSearch(searchQuery string, maxResults int, store vectorstores.VectorStore, options ...vectorstores.Option) (searchResults []schema.Document, err error):  
	- Performs a similarity search using the provided store and options  
	- Prints the similarity search results  
	- Returns the search results and any error  
  
  
  
# lib/embeddings/common.go  
package: embeddings  
imports: context, fmt, log, github.com/tmc/langchaingo/embeddings, github.com/tmc/langchaingo/llms/openai, github.com/tmc/langchaingo/vectorstores, github.com/tmc/langchaingo/vectorstores/pgvector, github.com/jackc/pgx/v5/pgxpool  
  
func LoadEnv():  
	- This function is not implemented.  
  
func GetVectorStore(ai_url string, api_token string, db_link string) (vectorstores.VectorStore, error):  
	- Takes ai_url, api_token, and db_link as input.  
	- Sets base_url to ai_url.  
	- Creates a connection to the database using the provided db_link.  
	- Creates an embeddings client using the OpenAI API with the provided api_token.  
	- Creates an embeddings client using the OpenAI API with the provided api_token.  
	- Creates a vector store using the pgvector library, connecting to the database and using the embeddings client.  
	- Returns the vector store and any errors encountered.  
  
func GetVectorStoreWithOptions(ai_url string, api_token string, db_link string, name string) (vectorstores.VectorStore, error):  
	- Takes ai_url, api_token, db_link, and name as input.  
	- Sets base_url to ai_url.  
	- Creates a connection to the database using the provided db_link.  
	- Creates an embeddings client using the OpenAI API with the provided api_token.  
	- Creates a vector store using the pgvector library, connecting to the database, using the embeddings client, and specifying the collection name as name.  
	- Returns the vector store and any errors encountered.  
  
  
  
# lib/localai/localai.go  
package: localai  
imports: bytes, encoding/json, fmt, io/ioutil, log, net/http, os, path/filepath, github.com/StarkBotsIndustries/telegraph  
  
external data:  
- http://localhost:8080/v1/chat/completions  
- http://localhost:8080/v1/images/generations  
  
cmdline arguments and flags:  
- prompt  
- modelName  
- url  
- s_pwd  
- u_pwd  
- size  
  
func main():  
	- prompt := "How are you?"  
	- modelName := "wizard-uncensored-13b"  
	- url := "http://localhost:8080/v1/chat/completions"  
	- calls GenerateCompletion(prompt, modelName, url)  
	- prints Assistant's response: (response from GenerateCompletion)  
  
func GenerateCompletion(prompt, modelName string, url string) (*ChatResponse, error):  
	- creates ChatRequest with modelName, prompt, and temperature  
	- converts ChatRequest to JSON  
	- sends POST request to url with JSON data  
	- reads response body  
	- parses JSON response into ChatResponse  
	- returns ChatResponse  
  
func GenerateCompletionWithPWD(prompt, modelName string, url string, s_pwd string, u_pwd string) (*ChatResponse, error):  
	- checks if passwords match  
	- calls GenerateCompletion(prompt, modelName, url)  
	- returns result  
  
func GenerateImageStableDissusion(prompt, size string) (string, error):  
	- creates payload with prompt and size  
	- converts payload to JSON  
	- sends POST request to url with JSON data  
	- reads response body  
	- parses JSON response into GenerationResponse  
	- returns image_url.URL  
  
func UploadToTelegraph(fileName string):  
	- gets absolute path to file  
	- opens file  
	- uploads file to telegraph  
	- returns link  
  
func deleteFromTemp(fileName string):  
	- gets absolute path to file  
	- deletes file from local machine  
  
  
  
# lib/localai/setupSequenceWithKey.go  
package: localai  
imports: context, log, sync, db, tgbotapi  
  
func SetupSequenceWithKey(  
	bot *tgbotapi.BotAPI,  
	user db.User,  
	language string,  
	ctx context.Context,  
	spwd string,  
	ai_endpoint string,  
) {  
	mu.Lock()  
	defer mu.Unlock()  
	chatID := user.ID  
	gptKey := user.AiSession.GptKey  
	log.Println("user GPT key from session: ", gptKey)  
	//u_network := user.Network  
	//log.Println("user network from session: ", u_network)  
	log.Println("user model from session: ", user.AiSession.GptModel)  
	//var client *openai.Client  
	u_pwd := user.AiSession.GptKey  
	log.Println("upwd: ", u_pwd)  
  
  
  
  
  
	//client := CreateClient(gptKey) // creating client (but we don't know if it works)  
	//log.Println("Setting up sequence with key")  
	//client := CreateLocalhostClientWithCheck(local_ap,gptKey)  
	//log.Println("local_ap: ", spwd)  
	//log.Println("client: ", client)  
	//log.Println("client: ", client.config)  
	//user.AiSession.GptClient = *client  
  
	  
  
	switch language {  
	case "English":  
		probe, err := tryLanguage(user, "", 1, ctx,ai_endpoint,spwd,u_pwd)  
		if err != nil {  
			errorMessage(err, bot, user)  
		} else {  
			msg := tgbotapi.NewMessage(chatID, probe)  
			bot.Send(msg)  
			user.DialogStatus = 4  
			db.UsersMap[chatID] = user  
		}  
	case "Russian":  
		probe, err := tryLanguage(user, "", 2, ctx,ai_endpoint,spwd,u_pwd)  
		if err != nil {  
			errorMessage(err, bot, user)  
		} else {  
			msg := tgbotapi.NewMessage(chatID, probe)  
			bot.Send(msg)  
			user.DialogStatus = 4  
			db.UsersMap[chatID] = user  
		}  
	default:  
		probe, err := tryLanguage(user, language, 0, ctx,ai_endpoint,spwd,u_pwd)  
		if err != nil {  
			errorMessage(err, bot, user)  
		} else {  
			msg := tgbotapi.NewMessage(chatID, probe)  
			bot.Send(msg)  
			user.DialogStatus = 4  
			db.UsersMap[chatID] = user  
		}  
	}  
}  
  
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context, ai_endpoint string, spwd string, upwd string) (string, error) {  
	var languagePromt string  
  
	switch languageCode {  
	case 1:  
		languagePromt = "Hi, Do you speak english?"  
	case 2:  
		languagePromt = "Привет, ты говоришь по-русски?"  
	default:  
		languagePromt = language  
	}  
  
	log.Printf("Language: %v\n", languagePromt)  
	model := user.AiSession.GptModel  
	//client := user.AiSession.GptClient  
	//log.Println("client: ", client)  
  
	/*  
	req := createComplexChatRequest(languagePromt, model)  
	log.Printf("request: %v\n", req)  
	*/  
  
	resp, err := GenerateCompletionWithPWD(languagePromt,model,ai_endpoint,spwd,upwd)  
	if err != nil {  
		return "", err  
	} else {  
		LogResponse(resp)  
		answer := resp.Choices[0].Message.Content  
		return answer, nil  
	}  
}  
  
  
  
# lib/localai/startDialogSequence.go  
package: localai  
imports: context, log, db, tgbotapi  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User):  
	- logs the error  
	- sends an error message to the user  
	- sends a message instructing the user to recreate the client and initialize a new session  
	- removes the user from the database  
  
func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string):  
	- locks a mutex  
	- retrieves the user from the database  
	- logs the GPT model and prompt  
	- generates a completion using the prompt and GPT model  
	- if there is an error, calls errorMessage  
	- if there is no error, logs the response, sends the response to the user, and updates the user's dialog status  
	- unlocks the mutex  
  
func LogResponse(resp *ChatResponse):  
	- logs various details about the response, including the full response object, created timestamp, ID, model, object, Choices[0], and usage statistics  
  
  
  
# lib/bot/command/addNewUsertoMap.go  
package: command  
imports: log, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func (c *Commander) AddNewUserToMap(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- user := database.User{  
		ID:           chatID,  
		Username:     updateMessage.From.UserName,  
		DialogStatus: 0,  
		Admin:        false,  
	}  
	- database.AddUser(user)  
	- log.Printf(  
		"Add new user to database: id: %v, username: %s\n",  
		user.ID,  
		user.Username,  
	)  
	- msg := tgbotapi.NewMessage(user.ID, msgTemplates["hello"])  
	- msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(  
		tgbotapi.NewKeyboardButtonRow(  
			tgbotapi.NewKeyboardButton("Start!")),  
	)  
	- c.bot.Send(msg)  
	- // check for registration  
	/*  
		if registred {  
			c.usersDb[chatID] = db.User{updateMessage.Chat.ID, updateMessage.Chat.UserName, 1}  
		}  
	*/  
  
  
  
# lib/bot/command/checkAdmin.go  
package: command  
imports: fmt, github.com/JackBekket/hellper/lib/bot/env, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func CheckAdmin(adminData map[string]env.AdminData, updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- iterates through adminData  
	- if admin.ID == chatID:  
		- if admin.GPTKey != "":  
			- c.AddAdminToMap(admin.GPTKey, updateMessage)  
			- return  
		- else:  
			- msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("env \"%s\" is missing.", evn))  
			- c.bot.Send(msg)  
			- c.AddNewUserToMap(updateMessage)  
			- return  
	- c.AddNewUserToMap(updateMessage)  
  
  
# lib/bot/command/newCommander.go  
package: command  
imports: context, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
type Commander:  
	- bot: *tgbotapi.BotAPI  
	- usersDb: map[int64]database.User  
	- ctx: context.Context  
  
func NewCommander(  
	bot *tgbotapi.BotAPI,  
	usersDb map[int64]database.User,  
	ctx context.Context,  
) *Commander:  
	- returns a new Commander instance with the given bot, usersDb, and ctx  
  
//func GetCommander():  
	- not implemented yet  
  
# lib/database/newAiSessionDataBase.go  
package: database  
imports: gogpt  
  
external data:  
- AiSessionMap: map[int64]AiSession  
  
func main():  
	- prints "Hello world!"  
  
  
# lib/embeddings/load.go  
package: embeddings  
imports: context, fmt, net/http, log, github.com/tmc/langchaingo/documentloaders, github.com/tmc/langchaingo/schema, github.com/tmc/langchaingo/textsplitter, github.com/tmc/langchaingo/vectorstores  
  
func LoadDocsToStore(docs []schema.Document, store vectorstores.VectorStore):  
	- prints "loading data from"  
	- prints "no. of documents to be loaded"  
	- adds documents to the vector store  
	- prints "data successfully loaded into vector store"  
	- logs any errors  
  
func getDocs(source string):  
	- fetches data from the given source  
	- loads and splits the data into documents  
	- returns the documents and any errors  
  
  
  
# lib/bot/command/addAdminTomap.go  
package: command  
imports: log, db, tgbotapi  
  
func (c *Commander) AddAdminToMap(adminKey string, updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- db.UsersMap[chatID] = db.User{  
		ID:           chatID,  
		Username:     updateMessage.From.UserName,  
		DialogStatus: 2,  
		Admin:        true,  
		AiSession: db.AiSession{  
			GptKey: adminKey,  
		},  
	}  
	- admin := db.UsersMap[chatID]  
	- log.Printf("%s authorized\n", admin.Username)  
	- msg := tgbotapi.NewMessage(admin.ID, "authorized: "+admin.Username)  
	- c.bot.Send(msg)  
	- msg = tgbotapi.NewMessage(admin.ID, msgTemplates["case1"])  
	- msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(  
		tgbotapi.NewKeyboardButtonRow(  
			tgbotapi.NewKeyboardButton("GPT-3.5")),  
		//tgbotapi.NewKeyboardButton("GPT-4"),  
		//tgbotapi.NewKeyboardButton("Codex")),  
	)  
	- c.bot.Send(msg)  
  
  
# lib/bot/command/utils.go  
package: command  
imports: fmt, log, math/rand, os, path/filepath, github.com/JackBekket/hellper/lib/database, github.com/JackBekket/hellper/lib/embeddings, github.com/go-telegram-bot-api/telegram-bot-api/v5, github.com/joho/godotenv  
  
func (c *Commander) HelpCommandMessage(updateMessage *tgbotapi.Message):  
	- gets chatID from updateMessage  
	- gets user from db.UsersMap using chatID  
	- creates a new message using msgTemplates["help_command"]  
	- sends the message to the user using c.bot.Send  
  
func (c *Commander) SearchDocuments(chatID int64, promt string, maxResults int):  
	- loads environment variables using godotenv.Load()  
	- gets PG_LINK and AI_BASEURL from environment variables  
	- gets user from db.UsersMap using chatID  
	- gets api_token from user.AiSession.GptKey  
	- gets vector store using embeddings.GetVectorStore  
	- performs semantic search using embeddings.SemanticSearch  
	- sends results to the user using c.bot.Send  
  
func (c *Commander) RAG(chatID int64, promt string, maxResults int):  
	- gets user from db.UsersMap using chatID  
	- loads environment variables using godotenv.Load()  
	- gets PG_LINK and AI_BASEURL from environment variables  
	- gets api_token from user.AiSession.GptKey  
	- gets vector store using embeddings.GetVectorStore  
	- performs RAG using embeddings.Rag  
	- sends result to the user using c.bot.Send  
  
func (c *Commander) GetUsage(chatID int64):  
	- gets user from db.UsersMap using chatID  
	- gets promt_tokens, completion_tokens, and total_tokens from user.AiSession.Usage  
	- sends usage information to the user using c.bot.Send  
  
func (c *Commander) SendMediaHelper(chatID int64):  
	- gets a list of files in the media directory  
	- selects a random file  
	- opens the video file  
	- creates a new video message  
	- sends the video message to the user using c.bot.Send  
  
  
  
# lib/bot/env/newEvn.go  
package: env  
imports: errors, log, strconv, github.com/joho/godotenv  
  
const envLoc = ".env"  
  
var env map[string]string  
  
type AdminData struct {  
	ID     int64  
	GPTKey string  
}  
  
func Load() error:  
	- reads environment variables from .env file  
	- returns error if reading fails  
  
func LoadAdminData() map[string]AdminData:  
	- creates a map to store admin data  
	- iterates through environment variables  
	- parses admin ID and GPT key for each admin  
	- returns a map with admin data  
  
func LoadTGToken() (string, error):  
	- returns telegram token from .env file  
	- returns error if token not found  
  
func LoadLocalPD() (string):  
	- returns local password from .env file  
  
func LoadLocalAI_Endpoint() (string):  
	- returns local AI endpoint from .env file  
  
func GetAdminToken() (string):  
	- returns admin key from .env file  
  
  
  
# lib/langchain/handler.go  
package: langchain  
imports: context, encoding/json, log, github.com/JackBekket/hellper/lib/database, github.com/tmc/langchaingo/llms, github.com/tmc/langchaingo/schema  
  
type ChainCallbackHandler struct{}  
  
func (h *ChainCallbackHandler) HandleAgentAction(ctx context.Context, action schema.AgentAction) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleAgentFinish(ctx context.Context, finish schema.AgentFinish) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleChainError(ctx context.Context, err error) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleLLMError(ctx context.Context, err error) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleLLMStart(ctx context.Context, prompts []string) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleRetrieverStart(ctx context.Context, query string) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleStreamingFunc(ctx context.Context, chunk []byte) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleToolEnd(ctx context.Context, output string) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleToolError(ctx context.Context, err error) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleToolStart(ctx context.Context, input string) {  
	//panic("unimplemented")  
}  
  
func (h *ChainCallbackHandler) HandleText(ctx context.Context, text string) {  
	// Implement this method if needed  
}  
  
func (h *ChainCallbackHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {  
	LogResponseContentChoice(ctx,res)	  
}  
  
func LogResponseContentChoice(ctx context.Context,resp *llms.ContentResponse) {  
	choice := resp.Choices[0]  
	log.Println("Content: ", choice.Content)  
	log.Println("Stop Reason: ", choice.StopReason)  
  
	log.Println("Context: ", ctx)  
	  
	user, ok := ctx.Value("user").(db.User)  
	if !ok {  
	  log.Println("No user in context")  
	  return  
	}  
	//chatID := user.ID  
  
  
  
	genInfo, err := json.Marshal(choice.GenerationInfo)  
	if err != nil {  
		log.Println("Error marshaling GenerationInfo: ", err)  
		return  
	}  
	log.Println("Generation Info: ", string(genInfo))  
  
	promt_tokens_str := choice.GenerationInfo["PromptTokens"]  
	completion_tokens_str := choice.GenerationInfo["CompletionTokens"]  
	total_tokens_str := choice.GenerationInfo["TotalTokens"]  
  
	pt, ok := promt_tokens_str.(int)  
	if !ok {  
  	log.Println("Error: value is not a string")  
  	return  
	}  
	ct, ok := completion_tokens_str.(int)  
	tt, ok := total_tokens_str.(int)  
	  
		  usage := map[string]int{  
			"Total": tt,  
			"Promt": pt,  
			"Completion": ct,  
		  }  
		  
		  db.UpdateSessionUsage(user.ID,usage)  
  
  
  
	if choice.FuncCall != nil {  
		log.Printf("Function Call: %+v\n", choice.FuncCall)  
	} else {  
		log.Println("No Function Call requested.")  
	}  
}  
  
  
  
# lib/langchain/langchain.go  
package: langchain  
imports: context, fmt, log, github.com/JackBekket/hellper/lib/database, github.com/tmc/langchaingo/chains, github.com/tmc/langchaingo/llms, github.com/tmc/langchaingo/memory, github.com/tmc/langchaingo/llms/openai  
  
func InitializeNewChatWithContextNoLimit(api_token string, model_name string, base_url string, user_initial_promt string) (*db.ChatSession, error):  
	- creates a new conversation using the provided API token, model name, and base URL.  
	- initializes a new conversation buffer and a new conversation object.  
	- returns a new chat session object with the conversation buffer and dialog thread.  
  
func StartNewChat(ctx context.Context, api_token string, model_name string, base_url string, user_initial_promt string) (string, *db.ChatSession, error):  
	- initializes a new chat session using the provided parameters.  
	- runs the chain using the initialized chat session and user prompt.  
	- returns the result, post-session, and any errors encountered.  
  
func RunChain(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error):  
	- runs the chain using the provided chat session and prompt.  
	- returns the result and the updated chat session.  
  
func ContinueChatWithContextNoLimit(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error):  
	- continues the dialog using the provided chat session and prompt.  
	- returns the result and the updated chat session.  
  
func GenerateContentInstruction(base_url string, promt string, model_name string, api_token string, network string, options ...llms.CallOption) (string, error):  
	- generates content based on the provided prompt, model name, API token, and network.  
	- returns the generated content and any errors encountered.  
  
  
  
# lib/bot/command/cases.go  
package: command  
  
imports:  
- context  
- log  
- net/url  
- os  
- path  
- github.com/JackBekket/hellper/lib/database  
- github.com/JackBekket/hellper/lib/langchain  
- github.com/JackBekket/hellper/lib/localai  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
- github.com/joho/godotenv  
  
type contextKey string  
const UserKey contextKey = "user"  
  
func (c *Commander) InputYourAPIKey(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- user := db.UsersMap[chatID]  
	- msg := tgbotapi.NewMessage(user.ID, msgTemplates["case0"])  
	- c.bot.Send(msg)  
	- user.DialogStatus = 3  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) ChooseNetwork(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- user := db.UsersMap[chatID]  
	- c.HelpCommandMessage(updateMessage)  
	- msg := tgbotapi.NewMessage(user.ID, msgTemplates["ch_network"])  
	- msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(  
		tgbotapi.NewKeyboardButtonRow(  
			tgbotapi.NewKeyboardButton("openai"),  
			tgbotapi.NewKeyboardButton("localai"),  
			tgbotapi.NewKeyboardButton("vastai")),  
  
	)  
	- c.bot.Send(msg)  
	- user.DialogStatus = 1  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) HandleNetworkChoose(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- network := updateMessage.Text  
	- user := db.UsersMap[chatID]  
	- switch network {  
	case "openai":  
		- user.Network = network  
		- user.AiSession.AI_Type = 0  
		- user.DialogStatus = 2  
		- db.UsersMap[chatID] = user  
		- c.InputYourAPIKey(updateMessage)  
	case "localai":  
		- user.Network = network  
		- user.AiSession.AI_Type = 1  
		- user.DialogStatus = 2  
		- db.UsersMap[chatID] = user  
		- c.InputYourAPIKey(updateMessage)  
	case "vastai":  
		- user.Network = network  
		- user.AiSession.AI_Type = 2  
		- user.DialogStatus = 2  
		- db.UsersMap[chatID] = user  
		- c.InputYourAPIKey(updateMessage)  
	default:  
		- c.WrongNetwork(updateMessage)  
	}  
  
func (c *Commander) ChooseModel(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- gptKey := updateMessage.Text  
	- user := db.UsersMap[chatID]  
	- network := user.Network  
	- user.AiSession.GptKey = gptKey  
	- switch network {  
	case "localai":  
		- c.RenderModelMenuLAI(chatID)  
		- user.DialogStatus = 4  
		- db.UsersMap[chatID] = user  
	case "openai":  
		- c.RenderModelMenuOAI(chatID)  
		- user.DialogStatus = 4  
		- db.UsersMap[chatID] = user  
	case "vastai":  
		- c.RenderModelMenuVAI(chatID)  
		- user.DialogStatus = 4  
		- db.UsersMap[chatID] = user  
	default:  
		- c.WrongNetwork(updateMessage)  
	}  
  
func (c *Commander) HandleModelChoose(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- model_name := updateMessage.Text  
	- user := db.UsersMap[chatID]  
	- network := user.Network  
	- switch network {  
	case "localai":  
		- switch model_name {  
		case "wizard-uncensored-13b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case "wizard-uncensored-30b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "deepseek-coder-6b-instruct":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "tiger-gemma-9b-v1-i1":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "wizard-uncensored-code-34b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		default:  
			- c.WrongModel(updateMessage)  
		}  
	case "openai":  
		- switch model_name {  
		case "gpt-3.5":  
			- model_name = "gpt-3.5-turbo"  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case "gpt-4":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		default:  
			- c.WrongModel(updateMessage)  
		}  
	case "vastai":  
		- switch model_name {  
		case "wizard-uncensored-13b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case "wizard-uncensored-30b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "deepseek-coder-6b-instruct":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "tiger-gemma-9b-v1-i1":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "big-tiger-gemma-27b-v1":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		case  "wizard-uncensored-code-34b":  
			- c.attachModel(model_name, chatID)  
			- user.AiSession.GptModel = model_name  
			- c.RenderLanguage(chatID)  
			- user.DialogStatus = 5  
			- db.UsersMap[chatID] = user  
		default:  
			- c.WrongModel(updateMessage)  
		}  
	}  
  
func (c *Commander) attachModel(model_name string, chatID int64):  
	- user := db.UsersMap[chatID]  
	- modelName := model_name  
	- user.AiSession.GptModel = modelName  
	- msg := tgbotapi.NewMessage(user.ID, "your session model: "+modelName)  
	- c.bot.Send(msg)  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) AttachKey(gpt_key string, chatID int64):  
	- log.Println("Key promt: ", gpt_key)  
	- user := db.UsersMap[chatID]  
	- user.AiSession.GptKey = gpt_key  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) ChangeDialogStatus(chatID int64, ds int8):  
	- user := db.UsersMap[chatID]  
	- old_status := user.DialogStatus  
	- log.Println("dialog status changed, old status is ", old_status)  
	- log.Println("new status is ", ds)  
	- user.DialogStatus = ds  
  
func (c *Commander) WrongModel(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- user := db.UsersMap[chatID]  
	- msg := tgbotapi.NewMessage(user.ID, "type wizard-uncensored-13b")  
	- c.bot.Send(msg)  
	- user.DialogStatus = 4  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) WrongNetwork(updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- user := db.UsersMap[chatID]  
	- msg := tgbotapi.NewMessage(user.ID, "type openai or localai")  
	- c.bot.Send(msg)  
	- user.DialogStatus = 0  
	- db.UsersMap[chatID] = user  
  
func (c *Commander) ConnectingToAiWithLanguage(updateMessage *tgbotapi.Message, ai_endpoint string):  
	- chatID := updateMessage.From.ID  
	- language := updateMessage.Text  
	- user := db.UsersMap[chatID]  
	- log.Println("check gpt key exist:", user.AiSession.GptKey)  
	- network := user.Network  
	- msg := tgbotapi.NewMessage(user.ID, "connecting to ai node")  
	- c.bot.Send(msg)  
	- ctx := context.WithValue(c.ctx, "user", user)  
	- if network == "localai" {  
		- log.Println("network: ", network)  
		- if ai_endpoint == "" {  
			- ai_endpoint = os.Getenv("AI_ENDPOINT")  
		}  
		- log.Println("local-ai endpoint is: ", ai_endpoint)  
		- go langchain.SetupSequenceWithKey(c.bot,user,language,ctx,ai_endpoint)	//local-ai  
	} else if network ==  "vastai" {  
		- log.Println("network: ", network)  
		- ai_endpoint := os.Getenv("VASTAI_ENDPOINT")  
		- log.Println("vast-ai endpoint is: ", ai_endpoint)  
		- go langchain.SetupSequenceWithKey(c.bot,user,language,ctx,ai_endpoint)	//vast-ai  
	} else {  
		- log.Println("network: ", network)  
		- go langchain.SetupSequenceWithKey(c.bot,user,language,ctx,"")	//openai  
	}  
  
func (c *Commander) DialogSequence(updateMessage *tgbotapi.Message, ai_endpoint string):  
	- chatID := updateMessage.From.ID  
	- user := db.UsersMap[chatID]  
	- if updateMessage != nil {  
		- promt := updateMessage.Text  
		- ctx := context.WithValue(c.ctx, "user", user)  
		- go langchain.StartDialogSequence(c.bot,chatID,promt,ctx,ai_endpoint)  
	}  
  
func (c *Commander) GenerateNewImageLAI_SD(promt string, chatID int64, bot *tgbotapi.BotAPI):  
	- size := "256x256"  
	- filepath, err := localai.GenerateImageStableDissusion(promt, size)  
	- if err != nil {  
		- log.Println(err)  
	}  
	- log.Println("url_path: ", filepath)  
	- sendImage(bot, chatID, filepath)  
  
func sendImage(bot *tgbotapi.BotAPI, chatID int64, path string):  
	- // Prepare a photo message  
	- fileName := transformURL(path)  
	- log.Println("local file name: ", fileName)  
	- telegraphLink := localai.UploadToTelegraph(fileName)  
	- log.Println("uploaded to telegraph successfully, link is: ", telegraphLink)  
	- msg := tgbotapi.NewMessage(chatID, telegraphLink)  
	- bot.Send(msg)  
  
func transformURL(inputURL string) string {  
	- // Replace "http://localhost:8080" with "/tmp" using strings.Replace  
	- parsedURL, _ := url.Parse(inputURL)  
	- // Use path.Base to get the filename from the URL path  
	- fileName := path.Base(parsedURL.Path)  
	- return fileName  
  
func (c *Commander) GetUsersDb() (map[int64]db.User):  
	- data_base := db.UsersMap  
	- return data_base  
  
func (c *Commander) GetUser(id int64) (db.User):  
	- user := db.UsersMap[id]  
	- return user  
  
# lib/bot/command/ui.go  
package: command  
imports: tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"  
  
func (c *Commander) RenderModelMenuOAI(chatID int64):  
	- creates a new message with the text from msgTemplates["case1"] and sends it to the chat with the given chatID  
	- sets the reply markup to a one-time reply keyboard with a single row containing two buttons: "gpt-3.5" and "gpt-4"  
	- sends the message using the bot  
  
func (c *Commander) RenderModelMenuLAI(chatID int64):  
	- creates a new message with the text from msgTemplates["case1"] and sends it to the chat with the given chatID  
	- sets the reply markup to a one-time reply keyboard with a single row containing three buttons: "wizard-uncensored-13b", "wizard-uncensored-30b", and "tiger-gemma-9b-v1-i1"  
	- sends the message using the bot  
  
func (c *Commander) RenderModelMenuVAI(chatID int64):  
	- creates a new message with the text from msgTemplates["case1"] and sends it to the chat with the given chatID  
	- sets the reply markup to a one-time reply keyboard with two rows:  
		- the first row contains two buttons: "deepseek-coder-6b-instruct" and "wizard-uncensored-code-34b"  
		- the second row contains two buttons: "tiger-gemma-9b-v1-i1" and "big-tiger-gemma-27b-v1"  
	- sends the message using the bot  
  
func (c *Commander) RenderLanguage(chat_id int64):  
	- creates a new message with the text "Choose a language or send 'Hello' in your desired language." and sends it to the chat with the given chatID  
	- sets the reply markup to a one-time reply keyboard with a single row containing two buttons: "English" and "Russian"  
	- sends the message using the bot  
  
  
  
# lib/database/newUserDataBase.go  
package: database  
imports: github.com/tmc/langchaingo/chains, github.com/tmc/langchaingo/memory  
  
func AddUser(user User):  
	- adds a new user to the UsersMap  
func UpdateUserUsage(id int64, usage map[string]int):  
	- updates the usage of a user's AiSession  
func UpdateSessionUsage(id int64, usage map[string]int):  
	- updates the usage of a session  
func GetSessionUsage(id int64) (map[string]int):  
	- returns the usage of a session  
  
User:  
	- ID: int64  
	- Username: string  
	- DialogStatus: int8  
	- Admin: bool  
	- AiSession: AiSession  
	- Network: string  
  
SessionUsage:  
	- ID: int64  
	- Usage: map[string]int  
  
AiSession:  
	- GptKey: string  
	- GptModel: string  
	- AI_Type: int8  
	- DialogThread: ChatSession  
	- Base_url: string  
	- Usage: map[string]int  
  
ChatSession:  
	- ConversationBuffer: memory.ConversationBuffer  
	- DialogThread: chains.LLMChain  
  
UsersMap:  
	- map[int64]User  
  
UsageMap:  
	- map[int64]SessionUsage  
  
  
  
