package main

import (
	"context"
	"os"
	"time"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/bot/dialog"
	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.DateTime,
		TimeLocation: time.Local,
	})

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal().Err(err).Msg("tg token missing")
	}

	// in-memory (cash) db.
	usersDatabase := database.UsersMap

	db, err := database.NewHandler(db_link)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create database service")
	}
	db_service, err := database.NewAIService(db)

	ctx := context.Background()
	comm := command.NewCommander(bot, usersDatabase, ctx)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	upd_ch := make(chan tgbotapi.Update)

	//updateHandler :=
	updates := bot.GetUpdatesChan(u)

	// handling any incoming updates through channel
	go dialog.HandleUpdates(upd_ch, bot, *comm, db_service)

	for update := range updates {
		upd_ch <- update
	}

} // end of main func
