package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/bot/dialog"
	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return value, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.DateTime,
		TimeLocation: time.Local,
	})

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load .env file")
	}
	log.Info().Msg(".env file loaded successfully")

	//If at least one environment variable is empty, fatal will be triggered
	token, err := getEnv("TG_KEY")
	db_link, err := getEnv("DB_LINK")
	ai_endpoint, err := getEnv("AI_ENDPOINT")
	baseURL, err := getEnv("AI_BASEURL")
	if err != nil {
		log.Fatal().
			Str("token", token).Str("db_link", db_link).
			Str("ai_endpoint", ai_endpoint).Str("baseURL", baseURL).
			Err(err).Msg("env variable is empty")
	}

	dbHandler, err := database.NewHandler(db_link)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create database service")
	}

	if err := dbHandler.DB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}
	log.Info().Msg("database ping successful")

	db_service, err := database.NewAIService(dbHandler)
	if err != nil {
		log.Fatal().Err(err).Msg("something wrong")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal().Err(err).Msg("tg token missing")
	}

	// in-memory (cash) db.
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
	go dialog.HandleUpdates(upd_ch, bot, *comm, db_service)

	for update := range updates {
		upd_ch <- update
	}

} // end of main func
