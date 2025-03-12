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

func getURL(baseURL string, endpoint string) string {
	joined, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		return baseURL + endpoint
	}
	return joined
}

// Constructs a Telegram file URL.
func urlTelegramServeFilesConstructor(token string, filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
}
