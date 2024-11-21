package dialog

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
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
				//comm.CheckAdmin(adminData, update.Message)
				comm.AddNewUserToMap(update.Message)
			}
			ai_endpoint := os.Getenv("AI_ENDPOINT")

			if ok {
				//chatID = int64(chatID)

				switch update.Message.Command() {

				case "image":
					msg := tgbotapi.NewMessage(user.ID, "Image link generation...")
					bot.Send(msg)
					baseUrl := os.Getenv("AI_ENDPOINT")
					promt := update.Message.CommandArguments()
					log.Printf("Command /image arg: %s\n", promt)
					if promt == "" {
						comm.GenerateNewImageLAI_SD("evangelion, neon, anime", baseUrl, chatID, bot)
					} else {
						comm.GenerateNewImageLAI_SD(promt, baseUrl, chatID, bot)
					}
					//go openaibot.StartImageSequence(c.bot, updateMessage, chatID, promt, c.ctx)

					//TODO: Consider adding continue to all other command options, since Hellper actively tries to answer on commands after their execution x)
					continue
				case "restart":
					msg := tgbotapi.NewMessage(user.ID, "Restarting session..., type any key")
					bot.Send(msg)
					userDb := database.UsersMap
					delete(userDb, user.ID)
				case "help":
					comm.HelpCommandMessage(update.Message)
				case "search_doc":
					promt := update.Message.CommandArguments()
					comm.SearchDocuments(chatID, promt, 3)
				case "rag":
					promt := update.Message.CommandArguments()
					comm.RAG(chatID, promt, 1)
				case "instruct":
					// this is calling local-ai within base template (and without langhain injections)
					promt := update.Message.CommandArguments()
					model_name := user.AiSession.GptModel
					api_token := user.AiSession.GptKey
					langchain.GenerateContentInstruction(user.AiSession.Base_url, promt, model_name, api_token, user.Network)
				case "usage":
					comm.GetUsage(chatID)
				case "helper":
					comm.SendMediaHelper(chatID)
				case "setContext":
					name := update.Message.CommandArguments()
					user := comm.GetUser(chatID)
					database.SetContext(user,name)
				case "clearContext":
					user := comm.GetUser(chatID)
					database.ClearContext(user)
				case "":
				default:
					continue
				}

				if update.Message == nil {
					continue
				}

				//chatID := update.Message.From.ID
				//user, ok := usersDatabase[chatID]

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
