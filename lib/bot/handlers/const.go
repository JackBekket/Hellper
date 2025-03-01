package handlers

import "fmt"

// Messages for the user
const (
	msg_Hello              = "Hey, this bot is working with LocalAI node! Please input your local-ai api_key 🐱"
	msg_Await              = "Awaiting..."
	msg_Choose_model       = "Choose model to use "
	msg_Session_model      = "Your session model: "
	msg_Choose_lang        = "Choose a language or send 'Hello' in your desired language"
	msg_Connecting_AI_node = "Connecting to AI node..."
	msg_Help_command       = "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire"
	msg_AI_client_failure  = "An error has occured. In order to proceed we need to recreate client and initialize new session"
)

// Base prompts for the AI
const (
	basePrompt_Lang_RU        = "Привет, ты говоришь по-русски?"
	basePrompt_Lang_EN        = "Hi, Do you speak english?"
	basePrompt_RecognizeImage = "What's in the image?"
	basePromt_GenerateImage   = "evangelion, neon, anime"
)

// Context values for bot commands and arguments
const (
	context_BotCommand = "command"
	context_CommandArg = "arg"
)

// Dialog status. The user's current position in the conversation with the bot
const (
	status_AIModelSelectionKeyboard           = 3
	status_AIModelSelectionChoice             = 4
	status_ConnectingToAiWithLang             = 5
	status_MainHandlerAfterUserIdentification = 6
)

// Error messages for the user
const (
	errorMsg_FailedToGenerateImage = "Failed to generate image"
	errMsg_FailedTrascribeVoice    = "Failed to transcribe the voice message"
	errMsg_FailedRecognizeImage    = "Failed to recognize the image"
)

// Default values for working with AI. Model names, URL suffixes
const (
	ai_StableDiffusionModel  = "stablediffusion"
	ai_ImageGenerationSuffix = "/v1/images/generations"

	ai_BunnyLLAMAModel      = "bunny-llama-3-8b-v"
	ai_ImageRecognizeSuffix = "/v1/chat/completions"
)

// Constructs a Telegram file URL.
func urlTelegramServeFilesConstructor(token string, filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
}
func urlConnectionToAIConstructor(endpoint string, suffix string) string {
	if suffix == "" {
		suffix = ai_ImageRecognizeSuffix
	}
	return endpoint + suffix
}

// Maximum number of results when searching for documents
const maxResultsForDoc = 3
