package handlers

import (
	"fmt"
	"net/url"
	"path"
)

func transformURL(inputURL string) string {
	parsedURL, _ := url.Parse(inputURL)
	fileName := path.Base(parsedURL.Path)
	return fileName
}

func getURL(baseURL string, enpoint string) string {
	if enpoint == "" {
		enpoint = aiImageGenerationEndpoint
	}
	joined, err := url.JoinPath(baseURL, enpoint)
	if err != nil {
		return baseURL + enpoint
	}
	return joined
}

// Constructs a Telegram file URL.
func urlTelegramServeFilesConstructor(token string, filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
}
