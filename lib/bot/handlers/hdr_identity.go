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

	user := database.User{
		ID:           chatID,
		Username:     update.Message.From.Username,
		DialogStatus: statusAIModelSelectionKeyboardForNewUser,
		Admin:        false,
	}

	h.cache.SetUser(chatID, user)

	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msgHello,
	}
	// послать клавиатуру
	_, err := tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

// download user data from database into cashe
func restoreUserSessionFromDB(ds *database.Service, chatID int64, username string) (database.User, error) {
	aiSession, err := ds.GetSession(chatID)
	if err != nil {
		return database.User{}, fmt.Errorf("failed to retrieve session: %w", err)
	}

	LocalAIToken, err := ds.GetToken(chatID, 1)
	if err != nil {
		return database.User{}, fmt.Errorf("error retrieving user's API key : %w", err)
	}

	history, err := ds.GetHistory(chatID, aiSession.AIProvider.ID, chatID, chatID, aiSession.Model)
	if err != nil {
		return database.User{}, fmt.Errorf("error loading user's history : %w", err)
	}

	user := database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: statusStartDialogSequence,
		Admin:        false,
		AiSession: database.AiSession{
			GptModel:     aiSession.Model,
			Base_url:     aiSession.AIProvider.BaseURL,
			LocalAIToken: LocalAIToken,
			DialogThread: database.ChatSessionGraph{
				ConversationBuffer: history,
			},
		},
	}
	return user, nil

}
func recoverUserAfterDrop(ds *database.Service, chatID int64, username string, BaseURL string) (database.User, error) {

	log.Info().
		Int64("chat_id", chatID).
		Str("username", username).
		Msg("Restoring a registered user")

	LocalAIToken, err := ds.GetToken(chatID, 1)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving GPT KEY")
		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to GPT KEY error")
		return database.User{}, err
	}
	user := database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: statusAIModelSelectionChoice,
		Admin:        false,
		AiSession: database.AiSession{
			Base_url:     BaseURL,
			LocalAIToken: LocalAIToken,
		},
	}

	return user, nil
}
