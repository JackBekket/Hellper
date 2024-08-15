package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/bot/dialog"
	"github.com/JackBekket/hellper/lib/bot/env"
	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

/*
type AdminData struct {
	ID     int64
	GPTKey string
	//localhost_password string
}
*/

func main() {

	_= godotenv.Load()
	api_token := os.Getenv("OPENAI_API_KEY")	// this is not openai key actually, it's local key for localai
	//conn_pg_link := os.Getenv("PG_LINK")

	/*
	err := env.Load()
	if err != nil {
		log.Panicf("could not load env from: %v", err)
	}

	token, err := env.LoadTGToken()
	if err != nil {
		log.Panic(err)
	}
	log.Println("TG token is: ", token)

	adminData := env.LoadAdminData()
	//local_access_pwd:= env.LoadLocalPD()
	ai_endpoint := env.LoadLocalAI_Endpoint()
	log.Println("ai endpoint is: ", ai_endpoint)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("tg token missing: %v\n", err)
	}
	*/


	//token := api_token
	token := os.Getenv("TG_KEY")
	log.Println("TG token is: ", token)

	admin_key := api_token
	admin_id := os.Getenv("ADMIN_ID")
	id, err := strconv.ParseInt(admin_id, 0, 64)
			if err != nil {
				log.Printf("admin id error parse: %s", admin_id)
			}
	ai_endpoint := os.Getenv("AI_ENDPOINT")
	log.Println("ai endpoint is: ", ai_endpoint)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("tg token missing: %v\n", err)
	}

	adminData := make(map[string]env.AdminData)
	adminData["ADMIN_ID"] = env.AdminData{
		ID:     id,
		GPTKey: admin_key,
	}






	// init database and commander
	usersDatabase := database.UsersMap
	ctx := context.Background()
	comm := command.NewCommander(bot, usersDatabase, ctx)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updateHandler := 
	updates := bot.GetUpdatesChan(u)
	go dialog.HandleUpdates(updates,bot,*comm)
	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	/*
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

	 } // range updates
	 */ 

	} // end of main func