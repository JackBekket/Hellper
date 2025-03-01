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
	cache database.Cacher
	// Postgres database and LLMHandlers
	db_Link    string
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
	// Router for tg bot command handlers
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, h.cmdRouter, cmdHandlerMiddleware)

	// Router for text message handlers
	tgb.RegisterHandlerMatchFunc(matchTextMessage, h.textMessageRouter)

	// Router for callbacks
	tgb.RegisterHandlerMatchFunc(
		matchCallbackQuery,
		h.callbackRouter,
		callbackSingleExecutionMiddleWare,
	)

	tgb.RegisterHandlerMatchFunc(matchPhoto, h.photoHandler)

	tgb.RegisterHandlerMatchFunc(matchVoice, h.voiceHandler)

	tgb.RegisterHandlerMatchFunc(matchTgGroup, h.tgGroupHandler)

}

// Rules for calling the handler

func matchTextMessage(update *models.Update) bool {
	return update.Message.Text != "" && update.Message.Photo == nil
}

func matchCallbackQuery(update *models.Update) bool {
	return update.CallbackQuery != nil
}

func matchPhoto(update *models.Update) bool {
	return update.Message.Photo != nil
}

func matchVoice(update *models.Update) bool {
	return update.Message.Voice != nil
}

func matchTgGroup(update *models.Update) bool {
	// Stub for the registration func
	return false
}
