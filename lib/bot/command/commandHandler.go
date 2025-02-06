package command

import (
	"log"
	"os"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommands(message *tgbotapi.Message, comm *Commander) {
	bot := comm.bot
	chatID := message.Chat.ID
	user := comm.GetUser(chatID)

	switch message.Command() {

	case "image":
		msg := tgbotapi.NewMessage(user.ID, "Image link generation...")
		bot.Send(msg)
		baseUrl := os.Getenv("AI_ENDPOINT")
		promt := message.CommandArguments()
		log.Printf("Command /image arg: %s\n", promt)
		if promt == "" {
			comm.GenerateNewImageLAI_SD("evangelion, neon, anime", baseUrl, chatID, bot)
		} else {
			comm.GenerateNewImageLAI_SD(promt, baseUrl, chatID, bot)
		}
		//go openaibot.StartImageSequence(c.bot, updateMessage, chatID, promt, c.ctx)
	case "restart":
		msg := tgbotapi.NewMessage(user.ID, "Restarting session..., type any key")
		bot.Send(msg)
		userDb := database.UsersMap
		delete(userDb, user.ID)
	case "help":
		comm.HelpCommandMessage(message)
	case "search_doc":
		promt := message.CommandArguments()
		comm.SearchDocuments(chatID, promt, 3)
	case "rag":
		promt := message.CommandArguments()
		comm.RAG(chatID, promt, 1)
	case "instruct":
		// this is calling local-ai within base template (and without langhain injections)
		promt := message.CommandArguments()
		model_name := user.AiSession.GptModel
		api_token := user.AiSession.GptKey
		langchain.GenerateContentInstruction(user.AiSession.Base_url, promt, model_name, api_token, user.Network)
	case "usage":
		comm.GetUsage(chatID)
	case "helper":
		comm.SendMediaHelper(chatID)
	case "setContext":
		name := message.CommandArguments()
		userDb := database.UsersMap
		user := userDb[chatID]
		log.Println("comnmand set context")
		log.Println("argument: ", name)
		log.Println("user:", user)
		user.SetContext(name)
	case "clearContext":
		user := comm.GetUser(chatID)
		user.ClearContext()
	default:
	}
}
