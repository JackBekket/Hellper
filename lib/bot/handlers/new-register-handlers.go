package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// structure to hold dependencies of other packages: postgres, cache, llmHandlers
type handlers struct {
	/*Pass dependencies here */
}

// Constructor of the handlers type
func NewHandlersBot( /*Pass dependencies here */ ) *handlers {
	return &handlers{}
}

// Central function for registering bot handlers. New used commands should be added here
func (h *handlers) NewRegisterHandlers(ctx context.Context, tgb *bot.Bot) {
	// Router for tg bot command handlers
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, h.cmdRouter)

	// Router for text message handlers
	tgb.RegisterHandlerMatchFunc(matchTextMessage, h.textMessageRouter)

	// Router for callbacks
	tgb.RegisterHandlerMatchFunc(matchCallbackQuery, h.callbackRouter)

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
