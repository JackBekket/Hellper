package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/JackBekket/hellper/lib/bot/handlers"
	"github.com/JackBekket/hellper/lib/config"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.DateTime,
		TimeLocation: time.Local,
	})

	//In the future, a check for empty variables in the .env file should be implemented
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load .env file")
	}
	log.Info().Msg(".env file loaded successfully")

	token := os.Getenv("TG_KEY")
	db_link := os.Getenv("DB_LINK")

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

	botHandlers := handlers.NewHandlersBot(
		cache, db_service,
		&config.AIConfig{
			AI_endpoint:            os.Getenv("AI_ENDPOINT"),
			BaseURL:                os.Getenv("AI_BASEURL"),
			ImageGenerationModel:   os.Getenv("IMAGE_GENERATION_MODEL"),
			ImageGenerationSuffix:  os.Getenv("IMAGE_GENERATION_SUFFIX"),
			ImageRecognitionModel:  os.Getenv("IMAGE_RECOGNITION_MODEL"),
			ImageRecognitionSuffix: os.Getenv("IMAGE_RECOGNITION_SUFFIX"),
			VoiceRecognitionModel:  os.Getenv("VOICE_RECOGNITION_MODEL"),
			VoiceRecognitionSuffix: os.Getenv("VOICE_RECOGNITION_SUFFIX"),
		},
	)

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
