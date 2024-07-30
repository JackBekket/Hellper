package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) RenderModelMenuOAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("gpt-3.5"),
		tgbotapi.NewKeyboardButton("gpt-4")),
	)
	c.bot.Send(msg)
}

func (c *Commander) RenderModelMenuLAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("wizard-uncensored-13b")),
		//tgbotapi.NewKeyboardButton("code-13b")),
	)

	c.bot.Send(msg)
}


// models choose menu
func (c *Commander) RenderModelMenuVAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("wizard-uncensored-13b"),
		tgbotapi.NewKeyboardButton("wizard-uncensored-30b")),
		
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("qwen14b"),
		),
	
	)

	c.bot.Send(msg)



}





// render language menu
func (c *Commander) RenderLanguage(chat_id int64) {
	chatID := chat_id
	//user := c.usersDb[chatID]

	msg := tgbotapi.NewMessage(chatID, "Choose a language or send 'Hello' in your desired language.")
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("English"),
			tgbotapi.NewKeyboardButton("Russian")),
	)
	c.bot.Send(msg)

}