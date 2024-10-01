package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Render OpenAI Model Menu with Inline Keyboard
func (c *Commander) RenderModelMenuOAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("gpt-3.5", "gpt-3.5"),
			tgbotapi.NewInlineKeyboardButtonData("gpt-4", "gpt-4"),
		),
	)
	c.bot.Send(msg)
}

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

// Render Various AI Models Menu with Inline Keyboard
func (c *Commander) RenderModelMenuVAI(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, msgTemplates["case1"])
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("deepseek-coder-6b-instruct", "deepseek-coder-6b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("wizard-uncensored-code-34b", "wizard-uncensored-code-34b"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("tiger-gemma-9b-v1-i1", "tiger-gemma-9b-v1-i1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("big-tiger-gemma-27b-v1", "big-tiger-gemma-27b-v1"),
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
