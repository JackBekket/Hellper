package dialog

import (
	"log"

	comm "github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleUpdates(updates <-chan tgbotapi.Update, bot tgbotapi.BotAPI)  {
	


for update := range updates {

	chatID := update.Message.From.ID
	user, ok := usersDatabase[chatID]
	if !ok {
		comm.CheckAdmin(adminData, update.Message)
	}
	if ok {

		switch update.Message.Command() {

		case "image":
			msg := tgbotapi.NewMessage(user.ID, "Image link generation...")
			bot.Send(msg)

			promt := update.Message.CommandArguments()
			log.Printf("Command /image arg: %s\n", promt)
			if (promt == "") {
				comm.GenerateNewImageLAI_SD("evangelion, neon, anime",chatID,bot)
			} else {
				comm.GenerateNewImageLAI_SD(promt,chatID,bot)
			}
			//go openaibot.StartImageSequence(c.bot, updateMessage, chatID, promt, c.ctx)

		case "restart":
			msg := tgbotapi.NewMessage(user.ID, "Restarting session..., type any key")
			bot.Send(msg)
			userDb := database.UsersMap
			delete(userDb, user.ID)
		case "help":
			comm.HelpCommandMessage(update.Message)
		case "search_doc":
			promt := update.Message.CommandArguments()
			comm.SearchDocuments(chatID,promt,3)
		case "rag":
			promt := update.Message.CommandArguments()
			comm.RAG(chatID,promt,1)
		case "instruct" :
			// this is calling local-ai within base template (and without langhain injections)
			promt := update.Message.CommandArguments()
			model_name := user.AiSession.GptModel
			api_token := user.AiSession.GptKey
			langchain.GenerateContentInstruction(user.AiSession.Base_url,promt,model_name,api_token,user.Network)
		case "usage" :
			comm.GetUsage(chatID)
		case "helper":
			comm.SendMediaHelper(chatID)
	}

	if update.Message == nil {
		continue
	}

	//chatID := update.Message.From.ID
	//user, ok := usersDatabase[chatID]
	if !ok {
		comm.CheckAdmin(adminData, update.Message)
	}
	if ok {

		log.Println("user dialog status:", user.DialogStatus)
		log.Println(user.ID)
		log.Println(user.Username)
		switch user.DialogStatus {
		// first check for user status, (for a new user status 0 is set automatically),
		// then user reply for the first bot message is logged to a database as name AND user status is updated
		case 0:
			comm.ChooseNetwork(update.Message)
		case 1:
			comm.HandleNetworkChoose(update.Message)
		case 2:
			comm.InputYourAPIKey(update.Message) 
		case 3:
			comm.ChooseModel(update.Message)
		case 4:
			comm.HandleModelChoose(update.Message)
		case 5:
			comm.ConnectingToAiWithLanguage(update.Message, ai_endpoint)	
		case 6: 
			comm.DialogSequence(update.Message,ai_endpoint)
			
		}

	}

	}

}
}  // end of main func