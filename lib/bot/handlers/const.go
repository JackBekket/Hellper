package handlers

var msgTemplates = map[string]string{
	"hello":        "Hey, this bot is working with LocalAI node! Please input your local-ai api_key 🐱",
	"await":        "Awaiting",
	"case1":        "Choose model to use. ",
	"help_command": "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire",
}
