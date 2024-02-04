package main

import (
	"context"
	"log"

	"github.com/JackBekket/uncensoredgpt_tgbot/internal/bot/command"
	"github.com/JackBekket/uncensoredgpt_tgbot/internal/bot/env"
	"github.com/JackBekket/uncensoredgpt_tgbot/internal/database"
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
			if update.Message.IsCommand() {
				// Handle the /image command
				if update.Message.Command() == "image" {
				 // Send a random image to the user
				// sendImage(bot, update.Message.Chat.ID)
				if (update.Message.Text =="") {
					comm.GenerateNewImageLAI_SD("evangelion, neon, genesis",chatID,bot)
				} else {
				comm.GenerateNewImageLAI_SD(update.Message.Text,chatID,bot)
				}}}
			log.Println("user dialog status:", user.DialogStatus)
			log.Println(user.ID)
			log.Println(user.Username)
			switch user.DialogStatus {
			// first check for user status, (for a new user status 0 is set automatically),
			// then user reply for the first bot message is logged to a database as name AND user status is updated
			case 0:
				comm.InputYourAPIKey(update.Message) // input key then net then model
			case 1:
				comm.ChooseModel(update.Message)

			case 2:
				comm.HandleModelChoose(update.Message)
				/*
				switch update.Message.Text {
				case "GPT-3.5":
					comm.ModelGPT3DOT5(update.Message)
				// case "GPT-4":
				// 	comm.ModelGPT4(update.Message)
				default:
					comm.WrongModel(update.Message)
				}
				*/
			case 3:
				comm.ConnectingToAiWithLanguage(update.Message,local_access_pwd,ai_endpoint)
			case 4:
				comm.DialogSequence(update.Message,ai_endpoint)
			case 5:
				//comm.CodexSequence(update.Message)
				
			}

		}

	}

} // end of main func
