package handlers

import (
	"github.com/go-telegram/bot/models"
)

// Render AI models menu with Inline Keyboard
func renderAIModelsInlineKeyboard(aiModelsList []string) models.InlineKeyboardMarkup {
	buttons := [][]models.InlineKeyboardButton{}
	for _, model := range aiModelsList {
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

// Render Language menu with Inline Keyboard
func renderLangInlineKeyboard() models.InlineKeyboardMarkup {
	return models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "English", CallbackData: "English"},
				{Text: "Russian", CallbackData: "Russian"},
			},
		},
	}
}
