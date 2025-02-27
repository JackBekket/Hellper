package handlers

import (
	"context"

	"github.com/go-telegram/bot"
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
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, h.cmdRouter)

	// Router for text message handlers
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, h.textMessageRouter)

}
