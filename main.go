package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/JackBekket/hellper/lib/bot/handlers"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
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

	cache := database.NewMemoryCache()

	botHandlers := handlers.NewHandlersBot(cache, db_service, ai_endpoint, baseURL)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(botHandlers.IdentifyUserMiddleware),
	}

	tgb, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("token is missing")
	}

	botHandlers.NewRegisterHandlers(ctx, tgb)

	botSelf, err := tgb.GetMe(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid token")
	}

	go tgb.Start(ctx)
	log.Info().Msg("Bot is starting")
	log.Info().Int64("id", botSelf.ID).Msgf("authorized on account: %s", botSelf.Username)

	<-ctx.Done()
	log.Info().Msg("Termination signal received. Shutting down...")
	if err := dbHandler.DB.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close database connection")
	} else {
		log.Info().Msg("database connection closed")
	}
	log.Info().Msg("Completed.")

}
