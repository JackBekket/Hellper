package command

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	db "github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	"github.com/JackBekket/hellper/lib/localai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type contextKey string

const UserKey contextKey = "user"

// Message:	case0 - "Input your openAI API key. It can be created at https://platform.openai.com/accousernamet/api-keys".
//
//	DialogStatus 2 -> 3
func (c *Commander) InputYourAPIKey(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]

	msg := tgbotapi.NewMessage(
		user.ID,
		msgTemplates["case0"],
	)
	c.bot.Send(msg)

	user.DialogStatus = 3
	db.UsersMap[chatID] = user
}

// DialogStatus 0 - > 1
func (c *Commander) ChooseNetwork(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]
	c.HelpCommandMessage(updateMessage)
	// render menu
	msg := tgbotapi.NewMessage(user.ID, msgTemplates["ch_network"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("openai"),
			tgbotapi.NewKeyboardButton("localai"),
			tgbotapi.NewKeyboardButton("vastai")),
	)
	c.bot.Send(msg)

	user.DialogStatus = 1      // this is output dialog status
	db.UsersMap[chatID] = user // commit changes

}

// Dialog status 1 -> 2
func (c *Commander) HandleNetworkChoose(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	network := updateMessage.Text
	user := db.UsersMap[chatID]
	switch network {
	case "openai":

		user.Network = network
		user.AiSession.AI_Type = 0
		user.DialogStatus = 2
		db.UsersMap[chatID] = user
		c.InputYourAPIKey(updateMessage)
	case "localai":

		user.Network = network
		user.AiSession.AI_Type = 1
		user.DialogStatus = 2
		db.UsersMap[chatID] = user
		c.InputYourAPIKey(updateMessage)
	case "vastai":
		user.Network = network
		user.AiSession.AI_Type = 2
		user.DialogStatus = 2
		db.UsersMap[chatID] = user
		c.InputYourAPIKey(updateMessage)
	default:
		c.WrongNetwork(updateMessage)
	}

}

// update Dialog_Status 3 -> 4
func (c *Commander) ChooseModel(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	gptKey := updateMessage.Text // handling previouse message
	user := db.UsersMap[chatID]
	network := user.Network

	// I can't validate key at this stage. The only way to validate key is to send test sequence (see case 3)
	// Since this part is oftenly get an usernamecaught exeption, we debug what user input as key. It's bad, I know, but usernametil we got key validation we need this part.
	log.Println("Key promt: ", gptKey)
	user.AiSession.GptKey = gptKey // store key in memory

	switch network {
	case "localai":
		c.RenderModelMenuLAI(chatID)
		user.DialogStatus = 4
		db.UsersMap[chatID] = user

	case "openai":
		c.RenderModelMenuOAI(chatID)
		user.DialogStatus = 4
		db.UsersMap[chatID] = user

	case "vastai":
		c.RenderModelMenuVAI(chatID)
		user.DialogStatus = 4
		db.UsersMap[chatID] = user
	default:
		c.WrongNetwork(updateMessage)
	}
}

// DialogStatus 4 -> 5
func (c *Commander) HandleModelChoose(updateMessage *tgbotapi.CallbackQuery) {
	chatID := updateMessage.From.ID
	messageID := updateMessage.Message.MessageID
	model_name := updateMessage.Data
	user := db.UsersMap[chatID]
	network := user.Network
	switch network {
	case "localai":
		switch model_name {
		case "wizard-uncensored-13b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "wizard-uncensored-30b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "deepseek-coder-6b-instruct":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "tiger-gemma-9b-v1-i1":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "wizard-uncensored-code-34b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user

		}

	case "openai":
		switch model_name {
		case "gpt-3.5":
			model_name = "gpt-3.5-turbo"
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "gpt-4":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		}
	case "vastai":
		switch model_name {
		case "wizard-uncensored-13b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "wizard-uncensored-30b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "deepseek-coder-6b-instruct":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "tiger-gemma-9b-v1-i1":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "big-tiger-gemma-27b-v1":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		case "wizard-uncensored-code-34b":
			c.attachModel(model_name, chatID)
			user.AiSession.GptModel = model_name
			c.RenderLanguage(chatID)

			user.DialogStatus = 5
			db.UsersMap[chatID] = user
		}

	}
	callbackResponse := tgbotapi.NewCallback(updateMessage.ID, "🐈💨")
	c.bot.Send(callbackResponse)

	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	c.bot.Send(deleteMsg)

}

// low level attach model name to user profile
func (c *Commander) attachModel(model_name string, chatID int64) {
	fmt.Println(model_name)
	// TODO: Write down user choise
	log.Printf("Model selected: %s\n", model_name)

	user := db.UsersMap[chatID]

	modelName := model_name
	user.AiSession.GptModel = modelName
	msg := tgbotapi.NewMessage(user.ID, "your session model: "+modelName)
	c.bot.Send(msg)
	db.UsersMap[chatID] = user
}

// internal for attach api key to a user
func (c *Commander) AttachKey(gpt_key string, chatID int64) {
	log.Println("Key promt: ", gpt_key)
	user := db.UsersMap[chatID]
	user.AiSession.GptKey = gpt_key // store key in memory
	db.UsersMap[chatID] = user
}

// Dangerouse! NOTE -- probably work only internal
func (c *Commander) ChangeDialogStatus(chatID int64, ds int8) {
	user := db.UsersMap[chatID]
	old_status := user.DialogStatus
	log.Println("dialog status changed, old status is ", old_status)
	log.Println("new status is ", ds)
	user.DialogStatus = ds
}

func (c *Commander) WrongResponse(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]

	msg := tgbotapi.NewMessage(user.ID, "Please use provided keyboard")
	c.bot.Send(msg)

}

// update Dialog_Status = 0
func (c *Commander) WrongNetwork(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]

	msg := tgbotapi.NewMessage(user.ID, "type openai or localai")
	c.bot.Send(msg)

	user.DialogStatus = 0
	db.UsersMap[chatID] = user
}

// update update Dialog_Status 5 -> 6
func (c *Commander) ConnectingToAiWithLanguage(updateMessage *tgbotapi.CallbackQuery, ai_endpoint string) {
	_ = godotenv.Load()
	messageID := updateMessage.Message.MessageID
	chatID := updateMessage.From.ID
	language := updateMessage.Data
	user := db.UsersMap[chatID]
	log.Println("check gpt key exist:", user.AiSession.GptKey)

	network := user.Network

	msg := tgbotapi.NewMessage(user.ID, "connecting to ai node")
	c.bot.Send(msg)

	ctx := context.WithValue(c.ctx, "user", user)

	if network == "localai" {
		log.Println("network: ", network)
		if ai_endpoint == "" {
			ai_endpoint = os.Getenv("AI_ENDPOINT")
		}
		log.Println("local-ai endpoint is: ", ai_endpoint)
		go langchain.SetupSequenceWithKey(c.bot, user, language, ctx, ai_endpoint) //local-ai
	} else if network == "vastai" {
		log.Println("network: ", network)
		ai_endpoint := os.Getenv("VASTAI_ENDPOINT")
		log.Println("vast-ai endpoint is: ", ai_endpoint)
		go langchain.SetupSequenceWithKey(c.bot, user, language, ctx, ai_endpoint) //vast-ai
	} else {
		log.Println("network: ", network)
		go langchain.SetupSequenceWithKey(c.bot, user, language, ctx, "") //openai
	}

	callbackResponse := tgbotapi.NewCallback(updateMessage.ID, "🐈💨")
	c.bot.Send(callbackResponse)

	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	c.bot.Send(deleteMsg)

}

// Generates an image with the /image command.
//
// Generates and sends text to the user. This is *main loop*
//
// update Dialog_Status 6 -> 6 (loop),
func (c *Commander) DialogSequence(updateMessage *tgbotapi.Message, ai_endpoint string) {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]

	if updateMessage != nil {

		promt := updateMessage.Text
		ctx := context.WithValue(c.ctx, "user", user)
		go langchain.StartDialogSequence(c.bot, chatID, promt, ctx, ai_endpoint)
	}
}

// stable diffusion
func (c *Commander) GenerateNewImageLAI_SD(promt string, chatID int64, bot *tgbotapi.BotAPI) {
	size := "256x256"
	filepath, err := localai.GenerateImageStableDissusion(promt, size)
	if err != nil {
		//return nil, err
		log.Println(err)
	}
	log.Println("url_path: ", filepath)
	sendImage(bot, chatID, filepath)
}

func sendImage(bot *tgbotapi.BotAPI, chatID int64, path string) {
	// Prepare a photo message
	fileName := transformURL(path)
	log.Println("local file name: ", fileName)

	telegraphLink := localai.UploadToTelegraph(fileName)
	log.Println("uploaded to telegraph successfully, link is: ", telegraphLink)

	// Path to the image/file locally
	// filePath := "/path/to/image.png" + local_path
	/*
			 // Creating a LocalFile object from the local path
			photoBytes, err := ioutil.ReadFile(filePath)
			if err != nil {
		    	log.Println(err)
						}
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
				}
	*/
	//message, err := bot.Send(tgbotapi.NewPhotoUpload(int64(chatID), photoFileBytes))
	/* photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(local_path))
	if _, err := bot.Send(photo); err != nil {
	log.Fatalln(err)
	} */
	msg := tgbotapi.NewMessage(chatID, telegraphLink)
	bot.Send(msg)
}

func transformURL(inputURL string) string {
	// Replace "http://localhost:8080" with "/tmp" using strings.Replace
	parsedURL, _ := url.Parse(inputURL)

	// Use path.Base to get the filename from the URL path
	fileName := path.Base(parsedURL.Path)
	return fileName
}

func (c *Commander) GetUsersDb() map[int64]db.User {
	data_base := db.UsersMap
	return data_base
}

func (c *Commander) GetUser(id int64) db.User {
	user := db.UsersMap[id]
	return user
}
