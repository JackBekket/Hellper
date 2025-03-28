package database

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/vectorstores"
)

// Constants for retrieving information from the Usage map
const (
	Usage_PromptTokens     = "Promt"
	Usage_CompletionTokens = "Completion"
	Usage_TotalTokens      = "Total"
)

// main in-memory struct for dialogs, key (int64) is telegram user id
type User struct {
	ID           int64
	Username     string
	DialogStatus int8
	Admin        bool
	AiSession    AiSession
	Network      string
	Topics       []int
	VectorStore  vectorstores.VectorStore
	//local_ai_pass string
}

type SessionUsage struct {
	ID    int64
	Usage map[string]int
}

// memory
type AiSession struct {
	AIToken      string
	GptModel     string
	AuthMethod   int64
	ProviderID   int64
	DialogThread ChatSessionGraph
	ProviderName string
	BaseURL      string
	Usage        map[string]int
}

// langgraph doesn't work with same types as langchain, so we have to improvise here.
type ChatSessionGraph struct {
	ConversationBuffer []llms.MessageContent
	//DialogThread string

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

func NewChatSessionGraph(buffer []llms.MessageContent) *ChatSessionGraph {
	return &ChatSessionGraph{
		ConversationBuffer: buffer,
	}
}

func (s *ChatSessionGraph) ClearBuffer() {
	s.ConversationBuffer = nil
}
