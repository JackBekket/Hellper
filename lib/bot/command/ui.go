package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Render LLaMA-based Model Menu with Inline Keyboard
func (c *Commander) RenderModelMenuLAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("wizard-uncensored-13b", "wizard-uncensored-13b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("wizard-uncensored-30b", "wizard-uncensored-30b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("tiger-gemma-9b-v1-i1", "tiger-gemma-9b-v1-i1"),
		),
	)
	c.bot.Send(msg)
}


// Render Language Menu with Inline Keyboard
func (c *Commander) RenderLanguage(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Choose a language or send 'Hello' in your desired language.")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("English", "English"),
			tgbotapi.NewInlineKeyboardButtonData("Russian", "Russian"),
		),
	)

	c.bot.Send(msg)
}
