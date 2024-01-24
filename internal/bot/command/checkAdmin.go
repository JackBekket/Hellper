package command

import (
	"fmt"

	"github.com/JackBekket/telegram-gpt/internal/bot/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Updates "dialogStatus" in the database. Admins - 2, other users - 0.
//
// Loads the key from env into the database.
func (c *Commander) CheckAdmin(adminData map[string]env.AdminData, updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	for evn, admin := range adminData {
		if admin.ID == chatID {
			if admin.GPTKey != "" {
				c.AddAdminToMap(admin.GPTKey, updateMessage)
				return
			} else {
				msg := tgbotapi.NewMessage(
					chatID,
					fmt.Sprintf("env \"%s\" is missing.", evn),
				)
				c.bot.Send(msg)
				// Directs to case 0
				c.AddNewUserToMap(updateMessage)
				return
			}
		}
	}
	c.AddNewUserToMap(updateMessage)
}
