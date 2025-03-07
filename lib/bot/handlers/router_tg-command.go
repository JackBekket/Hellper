package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/embeddings"
	"github.com/JackBekket/hellper/lib/langchain"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Function to extract the command and argument.
// Also removes the bot's name if the message was sent in a group chat.
// At the moment, only one argument is allowed
func extractCommandAndArg(msg string) (string, string) {
	msg = strings.TrimSpace(msg)
	if len(msg) == 0 || msg[0] != '/' {
		return "", ""
	}

	parts := strings.Fields(msg)
	command := strings.Split(parts[0], "@")[0]
	arg := strings.TrimSpace(strings.Join(parts[1:], " "))
	return command, arg
}

// Router for tg bot command handlers
func (h *handlers) cmdRouter(ctx context.Context, tgb *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	command, arg := extractCommandAndArg(update.Message.Text)

	log.Info().Int64("chat_id", chatID).Str("command", command).Str("arg", arg).Msg("processing command")

	switch command {
	case "/image":
		h.cmdGenerateImage(ctx, tgb, chatID, arg)
	case "/reload":
		h.cmdReload(ctx, tgb, chatID)
	case "/clear":
		h.cmdClear(ctx, tgb, chatID)
	case "/purge":
		h.cmdPurge(ctx, tgb, chatID)
	case "/drop":
		h.cmdDrop(ctx, tgb, chatID)
	case "/help":
		h.cmdHelp(ctx, tgb, chatID)
	case "/search_doc":
		h.cmdSearchDoc(ctx, tgb, chatID, arg)
	case "/instruct":
		h.cmdInstruct(ctx, tgb, chatID, arg)
	case "/usage":
		h.cmdUsage(ctx, tgb, chatID)
	case "/helper":
		h.cmdHelper(ctx, tgb, chatID)
	case "/setContext":
		h.cmdSetContext(ctx, tgb, chatID, arg)
	case "/clearContext":
		h.cmdClearContext(ctx, tgb, chatID)
	}
}

func (h *handlers) cmdReload(ctx context.Context, tgb *bot.Bot, chatID int64) {
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Reloading session..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
	log.Info().Int64("chat_id", chatID).Msg("User reloaded the session in bot")
	h.cache.DeleteUser(chatID)
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Done. Type any key..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

func (h *handlers) cmdClear(ctx context.Context, tgb *bot.Bot, chatID int64) {
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Deleting dialog thread from database..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.FlushMemory(h.dbService)
	h.cache.DeleteUser(chatID)
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Done. Type any key..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

// Completely removes all user records from the storage
func (h *handlers) cmdPurge(ctx context.Context, tgb *bot.Bot, chatID int64) {
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Deleting all user data from database and restarting session..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.Kill(h.dbService)
	h.cache.DeleteUser(chatID)
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Done. Type any key..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

func (h *handlers) cmdDrop(ctx context.Context, tgb *bot.Bot, chatID int64) {
	msg := &bot.SendMessageParams{ChatID: chatID, Text: "Dropping session..."}
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Dropping session..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.DropSession(h.dbService)
	user.FlushMemory(h.dbService)
	h.cache.DeleteUser(chatID)
	msg = &bot.SendMessageParams{ChatID: chatID, Text: "Done. Type any key..."}
	if _, err := tgb.SendMessage(ctx, msg); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

}

// Sends a message with instructions for working with the bot
func (h *handlers) cmdHelp(ctx context.Context, tgb *bot.Bot, chatID int64) {
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: msgHelpCommand}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

// Old Func - SearchDocuments.
// WARNING: The func uses an unsafe function - GetVectorStore
func (h *handlers) cmdSearchDoc(ctx context.Context, tgb *bot.Bot, chatID int64, prompt string) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	dbLink := h.dbLink
	baseURL := user.AiSession.BaseURL

	localAIToken := user.AiSession.AIToken
	store, err := embeddings.GetVectorStore(baseURL, localAIToken, dbLink) // WARNING: This function is unsafe! May call log.Fatal (╯°□°）╯︵ ┻━┻
	if err != nil {
		if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Something happened... error occured: " + err.Error()}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}
	}

	results, err := embeddings.SemanticSearch(prompt, maxResultsForDoc, store)
	if err != nil {
		msg := &bot.SendMessageParams{ChatID: chatID, Text: "Something happened... error occured: " + err.Error()}
		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}
	}

	for i, result := range results {
		content := result.PageContent
		if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("Result number: %d\nPage content: %s", i, content)}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}

		score := result.Score
		text := fmt.Sprintf("Score: %f", score)
		if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: text}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}
	}
}

// TODO
// this is calling local-ai within base template (and without langhain injections)
func (h *handlers) cmdInstruct(ctx context.Context, tgb *bot.Bot, chatID int64, prompt string) {
	if prompt == "" {
		_, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "You didn't enter a prompt. Format: /instruct prompt"})
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}
		return
	}
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	model := user.AiSession.GptModel
	localAIToken := user.AiSession.AIToken
	text, err := langchain.GenerateContentInstruction(user.AiSession.BaseURL, prompt, model, localAIToken, user.Network)
	if err != nil || text == "" {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error generate content instruction")
		if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Error generating content instruction"}); err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
			return
		}
		return
	}

	if _, err = tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Error generating content instruction"}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

// Shows token usage statistics
func (h *handlers) cmdUsage(ctx context.Context, tgb *bot.Bot, chatID int64) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}

	promt_tokens := user.AiSession.Usage[database.Usage_PromptTokens]
	completion_tokens := user.AiSession.Usage[database.Usage_CompletionTokens]
	total_tokens := user.AiSession.Usage[database.Usage_TotalTokens]

	text := fmt.Sprintf(
		"Promt tokens: %d\nCompletion tokens: %d\nTotal tokens: %d",
		promt_tokens, completion_tokens, total_tokens,
	)

	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: text}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}

// Send helper video error
// Get a list of all files in the media directory
// files, err := os.ReadDir("../../media/")
func (h *handlers) cmdHelper(ctx context.Context, tgb *bot.Bot, chatID int64) {
	videoMsg, err := getErrorMsgWithRandomVideo(chatID)
	if err != nil {
		log.Error().Err(err).Caller().Msg("")
		return
	}
	_, err = tgb.SendVideo(ctx, videoMsg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending video message")
	}
}

func (h *handlers) cmdSetContext(ctx context.Context, tgb *bot.Bot, chatID int64, name string) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	log.Info().
		Int64("chat_id", chatID).Str("command", "set context").Str("argument", name).Interface("user", user).
		Msg("executing set context command")

	if err := user.SetContext(name); err != nil {
		log.Error().Err(err).Int64("chat_id", user.ID).Str("collection", name).
			Msg("failed to set context for user")
	}
}

func (h *handlers) cmdClearContext(ctx context.Context, tgb *bot.Bot, chatID int64) {
	user, ok := ctx.Value(database.UserCtxKey).(database.User)
	if !ok {
		log.Error().Int64("chat_id", chatID).Caller().Msg("user not found in context")
		return
	}
	user.ClearContext()
	if _, err := tgb.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: "Done. Type any key..."}); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}
}
