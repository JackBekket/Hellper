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
			tgbotapi.NewKeyboardButton("wizard-uncensored-13b"),
		tgbotapi.NewKeyboardButton("wizard-uncensored-30b"),
		tgbotapi.NewKeyboardButton("tiger-gemma-9b-v1-i1"),
		))

	c.bot.Send(msg)
}


// models choose menu
func (c *Commander) RenderModelMenuVAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("deepseek-coder-6b-instruct"),
		tgbotapi.NewKeyboardButton("wizard-uncensored-code-34b")),
		
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("tiger-gemma-9b-v1-i1"),
			tgbotapi.NewKeyboardButton("big-tiger-gemma-27b-v1"),
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