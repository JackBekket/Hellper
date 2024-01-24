package openaibot

import (
	"context"
	"fmt"
	"log"

	db "github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartImageSequence(
	bot *tgbotapi.BotAPI,
	updateMessage *tgbotapi.Message,
	chatID int64,
	promt string,
	ctx context.Context,
) {
	mu.Lock()
	defer mu.Unlock()
	user := db.UsersMap[chatID]

	req := createImageRequest(promt)
	c := user.AiSession.GptClient

	resp, err := c.CreateImage(ctx, req)
	if err != nil {
		errorMessage(err, bot, user)
	} else {

		respUrl := resp.Data[0].URL
		log.Printf("url image: %s\n", respUrl)

		msg1 := tgbotapi.NewMessage(chatID, "Done!")
		bot.Send(msg1)

		msg := tgbotapi.NewEditMessageText(
			chatID,
			updateMessage.MessageID+1,
			fmt.Sprintf("[Result](%s)", respUrl),
		)

		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 4
		db.UsersMap[chatID] = user
	}
}
