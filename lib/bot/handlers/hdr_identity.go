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
		DialogStatus: status_AIModelSelectionKeyboard,
		Admin:        false,
		AiSession: database.AiSession{
			Base_url: h.config.AI_endpoint,
		},
	}

	h.cache.SetUser(chatID, user)

	msg := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   msg_Hello,
	}

	_, err := tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
	}
}

// download user data from database into cashe
func restoreUserSessionFromDB(ds *database.Service, chatID int64, username string) (*database.User, error) {
	ai_session, err := ds.GetSession(chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session: %w", err)
	}

	gptKey, err := ds.GetToken(chatID, 1)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user's API key : %w", err)
	}

	history, err := ds.GetHistory(chatID, ai_session.Endpoint.ID, chatID, chatID, *ai_session.Model)
	if err != nil {
		return nil, fmt.Errorf("error loading user's history : %w", err)
	}

	user := &database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: status_StartDialogSequence,
		Admin:        false,
		AiSession: database.AiSession{
			GptModel: *ai_session.Model,
			Base_url: ai_session.Endpoint.URL,
			GptKey:   gptKey,
			DialogThread: database.ChatSessionGraph{
				ConversationBuffer: history,
			},
		},
	}
	return user, nil

}
func recoverUserAfterDrop(ds *database.Service, chatID int64, username string, ai_endpoint string) (*database.User, error) {

	log.Info().
		Int64("chat_id", chatID).
		Str("username", username).
		Msg("Restoring a registered user")

	user := &database.User{
		ID:           chatID,
		Username:     username,
		DialogStatus: status_AIModelSelectionChoice,
		Admin:        false,
		AiSession: database.AiSession{
			Base_url: ai_endpoint,
		},
	}

	gptKey, err := ds.GetToken(chatID, 1)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error retrieving GPT KEY")
		log.Warn().Int64("chat_id", chatID).Msg("User redirected to registration handler due to GPT KEY error")
		return &database.User{}, err
	}

	user.AiSession.GptKey = gptKey
	return user, nil
}
