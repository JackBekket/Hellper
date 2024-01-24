package command

import (
	"log"

	"github.com/JackBekket/telegram-gpt/internal/openaibot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
)

// Message:	case0 - "Input your openAI API key. It can be created at https://platform.openai.com/accousernamet/api-keys".
//
//	update DialogStatus = 1
func (c *Commander) InputYourAPIKey(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	msg := tgbotapi.NewMessage(
		user.ID,
		msgTemplates["case0"],
	)
	c.bot.Send(msg)

	user.DialogStatus = 6
	c.usersDb[chatID] = user
}


// TODO:
// 1. figure out which dialog status is inputed here and which outputed
// 2. figure out how to place it in main menu
// 3. -----
// 4. this function should render choose network menu
func (c *Commander) ChooseNetwork(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	// validate message from previouse side
	// TODO
	gptKey := updateMessage.Text
	//c.AttachKey(gptKey,chatID)		// save gpt key
	user.AiSession.GptKey = gptKey


	// render menu
	msg := tgbotapi.NewMessage(user.ID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("openai"),
			tgbotapi.NewKeyboardButton("localai")),
		
	)
	c.bot.Send(msg)

	user.DialogStatus = 7	 // TODO: change it  // this is output dialog status
	c.usersDb[chatID] = user // commit changes

}



// Message: case1 - "Choose model to use. GPT3 is for text-based tasks, Codex for codegeneration.".
//
//	update Dialog_Status = 2
func (c *Commander) ChooseModel(updateMessage *tgbotapi.Message) {
	ok := false
	chatID := updateMessage.From.ID
	//gptKey := updateMessage.Text
	user := c.usersDb[chatID]
	//log.Println("chooseModel function worked")

	// I can't validate key at this stage. The only way to validate key is to send test sequence (see case 3)
	// Since this part is oftenly get an usernamecaught exeption, we debug what user input as key. It's bad, I know, but usernametil we got key validation we need this part.
	//log.Println("Key promt: ", gptKey)
	//user.AiSession.GptKey = gptKey // store key in memory

	//TODO get user network and check if it is openai or localai
	network := user.Network

	//TODO render different menu with different models for openai and localhost
	if (network == "openai") {
		c.RenderModelMenuOAI(chatID)
		ok = true
	}
	if (network == "localai") {
		c.RenderModelMenuLAI(chatID)
		ok = true
	}

	if (ok) {
		user.DialogStatus = 2
		c.usersDb[chatID] = user
	} else {
		user.DialogStatus = 7
		c.usersDb[chatID] = user
	}

}

func (c *Commander) HandleModelChoose(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	model_name := updateMessage.Text
	user := c.usersDb[chatID]
	switch model_name {
	case "GPT-3.5":
		c.ModelGPT3DOT5(updateMessage)
	case "wizard-uncensored-13b":
	 	c.attachModel(model_name,chatID)
		//c.ChangeDialogStatus(chatID,3)
		user.AiSession.GptModel = model_name
		user.DialogStatus = 3
		c.usersDb[chatID] = user
	case "wizard-uncensored-30b":
		c.attachModel(model_name,chatID)
		//c.ChangeDialogStatus(chatID,3)
		user.AiSession.GptModel = model_name
		user.DialogStatus = 3
		c.usersDb[chatID] = user
	default:
		c.WrongModel(updateMessage)
	}

}

// Message: "Choose language. If you have different languages then listed, then just send 'Hello' at your desired language".
//
//	update Dialog_Status = 3
func (c *Commander) ModelGPT3DOT5(updateMessage *tgbotapi.Message) {
	// TODO: Write down user choise
	log.Printf("Model selected: %s\n", updateMessage.Text)

	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	modelName := openai.GPT3Dot5Turbo // gpt-3.5
	user.AiSession.GptModel = modelName
	msg := tgbotapi.NewMessage(user.ID, "your session model: "+modelName)
	c.bot.Send(msg)

	msg = tgbotapi.NewMessage(user.ID, "Choose a language or send 'Hello' in your desired language.")
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("English"),
			tgbotapi.NewKeyboardButton("Russian")),
	)
	c.bot.Send(msg)

	user.DialogStatus = 3
	c.usersDb[chatID] = user
}

// Message: "your session model: Codex".
//
//	update Dialog_Status = 4
func (c *Commander) ModelCodex(updateMessage *tgbotapi.Message) {
	log.Printf("Model selected: %s\n", updateMessage.Text)
	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	modelCodex := openai.CodexCodeDavinci002
	user.AiSession.GptModel = modelCodex

	msg := tgbotapi.NewMessage(user.ID, "your session model :"+modelCodex)
	c.bot.Send(msg)

	msg = tgbotapi.NewMessage(user.ID, msgTemplates["codex_help"])
	msg.ParseMode = "MARKDOWN"
	c.bot.Send(msg)

	user.DialogStatus = 4
	c.usersDb[chatID] = user
}

// ModelGPT and ModelLL codes are the same.
// TODO
func (c *Commander) ModelGPT4(updateMessage *tgbotapi.Message) {
	// TODO: Write down user choise
	log.Printf("Model selected: %s\n", updateMessage.Text)

	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	modelName := openai.GPT4 // gpt-4
	user.AiSession.GptModel = modelName
	msg := tgbotapi.NewMessage(user.ID, "your session model: "+modelName)
	c.bot.Send(msg)

	msg = tgbotapi.NewMessage(user.ID, "Choose a language or send 'Hello' in your desired language.")
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("English"),
			tgbotapi.NewKeyboardButton("Russian")),
	)
	c.bot.Send(msg)

	user.DialogStatus = 3
	c.usersDb[chatID] = user
}

// low level attach model name to user profile
func (c *Commander) attachModel(model_name string, chatID int64) {
		// TODO: Write down user choise
		log.Printf("Model selected: %s\n", model_name)

		
		user := c.usersDb[chatID]
	
		modelName := model_name 
		user.AiSession.GptModel = modelName
		msg := tgbotapi.NewMessage(user.ID, "your session model: "+modelName)
		c.bot.Send(msg)
		c.usersDb[chatID] = user
}

func (c *Commander) AttachNetworkAndUpdDialog(network string, chatID int64) {
	c.AttachNetwork(network,chatID)
	user := c.usersDb[chatID]

	msg := tgbotapi.NewMessage(user.ID, "your session network: "+network)
	c.bot.Send(msg)

	user.DialogStatus = 1
	c.usersDb[chatID] = user
}

// internal for attach api key to a user
func (c *Commander) AttachKey(gpt_key string, chatID int64) {
	log.Println("Key promt: ", gpt_key)
	user := c.usersDb[chatID]
	user.AiSession.GptKey = gpt_key // store key in memory
	c.usersDb[chatID] = user
}

func (c *Commander) AttachNetwork(network string, chatID int64) {
	//chatID := updateMessage.From.ID
	user := c.usersDb[chatID]
	user.Network = network
	c.usersDb[chatID] = user
	log.Println("network attached")
}


// Dangerouse! NOTE -- probably work only internal
func (c *Commander) ChangeDialogStatus(chatID int64, ds int8) {
	user := c.usersDb[chatID]
	old_status := user.DialogStatus
	log.Println("dialog status changed, old status is ", old_status)
	log.Println("new status is ", ds)
	user.DialogStatus = ds
}



func (c *Commander) RenderModelMenuOAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("GPT-3.5")),
		//tgbotapi.NewKeyboardButton("GPT-4"),
	)
	c.bot.Send(msg)
}


func (c *Commander) RenderModelMenuLAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("wizard-uncensored-13b"),
		tgbotapi.NewKeyboardButton("wizard-uncensored-30b")),
	)
	c.bot.Send(msg)
}

// update Dialog_Status = 2
func (c *Commander) WrongModel(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]

	msg := tgbotapi.NewMessage(user.ID, "type GPT-3.5")
	c.bot.Send(msg)

	user.DialogStatus = 2
	c.usersDb[chatID] = user
}

// Message: "connecting to openAI"
//
// update update Dialog_Status = 4, for model GPT-3.5
func (c *Commander) ConnectingToOpenAiWithLanguage(updateMessage *tgbotapi.Message, lpwd string) {
	chatID := updateMessage.From.ID
	language := updateMessage.Text
	user := c.usersDb[chatID]
	log.Println("check gpt key exist:", user.AiSession.GptKey)

	msg := tgbotapi.NewMessage(user.ID, "connecting to node")
	c.bot.Send(msg)
	
	go openaibot.SetupSequenceWithKey(c.bot, user, language, c.ctx, lpwd)
}

// Generates an image with the /image command.
//
// Generates and sends text to the user.
//
// update Dialog_Status = 4, for model GPT-3.5
func (c *Commander) DialogSequence(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	user := c.usersDb[chatID]
	switch updateMessage.Command() {
	case "image":
		msg := tgbotapi.NewMessage(user.ID, "Image link generation...")
		c.bot.Send(msg)

		promt := updateMessage.CommandArguments()
		log.Printf("Command /image arg: %s\n", promt)
		go openaibot.StartImageSequence(c.bot, updateMessage, chatID, promt, c.ctx)

	default:
		promt := updateMessage.Text
		go openaibot.StartDialogSequence(c.bot, chatID, promt, c.ctx)
	}
}

// Generates and sends code to the user.
//
// At the moment there is no access to the Codex.
func (c *Commander) CodexSequence(updateMessage *tgbotapi.Message) {
	chatID := updateMessage.From.ID
	promt := updateMessage.Text
	go openaibot.StartCodexSequence(c.bot, chatID, promt, c.ctx)
	//user.DialogStatus = 0
	//userDatabase[chatID] = user
}



