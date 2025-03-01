//go:build ignore

package command

import (
	"context"
	"log"
	"os"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	tgbot "github.com/go-telegram/bot"
	tgbotapi "github.com/go-telegram/bot/models"
)

// old api
// done
func HandleCommands(message *tgbotapi.Message, comm *Commander, ds *database.Service) {
	bot := comm.bot
	chatID := message.Chat.ID
	user := comm.GetUser(chatID)

	switch message.Command() {

	case "image":
		msg := tgbot.SendMessageParams{ChatID: user.ID, Text: "Image link generation..."}
		bot.SendMessage(context.Background(), &msg)
		baseUrl := os.Getenv("AI_ENDPOINT")
		promt := message.CommandArguments()
		log.Printf("Command /image arg: %s\n", promt)
		if promt == "" {
			comm.GenerateNewImageLAI_SD("evangelion, neon, anime", baseUrl, chatID, bot)
		} else {
			comm.GenerateNewImageLAI_SD(promt, baseUrl, chatID, bot)
		}
		//go openaibot.StartImageSequence(c.bot, updateMessage, chatID, promt, c.ctx)
	case "reload":
		msg := tgbot.SendMessageParams{ChatID: user.ID, Text: "Reloading session..., type any key"}
		bot.SendMessage(context.Background(), &msg)
		userDb := database.UsersMap
		delete(userDb, user.ID)
	case "clear":
		msg := tgbot.SendMessageParams{ChatID: user.ID, Text: "Deleting dialog thread from database..., type any key"}
		bot.SendMessage(context.Background(), &msg)
		user.FlushMemory(ds)
		userDb := database.UsersMap
		delete(userDb, user.ID)
	case "purge":
		msg := tgbot.SendMessageParams{ChatID: user.ID, Text: "Deleting all user data from database and restarting session..., type any key"}
		bot.SendMessage(context.Background(), &msg)
		user.Kill(ds)
	case "drop":
		msg := tgbot.SendMessageParams{ChatID: user.ID, Text: "Dropping session..., type any key"}
		bot.SendMessage(context.Background(), &msg)
		user.DropSession(ds)
		user.FlushMemory(ds)
		userDb := database.UsersMap
		delete(userDb, user.ID)
	case "help":
		comm.HelpCommandMessage(message)
	case "search_doc":
		promt := message.CommandArguments()
		comm.SearchDocuments(chatID, promt, 3)
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

// in new approach we should delcare and register handler for each sepatare command (?)

// example
func helloHandler(ctx context.Context, b *tgbot.Bot, update *tgbotapi.Update) {
	b.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + tgbot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: tgbotapi.ParseModeMarkdown,
	})
}

/*
func imageHandler (ctx context.Context, b *tgbot.Bot, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	user_id := update.Message.From.ID
	msg := tgbot.SendMessageParams{ChatID: user_id, Text: "Image link generation..."}
		bot.SendMessage(context.Background(), &msg)
		baseUrl := os.Getenv("AI_ENDPOINT")
		promt :=
}
*/
