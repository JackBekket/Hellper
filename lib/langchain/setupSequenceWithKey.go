package langchain

//package main

import (
	"context"
	"log"
	"sync"

	db "github.com/JackBekket/hellper/lib/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mu = sync.Mutex{}
type contextKey string
const UserKey contextKey = "user"


func SetupSequenceWithKey(
	bot *tgbotapi.BotAPI,
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

	//tgbot = bot

	//ctx = context.WithValue(ctx, "user", user)

	/*
	// Initializing empty dialog thread
	thread, err := InitializeNewChatWithContextNoLimit(gptKey,user.AiSession.GptModel,ai_endpoint)
	if err != nil {
		log.Println(err)
	}
	user.AiSession.DialogThread = *thread
	db.UsersMap[chatID] = user // we need to store empty buffer *before* starting dialog
	*/
	

	switch language {
	case "English":
		response,probe, err := tryLanguage(user, "", 1, ctx,ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {

			msg := tgbotapi.NewMessage(chatID, response)
			bot.Send(msg)
			user.DialogStatus = 6
			user.AiSession.DialogThread = *probe
			usage := db.GetSessionUsage(user.ID)
			user.AiSession.Usage = usage
			//log.Println(user.AiSession.Usage)
			db.UsersMap[chatID] = user
		}
	case "Russian":
		response,probe, err := tryLanguage(user, "", 2, ctx,ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbotapi.NewMessage(chatID, response)
			bot.Send(msg)
			user.AiSession.DialogThread = *probe
			user.DialogStatus = 6
			usage := db.GetSessionUsage(user.ID)
			user.AiSession.Usage = usage
			log.Println(user.AiSession.Usage)
			db.UsersMap[chatID] = user
		}
	default:
		response,probe, err := tryLanguage(user, language, 0, ctx,ai_endpoint)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbotapi.NewMessage(chatID, response)
			bot.Send(msg)
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
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context, ai_endpoint string) (string,*db.ChatSession, error) {
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

	// Initializing empty dialog thread
	result,thread, err := StartNewChat(ctx,gptKey,model,ai_endpoint,languagePromt)
		if err != nil {
			log.Println(err)
			return "",nil,err
		}

	//user.AiSession.DialogThread = *thread
	//db.UsersMap[chatID] = user // we need to store empty buffer *before* starting dialog

	return result,thread, nil
	
	
	/*
	resp, err := ContinueChatWithContextNoLimit(thread,languagePromt)
	if err != nil {
		return "", err
	} else {
		return resp, nil
	}
	*/

	/*
	resp, err := GenerateContentLAI(ai_endpoint,model,languagePromt)
	if err != nil {
		return "", err
	} else {
		LogResponseContentChoice(resp)
		answer := resp.Choices[0].Content
		return answer, nil
	}
	*/
}
