package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Router for tg bot command handlers
func (h *handlers) cmdRouter(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (h *handlers) cmdImage(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdReload(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdClear(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdPurge(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdDrop(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdHelp(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdSearchDoc(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdInstruct(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdUsage(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdHelper(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdSetContext(ctx context.Context, tgb *bot.Bot, msg *models.Update)

func (h *handlers) cmdClearContext(ctx context.Context, tgb *bot.Bot, msg *models.Update)
