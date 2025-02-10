package main

import (
	"context"
	"log"
	"os"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/bot/dialog"
	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	token := os.Getenv("TG_KEY")

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

	upd_ch := make(chan tgbotapi.Update)

	//updateHandler :=
	updates := bot.GetUpdatesChan(u)

	// handling any incoming updates through channel
	go dialog.HandleUpdates(upd_ch, bot, *comm)

	for update := range updates {
		upd_ch <- update
	}

} // end of main func
