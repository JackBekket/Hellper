package openaibot

import (
	"context"
	"log"

	db "github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCodexSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()
	user := db.UsersMap[chatID]
	log.Printf(
		"GPT model: %s,\npromt: %s\n",
		user.AiSession.GptModel,
		promt,
	)

	req := createCodexRequest(promt)
	c := user.AiSession.GptClient

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		errorMessage(err, bot, user)
	} else {
		respText := resp.Choices[0].Text
		msg := tgbotapi.NewMessage(chatID, respText)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 5
		db.UsersMap[chatID] = user
	}

}
