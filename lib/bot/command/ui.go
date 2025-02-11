package command

import (
	"fmt"

	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-telegram/bot/models"
)

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


// Get list of models from endpoint and render them
func (c *Commander) RenderModels(chatID int64,db_service *database.Service, user database.User) {
	models_list,err := c.getModels(db_service,user)
	if err != nil {
		e_txt := fmt.Sprintf(err.Error())
		msg := tgbotapi.NewMessage(chatID, e_txt)
		c.bot.Send(msg)
	}
	msg := tgbotapi.NewMessage(chatID, "Choose model")
	msg.ReplyMarkup = CreateModelsMarkup(models_list)
	
	c.bot.Send(msg)
}

func (c *Commander) getModels(db_service *database.Service, user database.User) ([]string,error){
	ds := db_service
	api_key := user.AiSession.GptKey
	url := user.AiSession.Base_url
	list,err := ds.GetModelsList(url,api_key)
	if err != nil {
		return nil, err
	}
	return list, nil
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

func CreateModelsMarkup(llmModels []string) models.InlineKeyboardMarkup {
	buttons := [][]models.InlineKeyboardButton{}
	for _, model := range llmModels {
		buttons = append(buttons, []models.InlineKeyboardButton{
			{
				Text:         model,
				CallbackData: "model_" + model,
			},
		})
	}

	return models.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}
