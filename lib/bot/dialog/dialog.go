package dialog

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/JackBekket/hellper/lib/bot/command"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(updates <-chan tgbotapi.Update, bot *tgbotapi.BotAPI, comm command.Commander) {

	for update := range updates {
		if update.CallbackQuery == nil {

			var group = false
			if update.Message.Chat.ID < 0 {
				group = true
			}
			if group && !strings.Contains(update.Message.Text, bot.Self.UserName) && update.Message.Voice == nil && update.Message.Photo == nil && update.Message.Command() == "" {
				continue
			}

			if group && update.Message.Photo != nil && !strings.Contains(update.Message.Caption, bot.Self.UserName) {
				continue
			} else {
				re := regexp.MustCompile(`@?` + regexp.QuoteMeta(bot.Self.UserName))
				update.Message.Caption = re.ReplaceAllString(update.Message.Caption, "")
				update.Message.Caption = strings.TrimSpace(update.Message.Caption)
			}

			chatID := int64(update.Message.Chat.ID)
			db := comm.GetUsersDb()
			user, ok := db[int64(chatID)]
			if !ok {
				comm.AddNewUserToMap(update.Message)
			}
			ai_endpoint := os.Getenv("AI_ENDPOINT")

			if ok {

				if update.Message == nil {
					continue
				}

				if !ok {
					comm.AddNewUserToMap(update.Message)
				}
				if ok {

					log.Println("user dialog status:", user.DialogStatus)
					log.Println(user.ID)
					log.Println(user.Username)

					if group && update.Message.Voice != nil && user.DialogStatus != 6 {
						continue
					}
					if update.Message.Text != "" {
						re := regexp.MustCompile(`@?` + regexp.QuoteMeta(bot.Self.UserName))
						update.Message.Text = re.ReplaceAllString(update.Message.Text, "")
					}
					switch user.DialogStatus {
					// first check for user status, (for a new user status 0 is set automatically),
					// then user reply for the first bot message is logged to a database as name AND user status is updated
					case 0:
						fallthrough
					case 1:
						fallthrough
					case 2:
						comm.InputYourAPIKey(update.Message)
					case 3:
						comm.ChooseModel(update.Message)
					case 4, 5:
						comm.WrongResponse(update.Message)
					case 6:
						comm.DialogSequence(update.Message, ai_endpoint)

					}

				}

			} // usual handle end

		} else {
			//here goes the callback logic for inlines
			chatID := int64(update.CallbackQuery.Message.Chat.ID)
			db := comm.GetUsersDb()
			user := db[int64(chatID)]
			ai_endpoint := os.Getenv("AI_ENDPOINT")

			switch user.DialogStatus {
			case 4:
				comm.HandleModelChoose(update.CallbackQuery)
			case 5:
				comm.ConnectingToAiWithLanguage(update.CallbackQuery, ai_endpoint)
			}
		}
	} // end of main func
}
