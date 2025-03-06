package handlers

// Messages for the user
const (
	msgHello            = "Hey, this bot is working with LocalAI node! Please input your local-ai API token 🐱"
	msgAwait            = "Awaiting..."
	msgChooseModel      = "Choose model to use "
	msgSessionModel     = "Your session model: "
	msgChooseLang       = "Choose a language or send 'Hello' in your desired language"
	msgConnectingAINode = "Connecting to AI node..."
	msgHelpCommand      = "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /searchdoc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire"
	msgAIclientFailure  = "An error has occured. In order to proceed we need to recreate client and initialize new session"
)

// Base prompts for the AI
const (
	basePromptLangRU         = "Привет, ты говоришь по-русски?"
	basePromptLangEN         = "Hi, Do you speak english?"
	basePromptRecognizeImage = "What's in the image?"
	basePromptGenerateImage  = "evangelion, neon, anime"
)

// Dialog status. The user's current position in the conversation with the bot
const (
	statusAIModelSelectionKeyboardForNewUser = iota
	statusAIModelSelectionKeyboardForExistUser
	statusAIModelSelectionChoice
	statusConnectingToAiWithLang
	statusStartDialogSequence
)

// Error messages for the user
const (
	errMsgFailedToGenerateImage = "Failed to generate image"
	errMsgFailedTrascribeVoice  = "Failed to transcribe the voice message"
	errMsgFailedRecognizeImage  = "Failed to recognize the image"
)

// Default values for working with AI. Model names, URL endpointes
const (
	aiStableDiffusionModel    = "stablediffusion"
	aiImageGenerationEndpoint = "/v1/images/generations"

	aiBunnyLLAMAModel        = "bunny-llama-3-8b-v"
	aiImageRecognizeEndpoint = "/v1/chat/completions"
)

// Maximum number of results when searching for documents
const maxResultsForDoc = 3
