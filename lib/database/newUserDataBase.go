package database

import (
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/memory"
)

//import "github.com/JackBekket/uncensoredgpt_tgbot/lib/langchain"

// main database for dialogs, key (int64) is telegram user id
type User struct {
	ID           int64
	Username     string
	DialogStatus int8
	Admin        bool
	AiSession    AiSession
	Network      string
	local_ai_pass string
}



type AiSession struct {
	GptKey    string
	GptModel  string
	AI_Type	  int8
	DialogThread ChatSession
}

type ChatSession struct {
    ConversationBuffer *memory.ConversationBuffer
    DialogThread *chains.LLMChain
}

var UsersMap = make(map[int64]User)
