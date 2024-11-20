package database

//TODO
// user should be fully functional user class and all operation with user should be placed here (in separate user.go package)

import (
	"log"
	"os"

	e "github.com/JackBekket/hellper/lib/embeddings"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

// main database for dialogs, key (int64) is telegram user id
type User struct {
	ID           int64
	Username     string
	DialogStatus int8
	Admin        bool
	AiSession    AiSession
	Network      string
	Topics       []int
	VectorStore vectorstores.VectorStore
	//local_ai_pass string
}

type SessionUsage struct {
	ID    int64
	Usage map[string]int
}

type AiSession struct {
	GptKey       string
	GptModel     string
	AI_Type      int8
	DialogThread ChatSession
	Base_url     string
	Usage        map[string]int
}

type ChatSession struct {
	ConversationBuffer memory.ConversationBuffer
	DialogThread       chains.LLMChain
}

var UsersMap = make(map[int64]User)
var UsageMap = make(map[int64]SessionUsage)

func AddUser(user User) {
	UsersMap[user.ID] = user
}

func UpdateUserUsage(id int64, usage map[string]int) {
	user, exists := UsersMap[id]
	if exists {
		user.AiSession.Usage = usage
		//UsersMap[id] = user
	}
	UsersMap[id] = user
}

func UpdateSessionUsage(id int64, usage map[string]int) {
	su := UsageMap[id]
	su.ID = id
	su.Usage = usage
	UsageMap[id] = su

}

func GetSessionUsage(id int64) map[string]int {
	usage := UsageMap[id].Usage
	return usage
}

func SetContext(user User, collectionName string) error {
	_ = godotenv.Load()
	api_token := os.Getenv("OPENAI_API_KEY")
	ai_endpoint := os.Getenv("AI_ENDPOINT")
	//log.Println("ai endpoint is: ", ai_endpoint)
	db_link := os.Getenv("EMBEDDINGS_DB_URL")

    vectorStore, err := e.GetVectorStoreWithOptions(ai_endpoint, api_token, db_link, collectionName)
    if err != nil {
        return err
    }
    user.VectorStore = vectorStore

	defer func() {
		var pgvStore pgvector.Store
		pgvStore, ok := vectorStore.(pgvector.Store)
		if !ok {
		  log.Fatalf("store does not implement pgvector.Store")
		}
		pgvStore.Close()
	  }()


    return nil
}


func ClearContext(user User) {
	user.VectorStore = nil
}


