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

	// Telegram bot command handlers
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/image", bot.MatchTypeExact, h.cmdImage)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/reload", bot.MatchTypeExact, h.cmdReload)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/clear", bot.MatchTypeExact, h.cmdClear)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/purge", bot.MatchTypeExact, h.cmdPurge)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/drop", bot.MatchTypeExact, h.cmdDrop)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, h.cmdHelp)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/search_doc", bot.MatchTypeExact, h.cmdSearchDoc)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/instruct", bot.MatchTypeExact, h.cmdInstruct)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/usage", bot.MatchTypeExact, h.cmdUsage)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/helper", bot.MatchTypeExact, h.cmdHelper)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/setContext", bot.MatchTypeExact, h.cmdSetContext)
	tgb.RegisterHandler(bot.HandlerTypeMessageText, "/clearContext", bot.MatchTypeExact, h.cmdClearContext)

}
