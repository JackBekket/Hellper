package command

import (
	"log"

	"github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Adds a new user to the database and assigns "Dialog_status" = 0.
func (c *Commander) AddNewUserToMap(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	c.usersDb[chatID] = database.User{
		ID:           chatID,
		Username:     updateMessage.From.UserName,
		DialogStatus: 0,
		Admin:        false,
	}

	user := c.usersDb[chatID]
	log.Printf(
		"Add new user to database: id: %v, username: %s\n",
		user.ID,
		user.Username,
	)

	msg := tgbotapi.NewMessage(user.ID, msgTemplates["hello"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Start!")),
	)
	c.bot.Send(msg)

	// check for registration
	//	registred := IsAlreadyRegistred(session, chatID)
	/*
		if registred {
			c.usersDb[chatID] = db.User{updateMessage.Chat.ID, updateMessage.Chat.UserName, 1}
		}
	*/
}
