package localai

import (
	"context"
	"log"

	db "github.com/JackBekket/uncensoredgpt_tgbot/internal/database"
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

func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string) {
	mu.Lock()
	defer mu.Unlock()

	user := db.UsersMap[chatID]

	gptModel := user.AiSession.GptModel
	log.Printf(
		"GPT model: %s,\npromt: %s\n",
		gptModel,
		promt,
	)

	/*
	req := createComplexChatRequest(promt, gptModel)
	c := user.AiSession.GptClient
	*/

	resp, err := GenerateCompletion(promt,gptModel,ai_endpoint)
	if err != nil {
		errorMessage(err, bot, user)
	} else {
		LogResponse(resp)
		respText := resp.Choices[0].Message.Content
		msg := tgbotapi.NewMessage(chatID, respText)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 4
		db.UsersMap[chatID] = user
	}

}

func LogResponse(resp *ChatResponse) {
	log.Println("full response obj log: ", resp)
	log.Println("created: ", resp.Created)
	log.Println("resp id: ",resp.ID)
	log.Println("resp model: ", resp.Model)
	log.Println("resp object: ",resp.Object)
	log.Println("resp Choices[0]: ", resp.Choices[0])
	log.Println("resp_usage promt tokens: ", resp.Usage.PromptTokens)
	log.Println("resp_usage completion tokens: ", resp.Usage.CompletionTokens)
	log.Println("resp_usage total tokens: ",resp.Usage.TotalTokens)
}
