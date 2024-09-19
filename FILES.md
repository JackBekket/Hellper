# lib/bot/command/utils.go  
package: command  
imports: fmt, log, math/rand, os, path/filepath, github.com/JackBekket/hellper/lib/database, github.com/JackBekket/hellper/lib/embeddings, github.com/go-telegram-bot-api/telegram-bot-api/v5, github.com/joho/godotenv  
  
env: "PG_LINK", "AI_BASEURL"  
  
struct:  
	- Commander:  
		- HelpCommandMessage(updateMessage *tgbotapi.Message)  
		- SearchDocuments(chatID int64, promt string, maxResults int)  
		- RAG(chatID int64, promt string, maxResults int)  
		- GetUsage(chatID int64)  
		- SendMediaHelper(chatID int64)  
  
func HelpCommandMessage(updateMessage *tgbotapi.Message):  
	- gets chatID from updateMessage.From.ID  
	- gets user from db.UsersMap[chatID]  
	- creates a new message with msgTemplates["help_command"]  
	- sends the message using c.bot.Send  
  
func SearchDocuments(chatID int64, promt string, maxResults int):  
	- loads environment variables using godotenv.Load()  
	- gets PG_LINK and AI_BASEURL from environment variables  
	- gets db_conn from PG_LINK  
	- gets user from db.UsersMap[chatID]  
	- gets api_token from user.AiSession.GptKey  
	- gets vector store using embeddings.GetVectorStore  
	- performs semantic search using embeddings.SemanticSearch  
	- iterates through results and sends each result to the user  
  
func RAG(chatID int64, promt string, maxResults int):  
	- gets user from db.UsersMap[chatID]  
	- loads environment variables using godotenv.Load()  
	- gets PG_LINK and AI_BASEURL from environment variables  
	- gets db_conn from PG_LINK  
	- gets api_token from user.AiSession.GptKey  
	- gets vector store using embeddings.GetVectorStore  
	- performs RAG using embeddings.Rag  
	- sends the result to the user  
  
func GetUsage(chatID int64):  
	- gets user from db.UsersMap[chatID]  
	- gets promt_tokens, completion_tokens, and total_tokens from user.AiSession.Usage  
	- sends the usage information to the user  
  
func SendMediaHelper(chatID int64):  
	- gets a list of files in the media directory  
	- selects a random file  
	- opens the video file  
	- creates a new video message  
	- sends the video message to the user  
  
  
  
# lib/database/newAiSessionDataBase.go  
package: database  
imports: gogpt  
  
struct AiSession:  
	- Fields: GptKey, GptClient, GptModel  
  
var:  
	- AiSessionMap: map[int64]AiSession  
  
  
  
# lib/langchain/langchain.go  
package: langchain  
imports: context, fmt, log, github.com/JackBekket/hellper/lib/database, github.com/tmc/langchaingo/chains, github.com/tmc/langchaingo/llms, github.com/tmc/langchaingo/memory, github.com/tmc/langchaingo/llms/openai  
  
func InitializeNewChatWithContextNoLimit(api_token string, model_name string, base_url string, user_initial_promt string) (*db.ChatSession, error):  
	- creates new conversation with given parameters  
	- returns new chat session  
  
func StartNewChat(ctx context.Context, api_token string, model_name string, base_url string, user_initial_promt string) (string, *db.ChatSession, error):  
	- initializes new chat session  
	- runs chain with given prompt  
	- returns result, session, and error  
  
func RunChain(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error):  
	- runs chain with given prompt  
	- returns result, session, and error  
  
func ContinueChatWithContextNoLimit(ctx context.Context, session *db.ChatSession, prompt string) (string, *db.ChatSession, error):  
	- runs chain with given prompt  
	- returns result, session, and error  
  
func GenerateContentInstruction(base_url string, promt string, model_name string, api_token string, network string, options ...llms.CallOption) (string, error):  
	- generates content from single prompt  
	- returns result and error  
  
  
  
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
  
  
  
# lib/bot/command/cases.go  
package: command  
imports: context, log, net/url, os, path, github.com/JackBekket/hellper/lib/database, github.com/JackBekket/hellper/lib/langchain, github.com/JackBekket/hellper/lib/localai, github.com/go-telegram-bot-api/telegram-bot-api/v5, github.com/joho/godotenv  
  
struct:  
	- Commander:  
		- InputYourAPIKey:  
			- takes updateMessage as input  
			- sets DialogStatus to 3  
		- ChooseNetwork:  
			- takes updateMessage as input  
			- sets DialogStatus to 1  
		- HandleNetworkChoose:  
			- takes updateMessage as input  
			- sets DialogStatus to 2  
		- ChooseModel:  
			- takes updateMessage as input  
			- sets DialogStatus to 4  
		- HandleModelChoose:  
			- takes updateMessage as input  
			- sets DialogStatus to 5  
		- WrongModel:  
			- takes updateMessage as input  
			- sets DialogStatus to 4  
		- WrongNetwork:  
			- takes updateMessage as input  
			- sets DialogStatus to 0  
		- ConnectingToAiWithLanguage:  
			- takes updateMessage and ai_endpoint as input  
			- sets DialogStatus to 6  
		- DialogSequence:  
			- takes updateMessage and ai_endpoint as input  
			- sets DialogStatus to 6  
		- GenerateNewImageLAI_SD:  
			- takes promt, chatID, and bot as input  
		- sendImage:  
			- takes bot, chatID, and path as input  
		- transformURL:  
			- takes inputURL as input  
		- GetUsersDb:  
			- returns UsersMap  
		- GetUser:  
			- takes id as input  
			- returns User  
  
func:  
	- AttachKey:  
		- takes gpt_key and chatID as input  
		- sets user.AiSession.GptKey to gpt_key  
	- ChangeDialogStatus:  
		- takes chatID and ds as input  
		- sets user.DialogStatus to ds  
	- attachModel:  
		- takes model_name and chatID as input  
		- sets user.AiSession.GptModel to model_name  
	- RenderModelMenuLAI:  
		- takes chatID as input  
	- RenderModelMenuOAI:  
		- takes chatID as input  
	- RenderModelMenuVAI:  
		- takes chatID as input  
  
# lib/bot/command/msgTemplates.go  
package: command  
imports: none  
  
var msgTemplates:  
	- hello: "Hey, this bot is working with local ai node."  
	- case0: "If you choosed OpenAI then input your openai-key, if you choosed LocalAI then input password. DM me in telegram if you want password or buy key from openai"  
	- await: "Awaiting"  
	- case1: "Choose model to use. "  
	- ch_network: "choose network to work with. openai make calls to openai, localhost tries to access localai node installed alongside with bot"  
	- help_command: "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire"  
  
  
  
# lib/embeddings/query.go  
package: embeddings  
imports: context, fmt, chains, openai, schema, vectorstores  
  
func Rag(ai_url string,api_token string,question string, numOfResults int,store vectorstores.VectorStore) (result string,err error) {  
	base_url := ai_url  
	llm, err := openai.New(  
		openai.WithBaseURL(base_url),  
		openai.WithAPIVersion("v1"),  
		openai.WithToken(api_token),  
    	openai.WithModel("wizard-uncensored-13b"),  
    	openai.WithEmbeddingModel("text-embedding-ada-002"),  
	)  
	if err != nil {  
		return "",err  
	}  
	result, err = chains.Run(  
		context.Background(),  
		chains.NewRetrievalQAFromLLM(  
			llm,  
			vectorstores.ToRetriever(store, numOfResults),  
		),  
		question,  
		chains.WithMaxTokens(2048),  
	)  
	if err != nil {  
		return "",err  
	}  
	fmt.Println("====final answer====\n", result)  
	return result,nil  
  
}  
  
func SemanticSearch(searchQuery string, maxResults int, store vectorstores.VectorStore, options ...vectorstores.Option) (searchResults []schema.Document, err error) {  
	searchResults, err = store.SimilaritySearch(context.Background(), searchQuery, maxResults, options...)  
	if err != nil {  
		return nil,err  
	}  
	fmt.Println("============== similarity search results ==============")  
	for _, doc := range searchResults {  
		fmt.Println("similarity search info -", doc.PageContent)  
		fmt.Println("similarity search score -", doc.Score)  
		fmt.Println("============================")  
	}  
	return searchResults,nil  
  
}  
  
  
  
# lib/langchain/setupSequenceWithKey.go  
package: langchain  
imports: context, log, sync, db, tgbotapi  
  
struct db.User:  
	- Fields: ID, AiSession  
		- AiSession:  
			- Fields: GptKey, GptModel, DialogThread, Usage  
struct db.ChatSession:  
	- Fields: DialogThread, Usage  
  
func SetupSequenceWithKey(  
	bot *tgbotapi.BotAPI,  
	user db.User,  
	language string,  
	ctx context.Context,  
	ai_endpoint string,  
) {  
	- Acquires a lock on the mutex  
	- Retrieves the chat ID, GPT key, and GPT model from the user's AiSession  
	- Prints the GPT key and GPT model from the user's AiSession  
	- Determines the language based on the provided language string  
	- Calls the tryLanguage function to handle the language-specific logic  
	- Updates the user's DialogStatus and AiSession.DialogThread based on the language  
	- Retrieves the session usage and updates the user's AiSession.Usage  
	- Stores the updated user in the db.UsersMap  
  
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context, ai_endpoint string) (string,*db.ChatSession, error) {  
	- Initializes a language prompt based on the language code  
	- Retrieves the GPT key, GPT model, and chat ID from the user's AiSession  
	- Calls the StartNewChat function to initiate a new chat session  
	- Returns the chat result, thread, and any errors encountered  
  
func StartNewChat(ctx context.Context, gptKey string, model string, ai_endpoint string, languagePromt string) (string,*db.ChatSession, error) {  
	// ... (implementation not provided)  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User) {  
	// ... (implementation not provided)  
  
  
  
# lib/langchain/startDialogSequence.go  
package: langchain  
imports: context, io/ioutil, log, math/rand, os, path/filepath, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
env: "PWD", "BAZ"  
  
struct FooBar:  
	- Fields: Foo, Bar  
var:  
	- zab: "zab"  
const:  
	- zoob: "zoob"  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User):  
	- logs error  
	- sends error message to user  
	- sends message to user to recreate client and initialize new session  
	- gets a random video file from the media directory  
	- opens the video file  
	- creates a new video message  
	- sends the video message  
	- removes user from the database  
  
func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string):  
	- locks mutex  
	- gets user from database  
	- gets GPT model and prompt  
	- gets dialog thread  
	- continues chat with context and prompt  
	- sends AI response to user  
	- updates user's dialog status  
	- updates user's session usage  
	- gets chat history  
	- updates user's dialog thread  
	- updates user in database  
	- unlocks mutex  
  
  
  
# lib/localai/startDialogSequence.go  
package: localai  
imports: context, log, db, tgbotapi  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User):  
	- logs the error  
	- sends an error message to the user via the bot  
	- sends a message instructing the user to recreate the client and initialize a new session  
	- removes the user from the database  
  
func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string):  
	- acquires a lock  
	- retrieves the user from the database  
	- logs the GPT model and prompt  
	- generates a completion using the prompt and GPT model  
	- handles errors by calling errorMessage  
	- logs the response  
	- sends the response to the user via the bot  
	- updates the user's dialog status  
	- releases the lock  
  
func LogResponse(resp *ChatResponse):  
	- logs various details about the response, including the full response object, created timestamp, ID, model, object, Choices[0], and usage statistics  
  
  
  
# lib/bot/env/newEvn.go  
package: env  
imports: errors, log, strconv, godotenv  
  
const:  
	- envLoc: ".env"  
  
var:  
	- env: map[string]string  
  
type AdminData:  
	- Fields: ID, GPTKey  
  
func Load(): error  
	- reads environment variables from .env file  
	- returns error if reading fails  
  
func LoadAdminData(): map[string]AdminData  
	- creates a map to store admin data  
	- iterates through environment variables  
	- parses admin IDs and GPT keys  
	- returns a map with admin data  
  
func LoadTGToken() (string, error)  
	- returns telegram token from environment variable  
	- returns error if token not found  
  
func LoadLocalPD() (string)  
	- returns localhost password from environment variable  
  
func LoadLocalAI_Endpoint() (string)  
	- returns local AI endpoint from environment variable  
  
func GetAdminToken() (string)  
	- returns admin key from environment variable  
  
  
  
# lib/database/newUserDataBase.go  
package: database  
imports: github.com/tmc/langchaingo/chains, github.com/tmc/langchaingo/memory  
  
struct User:  
	- Fields: ID, Username, DialogStatus, Admin, AiSession, Network  
struct SessionUsage:  
	- Fields: ID, Usage  
struct AiSession:  
	- Fields: GptKey, GptModel, AI_Type, DialogThread, Base_url, Usage  
struct ChatSession:  
	- Fields: ConversationBuffer, DialogThread  
  
var:  
	- UsersMap: map[int64]User  
	- UsageMap: map[int64]SessionUsage  
  
func AddUser(user User):  
	- adds user to UsersMap  
  
func UpdateUserUsage(id int64, usage map[string]int):  
	- updates usage of user with given id in UsersMap  
  
func UpdateSessionUsage(id int64, usage map[string]int):  
	- updates usage of session with given id in UsageMap  
  
func GetSessionUsage(id int64):  
	- returns usage of session with given id from UsageMap  
  
  
  
# lib/embeddings/common.go  
package: embeddings  
imports: context, fmt, log, github.com/tmc/langchaingo/embeddings, github.com/tmc/langchaingo/llms/openai, github.com/tmc/langchaingo/vectorstores, github.com/tmc/langchaingo/vectorstores/pgvector, github.com/jackc/pgx/v5/pgxpool  
  
func LoadEnv(): unimplemented  
func GetVectorStore(ai_url string, api_token string, db_link string) (vectorstores.VectorStore, error):  
	- sets base_url to ai_url  
	- gets pgConnURL from db_link  
	- creates a pgxpool.Config from pgConnURL  
	- creates a pgxpool.Pool from the config  
	- creates an embeddings.Embedder using the OpenAI API with the given api_token and base_url  
	- creates a pgvector.VectorStore using the Embedder and the pgxpool.Pool  
	- returns the VectorStore and any error  
  
func GetVectorStoreWithOptions(ai_url string, api_token string, db_link string, name string) (vectorstores.VectorStore, error):  
	- sets base_url to ai_url  
	- gets pgConnURL from db_link  
	- creates a pgxpool.Config from pgConnURL  
	- creates a pgxpool.Pool from the config  
	- creates an embeddings.Embedder using the OpenAI API with the given api_token and base_url  
	- creates a pgvector.VectorStore using the Embedder, the pgxpool.Pool, and the given name for the collection  
	- returns the VectorStore and any error  
  
  
  
# lib/langchain/handler.go  
package: langchain  
imports: context, encoding/json, log, github.com/JackBekket/hellper/lib/database, github.com/tmc/langchaingo/llms, github.com/tmc/langchaingo/schema  
  
struct ChainCallbackHandler:  
	- HandleAgentAction: unimplemented  
	- HandleAgentFinish: unimplemented  
	- HandleChainEnd: unimplemented  
	- HandleChainError: unimplemented  
	- HandleChainStart: unimplemented  
	- HandleLLMError: unimplemented  
	- HandleLLMGenerateContentStart: unimplemented  
	- HandleLLMStart: unimplemented  
	- HandleRetrieverEnd: unimplemented  
	- HandleRetrieverStart: unimplemented  
	- HandleStreamingFunc: unimplemented  
	- HandleToolEnd: unimplemented  
	- HandleToolError: unimplemented  
	- HandleToolStart: unimplemented  
	- HandleText: unimplemented  
	- HandleLLMGenerateContentEnd:  
		- Logs response content choice  
		- Logs stop reason  
		- Logs context  
		- Gets user from context  
		- Logs generation info  
		- Logs prompt tokens  
		- Logs completion tokens  
		- Logs total tokens  
		- Updates user's usage information  
		- Saves user usage to database  
  
func LogResponseContentChoice(ctx context.Context,resp *llms.ContentResponse):  
	- Logs content, stop reason, context, generation info, prompt tokens, completion tokens, and total tokens.  
	- Updates user's usage information.  
	- Saves user usage to database.  
  
  
  
# lib/localai/setupSequenceWithKey.go  
package: localai  
imports: context, log, sync, db, tgbotapi  
env: "PWD", "BAZ"  
  
struct db.User:  
	- Fields: ID, AiSession  
		- AiSession:  
			- GptKey  
			- GptModel  
			- GptClient  
			- Network  
  
var:  
	- mu: sync.Mutex  
  
func SetupSequenceWithKey(  
	bot *tgbotapi.BotAPI,  
	user db.User,  
	language string,  
	ctx context.Context,  
	spwd string,  
	ai_endpoint string,  
) {  
	- acquires lock on mu  
	- assigns chatID to user.ID  
	- assigns gptKey to user.AiSession.GptKey  
	- logs user GPT key from session  
	- logs user model from session  
	- logs upwd  
	- checks if language is "English", "Russian" or default  
	- calls tryLanguage function based on language  
	- updates user.DialogStatus to 4  
	- updates db.UsersMap with user  
  
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context, ai_endpoint string, spwd string, upwd string) (string, error) {  
	- assigns languagePromt based on languageCode  
	- logs languagePromt  
	- assigns model to user.AiSession.GptModel  
	- calls GenerateCompletionWithPWD function  
	- logs response  
	- returns answer from response  
  
func GenerateCompletionWithPWD(languagePromt string, model string, ai_endpoint string, spwd string, upwd string) (resp, err) {  
	// TODO: implement GenerateCompletionWithPWD  
}  
  
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User) {  
	// TODO: implement errorMessage  
}  
  
  
  
# main.go  
package: main  
imports: context, log, os, strconv, github.com/JackBekket/hellper/lib/bot/command, github.com/JackBekket/hellper/lib/bot/dialog, github.com/JackBekket/hellper/lib/bot/env, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5, github.com/joho/godotenv  
  
env: OPENAI_API_KEY, TG_KEY, ADMIN_ID, AI_ENDPOINT  
  
func main():  
	- loads environment variables using godotenv.Load()  
	- retrieves OPENAI_API_KEY from environment variable OPENAI_API_KEY  
	- retrieves TG_KEY from environment variable TG_KEY  
	- retrieves ADMIN_ID from environment variable ADMIN_ID and parses it to int64  
	- retrieves AI_ENDPOINT from environment variable AI_ENDPOINT  
	- creates a new instance of tgbotapi.BotAPI using the retrieved TG_KEY  
	- creates a map of AdminData with ADMIN_ID as key and env.AdminData as value  
	- initializes database and commander  
	- prints a message indicating that the bot is authorized on the specified account  
	- creates a channel for handling updates  
	- starts a goroutine to handle updates using dialog.HandleUpdates  
	- iterates through incoming updates and checks if the user is new, if so, creates a new entry in the database  
	- sends the update to the upd_ch channel  
  
  
  
# lib/bot/command/ui.go  
package: command  
imports: tgbotapi  
  
func (c *Commander) RenderModelMenuOAI(chatID int64):  
	- creates a new message with chatID and msgTemplates["case1"]  
	- sets the reply markup to a one-time reply keyboard with a single row containing two buttons: "gpt-3.5" and "gpt-4"  
	- sends the message using c.bot.Send(msg)  
  
func (c *Commander) RenderModelMenuLAI(chatID int64):  
	- creates a new message with chatID and msgTemplates["case1"]  
	- sets the reply markup to a one-time reply keyboard with a single row containing three buttons: "wizard-uncensored-13b", "wizard-uncensored-30b", and "tiger-gemma-9b-v1-i1"  
	- sends the message using c.bot.Send(msg)  
  
func (c *Commander) RenderModelMenuVAI(chatID int64):  
	- creates a new message with chatID and msgTemplates["case1"]  
	- sets the reply markup to a one-time reply keyboard with two rows:  
		- first row contains two buttons: "deepseek-coder-6b-instruct" and "wizard-uncensored-code-34b"  
		- second row contains two buttons: "tiger-gemma-9b-v1-i1" and "big-tiger-gemma-27b-v1"  
	- sends the message using c.bot.Send(msg)  
  
func (c *Commander) RenderLanguage(chat_id int64):  
	- creates a new message with chatID and a message asking the user to choose a language or send "Hello" in their desired language  
	- sets the reply markup to a one-time reply keyboard with a single row containing two buttons: "English" and "Russian"  
	- sends the message using c.bot.Send(msg)  
  
  
  
# lib/bot/dialog/dialog.go  
package: dialog  
imports: log, github.com/JackBekket/hellper/lib/bot/command, github.com/JackBekket/hellper/lib/database, github.com/JackBekket/hellper/lib/langchain, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func HandleUpdates(updates <-chan tgbotapi.Update, bot *tgbotapi.BotAPI, comm command.Commander):  
	for update := range updates {  
		var chatID int64  
		chatID = int64(update.Message.From.ID)  
		db := comm.GetUsersDb()  
		user, ok := db[int64(chatID)]  
		if !ok {  
			comm.AddNewUserToMap(update.Message)  
		}  
		if ok {  
			switch update.Message.Command() {  
			case "image":  
				msg := tgbotapi.NewMessage(user.ID, "Image link generation...")  
				bot.Send(msg)  
				promt := update.Message.CommandArguments()  
				log.Printf("Command /image arg: %s\n", promt)  
				if (promt == "") {  
					comm.GenerateNewImageLAI_SD("evangelion, neon, anime",chatID,bot)  
				} else {  
					comm.GenerateNewImageLAI_SD(promt,chatID,bot)  
				}  
			case "restart":  
				msg := tgbotapi.NewMessage(user.ID, "Restarting session..., type any key")  
				bot.Send(msg)  
				userDb := database.UsersMap  
				delete(userDb, user.ID)  
			case "help":  
				comm.HelpCommandMessage(update.Message)  
			case "search_doc":  
				promt := update.Message.CommandArguments()  
				comm.SearchDocuments(chatID,promt,3)  
			case "rag":  
				promt := update.Message.CommandArguments()  
				comm.RAG(chatID,promt,1)  
			case "instruct" :  
				promt := update.Message.CommandArguments()  
				model_name := user.AiSession.GptModel  
				api_token := user.AiSession.GptKey  
				langchain.GenerateContentInstruction(user.AiSession.Base_url,promt,model_name,api_token,user.Network)  
			case "usage" :  
				comm.GetUsage(chatID)  
			case "helper":  
				comm.SendMediaHelper(chatID)  
			}  
		}  
	}  
}  
  
  
# lib/embeddings/load.go  
package: embeddings  
imports: context, fmt, net/http, log, github.com/tmc/langchaingo/documentloaders, github.com/tmc/langchaingo/schema, github.com/tmc/langchaingo/textsplitter, github.com/tmc/langchaingo/vectorstores  
  
func LoadDocsToStore(docs []schema.Document, store vectorstores.VectorStore):  
	- prints "loading data from"  
	- prints "no. of documents to be loaded" followed by the length of the docs slice  
	- calls store.AddDocuments with context.Background() and the docs slice  
	- prints "data successfully loaded into vector store"  
	- prints the error if any  
  
func getDocs(source string):  
	- sends a GET request to the given source  
	- closes the response body after use  
	- loads and splits the response body using documentloaders.NewHTML and textsplitter.NewRecursiveCharacter  
	- returns the loaded documents and any error  
  
  
  
# lib/localai/localai.go  
package: localai  
imports: bytes, encoding/json, fmt, io/ioutil, log, net/http, os, path/filepath, github.com/StarkBotsIndustries/telegraph  
  
structs:  
	- ChatRequest:  
		- Fields: Model, Messages, Temperature  
	- Message:  
		- Fields: Role, Content  
	- ChatResponse:  
		- Fields: Created, Object, ID, Model, Choices, Usage  
	- Choice:  
		- Fields: Index, FinishReason, Message  
	- UsageStatistics:  
		- Fields: PromptTokens, CompletionTokens, TotalTokens  
	- GenerationResponse:  
		- Fields: Created, ID, Data, Usage  
	- GenerationData:  
		- Fields: Embedding, Index, URL  
	- GenerationUsage:  
		- Fields: PromptTokens, CompletionTokens, TotalTokens  
	- WrongPwdError:  
		- Fields: message  
  
func GenerateCompletion(prompt, modelName string, url string) (*ChatResponse, error):  
	- Creates a ChatRequest struct with the given prompt, model name, and temperature.  
	- Converts the ChatRequest struct to JSON.  
	- Sends a POST request to the given URL with the JSON data.  
	- Reads the response body and parses it as a ChatResponse struct.  
	- Returns the ChatResponse struct and any errors encountered.  
  
func GenerateCompletionWithPWD(prompt, modelName string, url string, s_pwd string, u_pwd string) (*ChatResponse, error):  
	- Checks if the user-provided password matches the stored password.  
	- If the passwords match, calls GenerateCompletion with the given parameters and returns the result.  
	- If the passwords don't match, returns an error.  
  
func GenerateImageStableDissusion(prompt, size string) (string, error):  
	- Creates a payload with the given prompt and size.  
	- Converts the payload to JSON.  
	- Sends a POST request to the given URL with the JSON data.  
	- Reads the response body and parses it as a GenerationResponse struct.  
	- Returns the image URL from the GenerationResponse struct and any errors encountered.  
  
func UploadToTelegraph(fileName string) string:  
	- Gets the absolute path to the file.  
	- Opens the file using the absolute path.  
	- Uploads the file to Telegraph using the telegraph.Upload function.  
	- Returns the uploaded file link.  
  
func deleteFromTemp(fileName string):  
	- Gets the absolute path to the file.  
	- Deletes the file from the temporary directory.  
  
  
  
# lib/bot/command/addNewUsertoMap.go  
package: command  
imports: log, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func AddNewUserToMap(updateMessage *tgbotapi.Message):  
	- extracts chatID from updateMessage  
	- creates a new User struct with chatID, username, DialogStatus: 0, and Admin: false  
	- calls database.AddUser to add the new user to the database  
	- prints a log message with the user's ID and username  
	- creates a new message with the "hello" template and a one-time reply keyboard with a "Start!" button  
	- sends the message to the user's chatID using the bot  
	- (commented out) checks if the user is already registered and assigns the appropriate DialogStatus if they are  
  
  
  
# lib/bot/command/checkAdmin.go  
package: command  
imports: fmt, github.com/JackBekket/hellper/lib/bot/env, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
func CheckAdmin(adminData map[string]env.AdminData, updateMessage *tgbotapi.Message):  
	- chatID := updateMessage.From.ID  
	- iterates through adminData  
	- if admin.ID == chatID:  
		- if admin.GPTKey != "":  
			- calls c.AddAdminToMap(admin.GPTKey, updateMessage)  
			- returns  
		- else:  
			- sends a message to chatID with the message "env \"%s\" is missing."  
			- calls c.AddNewUserToMap(updateMessage)  
			- returns  
	- calls c.AddNewUserToMap(updateMessage)  
  
  
  
# lib/bot/command/newCommander.go  
package: command  
imports: context, github.com/JackBekket/hellper/lib/database, github.com/go-telegram-bot-api/telegram-bot-api/v5  
  
type Commander:  
	- Fields: bot, usersDb, ctx  
	- bot: *tgbotapi.BotAPI  
	- usersDb: map[int64]database.User  
	- ctx: context.Context  
  
func NewCommander(  
	bot *tgbotapi.BotAPI,  
	usersDb map[int64]database.User,  
	ctx context.Context,  
) *Commander:  
	- returns a new instance of Commander with the given bot, usersDb, and ctx  
  
//func GetCommander(): unimplemented  
  
