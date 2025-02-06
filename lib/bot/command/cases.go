package command

import (
	"context"
	"fmt"
	"log"
	"strings"

	db "github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/langchain"
	"github.com/JackBekket/hellper/lib/localai"
	stt "github.com/JackBekket/hellper/lib/localai/audioRecognition"
	imgrec "github.com/JackBekket/hellper/lib/localai/imageRecognition"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type contextKey string

const UserKey contextKey = "user"

// Message:	case0 - "Input your openAI API key. It can be created at https://platform.openai.com/accousernamet/api-keys".
//
//	DialogStatus 2 -> 3
func (c *Commander) InputYourAPIKey(updateMessage *tgbotapi.Message) {
	updateMessage.Text = strings.ReplaceAll(updateMessage.Text, " ", "")
	chatID := updateMessage.Chat.ID
	user := db.UsersMap[chatID]

	msg := tgbotapi.NewMessage(
		user.ID,
		msgTemplates["case0"],
	)
	c.bot.Send(msg)

	user.DialogStatus = 3
	db.UsersMap[chatID] = user
}

// update Dialog_Status 3 -> 4
func (c *Commander) ChooseModel(updateMessage *tgbotapi.Message) {
	updateMessage.Text = strings.TrimSpace(updateMessage.Text)
	chatID := updateMessage.Chat.ID
	gptKey := updateMessage.Text // handling previouse message
	user := db.UsersMap[chatID]

	// I can't validate key at this stage. The only way to validate key is to send test sequence (see case 3)
	// Since this part is oftenly get an usernamecaught exeption, we debug what user input as key. It's bad, I know, but usernametil we got key validation we need this part.
	log.Println("Key promt: ", gptKey)
	user.AiSession.GptKey = gptKey // store key in memory
	c.RenderModelMenuLAI(chatID)
	user.DialogStatus = 4
	db.UsersMap[chatID] = user
}

// DialogStatus 4 -> 5
func (c *Commander) HandleModelChoose(updateMessage *tgbotapi.CallbackQuery) {
	chatID := updateMessage.Message.Chat.ID
	messageID := updateMessage.Message.MessageID
	model_name := updateMessage.Data
	user := db.UsersMap[chatID]
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

	callbackResponse := tgbotapi.NewCallback(updateMessage.ID, "🐈💨")
	c.bot.Send(callbackResponse)

	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	c.bot.Send(deleteMsg)

}

// low level attach model name to user profile
func (c *Commander) attachModel(model_name string, chatID int64) {
	fmt.Println(model_name)
	// TODO: Write down user choice
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
	chatID := updateMessage.Chat.ID
	user := db.UsersMap[chatID]

	msg := tgbotapi.NewMessage(user.ID, "Please use provided keyboard")
	c.bot.Send(msg)

}

// update update Dialog_Status 5 -> 6
func (c *Commander) ConnectingToAiWithLanguage(updateMessage *tgbotapi.CallbackQuery, ai_endpoint string) {
	_ = godotenv.Load()
	messageID := updateMessage.Message.MessageID
	chatID := updateMessage.Message.Chat.ID
	language := updateMessage.Data
	user := db.UsersMap[chatID]
	log.Println("check gpt key exist:", user.AiSession.GptKey)

	//network := user.Network

	msg := tgbotapi.NewMessage(user.ID, "connecting to ai node")
	c.bot.Send(msg)

	ctx := context.WithValue(c.ctx, "user", user)
	//ai_endpoint = os.Getenv("AI_ENDPOINT")
	log.Println("local-ai endpoint is: ", ai_endpoint)
	go langchain.SetupSequenceWithKey(c.bot, user, language, ctx, ai_endpoint) //local-ai

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
	chatID := updateMessage.Chat.ID
	user := db.UsersMap[chatID]

	if updateMessage.Command() != "" {
		HandleCommands(updateMessage, c)
	} else {

		if updateMessage != nil {

			if updateMessage.Text != "" && updateMessage.Photo == nil {
				promt := updateMessage.Text
				ctx := context.WithValue(c.ctx, "user", user)
				go langchain.StartDialogSequence(c.bot, chatID, promt, ctx, ai_endpoint)
			} else if updateMessage.Voice != nil {
				voicePath, err := stt.HandleVoiceMessage(updateMessage, *c.bot)
				if err != nil {
					log.Println(err)
				}
				url, model := stt.GetEnvsForSST()
				transcription, err := localai.TranscribeWhisper(url, model, voicePath)
				if err != nil {
					log.Println(err)
				}
				msg := tgbotapi.NewMessage(chatID, transcription)
				c.bot.Send(msg)
				DeleteFile(voicePath)
			} else if updateMessage.Photo != nil {
				response, err := imgrec.RecognizeImage(c.bot, updateMessage)
				if err != nil {
					log.Println(err)
				}
				msg := tgbotapi.NewMessage(chatID, response)
				c.bot.Send(msg)
			}
		}
	}
}

func (c *Commander) GetUsersDb() map[int64]db.User {
	data_base := db.UsersMap
	return data_base
}

func (c *Commander) GetUser(id int64) db.User {
	user := db.UsersMap[id]
	return user
}
