package command

import (
	"log"

	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Adds a new user to the database and assigns "Dialog_status" = 0.
func (c *Commander) AddNewUserToMap(updateMessage *tgbotapi.Message, base_url string) {
	chatID := updateMessage.Chat.ID
	user := database.User{
		ID:           chatID,
		Username:     updateMessage.From.UserName,
		DialogStatus: 3,
		Admin:        false,
	}
	user.AiSession.Base_url = base_url

	database.AddUser(user)

	//user := c.usersDb[chatID]
	log.Printf(
		"Add new user to database: id: %v, username: %s\n",
		user.ID,
		user.Username,
	)

	msg := tgbotapi.NewMessage(user.ID, msgTemplates["hello"])
	c.bot.Send(msg)


}
