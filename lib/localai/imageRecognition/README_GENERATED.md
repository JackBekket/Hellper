# imageRecognition

This package provides a Telegram bot that recognizes images using an AI model. It uses the OpenAI API to perform image recognition.

## Overview

The package is designed to interact with the Telegram bot API (via `github.com/go-telegram-bot-api/telegram-bot-api/v5`) and perform image recognition tasks. It retrieves environment variables for the AI endpoint, model, and API token, and provides default values if any of these variables are not set.

## Configuration

* Environment variables:
	+ `AI_ENDPOINT`
	+ `IMAGE_RECOGNITION_SUFFIX`
	+ `IMAGE_RECOGNITION_MODEL`
	+ `OPENAI_API_KEY`
* No other configuration options are available.

## Usage

### Launching the application

The package can be launched as a Telegram bot by running the `imageRecognition.go` file. The bot can be interacted with by sending messages to the bot, which will then respond with the recognized text.

### Edge cases

* The package does not support any edge cases, as it is designed to work with a specific use case (image recognition via Telegram bot).

### Files and paths

* `imageRecognition.go`: main entry point for the package
* `lib/localai/imageRecognition/imageRecognition.go`: not used

## Relations between code entities

The package is structured in a way that `getEnvsForImgRec` function retrieves environment variables, which are then used by `RecognizeImage` function to perform image recognition. `handleImageMessage` function is used to retrieve the image file from the Telegram bot, and `imageRecognitionLAI` function sends a POST request to the OpenAI API with the image link and prompt.

## Unclear places or dead code

None found.

