//go:build ignore

package command

import (
	"context"

	"github.com/JackBekket/hellper/lib/database"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram/bot"
)

type Commander struct {
	bot     *tgbotapi.Bot
	usersDb map[int64]database.User
	ctx     context.Context
}

func NewCommander(
	bot *tgbotapi.Bot,
	usersDb map[int64]database.User,
	ctx context.Context,
) *Commander {
	return &Commander{
		bot:     bot,
		usersDb: usersDb,
		ctx:     ctx,
	}
}
