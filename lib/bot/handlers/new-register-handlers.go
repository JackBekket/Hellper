package handlers

import (
	"context"

	"github.com/JackBekket/hellper/lib/config"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot interface {
	// The global middleware checks for the user's presence in the cache and PostgreSQL.
	// If the user is absent, it initiates the registration or data recovery process
	IdentifyUserMiddleware(next bot.HandlerFunc) bot.HandlerFunc
	// Central function for registering bot handlers. New used commands should be added here
	NewRegisterHandlers(ctx context.Context, tgb *bot.Bot)
}

// structure to hold dependencies of other packages: postgres, cache, llmHandlers
type handlers struct {
	botUsername string
	cache       database.Cacher
	db_Link     string
	// Postgres database and LLMHandlers
	db_service *database.Service
	config     *config.AIConfig

	// Pass dependencies here
}

// Constructor of the handlers type
func NewHandlersBot(cache database.Cacher, db_service *database.Service, db_Link string, config *config.AIConfig) Bot {
	return &handlers{
		cache:      cache,
		db_service: db_service,
		db_Link:    db_Link,
		config:     config,
	}
}

// Central function for registering bot handlers. New used commands should be added here
func (h *handlers) NewRegisterHandlers(ctx context.Context, tgb *bot.Bot) {
	botSelf, _ := tgb.GetMe(ctx)
	h.botUsername = botSelf.Username
	// Router for tg bot command handlers
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, h.cmdRouter)

	// Router for text message handlers
	tgb.RegisterHandlerMatchFunc(matchTextMessage, h.textMessageRouter, h.filterGroupMessagesMiddleware)

	// Router for callbacks
	tgb.RegisterHandlerMatchFunc(
		matchCallbackQuery,
		h.callbackRouter,
		callbackSingleExecutionMiddleWare,
	)

	tgb.RegisterHandlerMatchFunc(matchPhoto, h.handleRecognizeImage, h.filterGroupMessagesMiddleware)
	tgb.RegisterHandlerMatchFunc(matchVoice, h.handleVoiceTranscriber, h.filterGroupMessagesMiddleware)

}

// Rules for calling the handler

func matchTextMessage(update *models.Update) bool {
	return update.Message != nil && update.Message.Text != ""
}

func matchCallbackQuery(update *models.Update) bool {
	return update.CallbackQuery != nil
}

func matchPhoto(update *models.Update) bool {
	return update.Message != nil && update.Message.Photo != nil
}

func matchVoice(update *models.Update) bool {
	return update.Message != nil && update.Message.Voice != nil
}
