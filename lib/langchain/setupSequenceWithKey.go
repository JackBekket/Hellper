//go:build ignore

package langchain

// Deprecated

//package main

import (
	"context"
	"log"
	"sync"

	db "github.com/JackBekket/hellper/lib/database"

	tgbot "github.com/go-telegram/bot"
)

var mu = sync.Mutex{}

type contextKey string

const UserKey contextKey = "user"

// done
func SetupSequenceWithKey(
	bot *tgbot.Bot,
	user db.User,
	language string,
	ctx context.Context,
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

	switch language {
	case "English":
		response, probe, err := tryLanguage(user, "", 1, ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {

			msg := tgbot.SendMessageParams{ChatID: chatID, Text: response}
			bot.SendMessage(ctx, &msg)
			user.DialogStatus = 6
			user.AiSession.DialogThread = *probe
			usage := db.GetSessionUsage(user.ID)
			user.AiSession.Usage = usage
			//log.Println(user.AiSession.Usage)
			db.UsersMap[chatID] = user
		}
	case "Russian":
		response, probe, err := tryLanguage(user, "", 2, ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbot.SendMessageParams{ChatID: chatID, Text: response}
			bot.SendMessage(ctx, &msg)
			user.AiSession.DialogThread = *probe
			user.DialogStatus = 6
			usage := db.GetSessionUsage(user.ID)
			user.AiSession.Usage = usage
			log.Println(user.AiSession.Usage)
			db.UsersMap[chatID] = user
		}
	default:
		response, probe, err := tryLanguage(user, language, 0, ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbot.SendMessageParams{ChatID: chatID, Text: response}
			bot.SendMessage(ctx, &msg)
			user.AiSession.DialogThread = *probe
			user.DialogStatus = 6
			usage := db.GetSessionUsage(user.ID)
			user.AiSession.Usage = usage
			//log.Println(user.AiSession.Usage)
			db.UsersMap[chatID] = user
		}
	}

}

// LanguageCode: 0 - default, 1 - Russian, 2 - English
// done
func tryLanguage(user db.User, language string, languageCode int, ai_endpoint string) (string, *db.ChatSessionGraph, error) {
	var languagePromt string
	//var languageResponse string
	model := user.AiSession.GptModel

	switch languageCode {
	case 1:
		languagePromt = "Hi, Do you speak english?"
		//languageResponse = "Yes, I do, how can I help you today?"
	case 2:
		languagePromt = "Привет, ты говоришь по-русски?"
		//languageResponse = "Да, я говорю по русски, чем я могу помочь тебе сегодня?"
	default:
		languagePromt = language
	}
	log.Printf("Language: %v\n", languagePromt)

	gptKey := user.AiSession.GptKey
	//model := user.AiSession.GptModel
	//chatID := user.ID

	//result,thread, err := StartNewChat(ctx,gptKey,model,ai_endpoint,languagePromt)
	result, thread, err := RunNewAgent(gptKey, model, ai_endpoint, languagePromt)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return thread, result, nil
}
