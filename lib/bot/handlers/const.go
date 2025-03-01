package handlers

const (
	msg_Hello              = "Hey, this bot is working with LocalAI node! Please input your local-ai api_key 🐱"
	msg_Await              = "Awaiting..."
	msg_Choose_model       = "Choose model to use "
	msg_Session_model      = "Your session model: "
	msg_Choose_lang        = "Choose a language or send 'Hello' in your desired language"
	msg_Connecting_AI_node = "Connecting to AI node..."
	msg_Help_command       = "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire"
	msg_AI_client_failure  = "An error has occured. In order to proceed we need to recreate client and initialize new session"

	initialPrompt_Lang_RU = "Привет, ты говоришь по-русски?"
	initialPrompt_Lang_EN = "Hi, Do you speak english?"
)
