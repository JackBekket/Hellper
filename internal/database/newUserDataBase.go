package database

import gogpt "github.com/sashabaranov/go-openai"

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
	GptClient gogpt.Client
	GptModel  string
}

var UsersMap = make(map[int64]User)
