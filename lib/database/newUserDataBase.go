package database

import (
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
)

// main database for dialogs, key (int64) is telegram user id
type User struct {
	ID           int64
	Username     string
	DialogStatus int8
	Admin        bool
	AiSession    AiSession
	Network      string
	//local_ai_pass string
}

type SessionUsage struct {
	ID			int64
	Usage		map[string]int
}


type AiSession struct {
	GptKey    string
	GptModel  string
	AI_Type	  int8
	DialogThread ChatSession
	Base_url  string
	Usage	  map[string]int
}

type ChatSession struct {
    ConversationBuffer memory.ConversationBuffer
    DialogThread chains.LLMChain
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




func UpdateSessionUsage(id int64, usage map[string]int)  {
	su := UsageMap[id]
	su.ID = id
	su.Usage = usage
	UsageMap[id] = su

}

func GetSessionUsage(id int64) (map[string]int) {
	usage := UsageMap[id].Usage
	return usage
}
