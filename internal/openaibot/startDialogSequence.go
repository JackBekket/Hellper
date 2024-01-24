package openaibot

import (
	"context"
	"log"

	db "github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Notifies the user that an error occurred while creating the request.
// "An error has occured. In order to proceed we need to recreate client and initialize new session"
// Removes a user from the database (тemporary solution).
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User) {
	log.Println("error :", err)
	msg := tgbotapi.NewMessage(user.ID, err.Error())
	bot.Send(msg)
	msg = tgbotapi.NewMessage(user.ID, "an error has occured. In order to proceed we need to recreate client and initialize new session")
	bot.Send(msg)

	userDb := db.UsersMap
	delete(userDb, user.ID)
	// updateDb := userDatabase[ID]
	// updateDb.Dialog_status = 0
	// userDatabase[ID] = updateDb
}

func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()

	user := db.UsersMap[chatID]

	gptModel := user.AiSession.GptModel
	log.Printf(
		"GPT model: %s,\npromt: %s\n",
		gptModel,
		promt,
	)

	req := createComplexChatRequest(promt, gptModel)
	c := user.AiSession.GptClient

	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		errorMessage(err, bot, user)
	} else {
		respText := resp.Choices[0].Message.Content
		msg := tgbotapi.NewMessage(chatID, respText)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 4
		db.UsersMap[chatID] = user
	}

}
