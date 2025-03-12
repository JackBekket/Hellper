package handlers

import (
	"context"
	"strings"

	"github.com/JackBekket/hellper/lib/config"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// 1. Enforce encapsulation: external packages should not have access to internal `handler` structures except through the interface.
// 2. All functions that send messages to the user on behalf of the bot should be implemented exclusively in this package.
// 3. The logic for sending messages should not be moved to other packages.
// 4. Configuration and dependencies should only be passed through the constructor. Global variables are prohibited.

// Thanks! (=ↀωↀ=)

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
	dbLink      string
	// Postgres database and LLMHandlers
	dbService *database.Service
	config    *config.AIConfig

	// Pass dependencies here
}

// Constructor of the handlers type
func NewHandlersBot(cache database.Cacher, db_service *database.Service, dbLink string, config *config.AIConfig) Bot {
	return &handlers{
		cache:     cache,
		dbService: db_service,
		dbLink:    dbLink,
		config:    config,
	}
}

// Central function for registering bot handlers. New used commands should be added here
func (h *handlers) NewRegisterHandlers(ctx context.Context, tgb *bot.Bot) {
	botSelf, _ := tgb.GetMe(ctx)
	h.botUsername = botSelf.Username
	// Router for tg bot command handlers
	tgb.RegisterHandlerMatchFunc(matchTypePrefix, h.cmdRouter)

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
//The telegram file handling moment needs to be processed.

func matchTextMessage(update *models.Update) bool {
	return update.Message != nil && update.Message.Text != "" && update.CallbackQuery == nil && update.Message.Sticker == nil
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

// Stub for the /start command, it needs to be designed.
func matchTypePrefix(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	data := update.Message.Text
	return strings.HasPrefix(data, "/") && data != "/start"
}
