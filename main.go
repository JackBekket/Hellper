package main

import (
	"context"
	"log"

	"github.com/JackBekket/uncensoredgpt_tgbot/lib/bot/command"
	"github.com/JackBekket/uncensoredgpt_tgbot/lib/bot/env"
	"github.com/JackBekket/uncensoredgpt_tgbot/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Panicf("could not load env from: %v", err)
	}

	token, err := env.LoadTGToken()
	if err != nil {
		log.Panic(err)
	}
	log.Println("TG token is: ", token)

	adminData := env.LoadAdminData()
	local_access_pwd:= env.LoadLocalPD()
	ai_endpoint := env.LoadLocalAI_Endpoint()
	log.Println("ai endpoint is: ", ai_endpoint)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("tg token missing: %v\n", err)
	}

	// init database and commander
	usersDatabase := database.UsersMap
	ctx := context.Background()
	comm := command.NewCommander(bot, usersDatabase, ctx)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message == nil {
			continue
		}

		chatID := update.Message.From.ID
		user, ok := usersDatabase[chatID]
		if !ok {
			comm.CheckAdmin(adminData, update.Message)
		}
		if ok {

			log.Println("user dialog status:", user.DialogStatus)
			log.Println(user.ID)
			log.Println(user.Username)
			switch user.DialogStatus {
			// first check for user status, (for a new user status 0 is set automatically),
			// then user reply for the first bot message is logged to a database as name AND user status is updated
			case 0:
				comm.ChooseNetwork(update.Message)
			case 1:
				comm.HandleNetworkChoose(update.Message)
			case 2:
				comm.InputYourAPIKey(update.Message) 
			case 3:
				comm.ChooseModel(update.Message)
			case 4:
				comm.HandleModelChoose(update.Message)
			case 5:
				comm.ConnectingToAiWithLanguage(update.Message,local_access_pwd,ai_endpoint)	
			case 6: 
				comm.DialogSequence(update.Message,ai_endpoint)
				
			}

		}

	}

} // end of main func
