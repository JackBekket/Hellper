package handlers

import (
	"context"
	"fmt"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// old func name: AddNewUserToMap.
// Creates a session for a new user. Sends a welcome message.
func (h *handlers) handleNewUserRegistration(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	chatID := update.Message.Chat.ID
	h.cache.SetUser(chatID, database.User{
		ID:           chatID,
		Username:     update.Message.From.Username,
		DialogStatus: statusAuthMethodCallback,
		Admin:        false,
	})

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        msgHello,
		ReplyMarkup: renderAIServicesInlineKeyboard(),
	}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

// download user data from database into cashe
func restoreUserSessionFromDB(ds *database.Service, chatID int64, username string) (database.User, error) {
	aiSession, err := ds.GetAISession(chatID)
	if err != nil {
		return database.User{}, fmt.Errorf("failed to retrieve session: %w", err)
	}

	aiToken, err := ds.GetToken(chatID, aiSession.AIProvider.AuthMethod)
	if err != nil {
		return database.User{}, fmt.Errorf("error retrieving user's API key : %w", err)
	}

	history, err := ds.GetHistory(chatID, aiSession.AIProvider.ID, chatID, chatID, aiSession.Model)
	if err != nil {
		return database.User{}, fmt.Errorf("error loading user's history : %w", err)
	}

	return database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: statusStartDialogSequence,
		Admin:        false,
		AiSession: database.AiSession{
			GptModel:     aiSession.Model,
			BaseURL:      aiSession.AIProvider.BaseURL,
			AIToken:      aiToken,
			AuthMethod:   aiSession.AIProvider.AuthMethod,
			ProviderID:   aiSession.AIProvider.ID,
			ProviderName: aiSession.AIProvider.Name,
			DialogThread: database.ChatSessionGraph{
				ConversationBuffer: history,
			},
		},
	}, nil
}

// func recoverUserAfterDrop(ds *database.Service, chatID int64, username string) (database.User, error) {
// 	log.Info().Int64("chat_id", chatID).Str("username", username).Msg("Restoring a registered user")

// 	auth, err := ds.GetUserTokensWithAuthMethods(chatID)
// 	if err != nil {
// 		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving GPT KEY")
// 		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to GPT KEY error")
// 		return database.User{}, err
// 	}

// 	if len(auth) > 1 {

// 	}

// 	return database.User{
// 		ID:           chatID,
// 		Username:     username,
// 		DialogStatus: statusAuthMethodCallback,
// 		Admin:        false,
// 		AiSession:    database.AiSession{},
// 	}, nil
// }
