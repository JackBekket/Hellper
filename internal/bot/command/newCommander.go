package command

import (
	"context"

	"github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	usersDb map[int64]database.User
	ctx     context.Context
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	usersDb map[int64]database.User,
	ctx context.Context,
) *Commander {
	return &Commander{
		bot:     bot,
		usersDb: usersDb,
		ctx:     ctx,
	}
}
