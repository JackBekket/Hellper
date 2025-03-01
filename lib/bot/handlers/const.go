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

// Initial prompts for starting a conversation with the AI
const (
	initialPrompt_Lang_RU = "Привет, ты говоришь по-русски?"
	initialPrompt_Lang_EN = "Hi, Do you speak english?"
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
)

// Default values for working with AI. Model names, URL suffixes, prompts
const (
	ai_StableDiffusionModel  = "stablediffusion"
	ai_ImageGenerationSuffix = "/v1/images/generations"
	ai_DefaultPromptForImage = "evangelion, neon, anime"
)

// Constructs a Telegram file URL.
func urlTelegramServeFilesConstructor(token string, filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, filePath)
}

// Maximum number of results when searching for documents
const maxResultsForDoc = 3
