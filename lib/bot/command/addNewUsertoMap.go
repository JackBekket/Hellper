package command

import (
	"context"
	"log"

	"github.com/JackBekket/hellper/lib/database"
	tgbot "github.com/go-telegram/bot"
	tgbotapi "github.com/go-telegram/bot/models"
)

// Adds a new user to the database and assigns "Dialog_status" = 0.
func (c *Commander) AddNewUserToMap(updateMessage *tgbotapi.Message, base_url string) {
	chatID := updateMessage.Chat.ID
	user := database.User{
		ID:           chatID,
		Username:     updateMessage.From.Username,
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

	msg := tgbot.SendMessageParams{ChatID: user.ID, Text:  msgTemplates["hello"]}
	c.bot.SendMessage(context.Background(),&msg)


}
