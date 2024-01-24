package openaibot

import (
	"context"
	"log"
	"sync"

	db "github.com/JackBekket/telegram-gpt/internal/database"
	"github.com/sashabaranov/go-openai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mu = sync.Mutex{}

func SetupSequenceWithKey(
	bot *tgbotapi.BotAPI,
	user db.User,
	language string,
	ctx context.Context,
	local_ap string,
) {
	mu.Lock()
	defer mu.Unlock()
	chatID := user.ID
	gptKey := user.AiSession.GptKey
	log.Println("user GPT key from session: ", gptKey)
	u_network := user.Network
	log.Println("user network from session: ", u_network)
	log.Println("user model from session: ", user.AiSession.GptModel)
	var client *openai.Client
	//u_pwd := 


	if (u_network == "openai") {
		log.Println("Setting up sequence with key")
		log.Println("Network is openai, initiating creating client")
		client = CreateClient(gptKey)
	} 
	if (u_network == "localai") {
		log.Println("Network is localhost, creating local client")
		//var err error
		client_check, err := CreateLocalhostClientWithCheck(local_ap,gptKey)
		if err != nil {
			errorMessage(err,bot,user)
		} 
		client = client_check
	}

	//client := CreateClient(gptKey) // creating client (but we don't know if it works)
	//log.Println("Setting up sequence with key")
	//client := CreateLocalhostClientWithCheck(local_ap,gptKey)
	log.Println("local_ap: ", local_ap)
	log.Println("client: ", client)
	//log.Println("client: ", client.config)
	user.AiSession.GptClient = *client

	

	switch language {
	case "English":
		probe, err := tryLanguage(user, "", 1, ctx)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbotapi.NewMessage(chatID, probe)
			bot.Send(msg)
			user.DialogStatus = 4
			db.UsersMap[chatID] = user
		}
	case "Russian":
		probe, err := tryLanguage(user, "", 2, ctx)
		if err != nil {
			errorMessage(err, bot, user)
		} else {
			msg := tgbotapi.NewMessage(chatID, probe)
			bot.Send(msg)
			user.DialogStatus = 4
			db.UsersMap[chatID] = user
		}
	default:
		probe, err := tryLanguage(user, language, 0, ctx)
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

// LanguageCode: 0 - default, 1 - Russian, 2 - English
func tryLanguage(user db.User, language string, languageCode int, ctx context.Context) (string, error) {
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
	client := user.AiSession.GptClient
	log.Println("client: ", client)

	req := createComplexChatRequest(languagePromt, model)
	log.Printf("request: %v\n", req)

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	} else {
		answer := resp.Choices[0].Message.Content
		return answer, nil
	}
}
