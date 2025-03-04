package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/JackBekket/hellper/lib/localai"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
)

// Makes a request for image generation and sends the user the path where the image is located.
// Currently, the default model used is Stable Diffusion
func (h *handlers) cmdGenerateImage(ctx context.Context, tgb *bot.Bot, chatID int64, prompt string) {
	msg := &bot.SendMessageParams{ChatID: chatID, Text: "Image link generation..."}
	_, err := tgb.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	msgFailedGenerateImageFunc := func() {
		msg := &bot.SendMessageParams{ChatID: chatID, Text: errorMsg_FailedToGenerateImage}
		_, err := tgb.SendMessage(ctx, msg)
		if err != nil {
			log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		}
	}

	if prompt == "" {
		prompt = basePromt_GenerateImage
	}

	url := getURL(h.config.AIEndpoint, h.config.ImageGenerationSuffix)
	size := "256x256"
	model := h.config.ImageGenerationModel
	if model == "" {
		model = ai_StableDiffusionModel
	}

	user, ok := h.cache.GetUser(chatID)
	if !ok {
		log.Error().Int64("chat_id", chatID).Msg("user not found in cache")
		return
		// todo: Add actions in case the user is not found in the cache
	}

	localAIToken := user.AiSession.LocalAIToken

	pathToImage, err := localai.GenerateImageStableDiffusion(prompt, size, url, model, localAIToken)
	if err != nil {
		msgFailedGenerateImageFunc()
		log.Error().Err(err).Str("prompt", prompt).Str("size", size).Str("url", url).Str("model", model).Caller().
			Msg("failed to generate image with Stable Diffusion")
		return
	}

	imageMsg, err := getMsgWithImage(chatID, pathToImage, localAIToken)
	if err != nil {
		msgFailedGenerateImageFunc()
		log.Error().Err(err).Str("prompt", prompt).Str("size", size).Str("url", url).Str("model", model).Caller().
			Msg("failed to generate image with Stable Diffusion")
		return
	}

	_, err = tgb.SendPhoto(ctx, imageMsg)
	if err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error sending message")
		return
	}

	if err := os.Remove(pathToImage); err != nil {
		log.Error().Err(err).Int64("chat_id", chatID).Caller().Msg("error deleting image")
		return
	}
}

func getMsgWithImage(chatID int64, pathToImage string, localAIToken string) (*bot.SendPhotoParams, error) {
	fileName, err := getImage(pathToImage, localAIToken)
	if err != nil {
		return &bot.SendPhotoParams{}, err
	}

	path := filepath.Join("tmp", "generated", "images", fileName)
	imageFile, err := os.Open(path)
	if err != nil {
		return &bot.SendPhotoParams{}, err
	}
	defer imageFile.Close()

	if err != nil {
		return &bot.SendPhotoParams{}, err
	}

	return &bot.SendPhotoParams{
		ChatID: chatID,
		Photo: &models.InputFileUpload{
			Filename: "picture",
			Data:     imageFile,
		},
	}, nil

}

func getImage(imageURL, localAITokenHeader string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}
	req.Header.Add("localAITokenorization", "Bearer "+localAITokenHeader)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch the image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch image, status code: %d", resp.StatusCode)
	}

	fileName := transformURL(imageURL)

	dir := filepath.Join("tmp", "generated", "images")
	filePath := filepath.Join(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	return fileName, nil

}
