package handlers

import (
	"net/url"
	"path"
)

func transformURL(inputURL string) string {
	parsedURL, _ := url.Parse(inputURL)
	fileName := path.Base(parsedURL.Path)
	return fileName
}

func getURL(endpoint string, suffix string) string {
	if suffix == "" {
		suffix = ai_ImageGenerationSuffix
	}
	joined, err := url.JoinPath(endpoint, suffix)
	if err != nil {
		return endpoint + suffix
	}
	return joined
}
