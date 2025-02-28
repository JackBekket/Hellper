package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

//todo

const (
	status_AIModelSelectionChoice = "AIModelSelectionChoice" // old dialogStatus - 4
	status_ConnectingToAiWithLang = "connectingToAiWithLang" // old dialogStatus - 5

)

func (h *handlers) callbackRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	// Stub for the user cache structure
	dialogStatus := status_AIModelSelectionChoice
	switch dialogStatus {
	case status_AIModelSelectionChoice:
		h.handleAIModelSelectionCallback(ctx, tgb, update)
	case status_ConnectingToAiWithLang:
		h.handleConnectingToAiWithLangCallback(ctx, tgb, update)

	default: // todo: error msg
	}

}
