package command

var msgTemplates = map[string]string{
	"hello":      "Hey, this bot is working with local ai node.",
	"case0":      "If you choosed OpenAI then input your openai-key, if you choosed LocalAI then input password. DM me in telegram if you want password or buy key from openai",
	"await":      "Awaiting",
	"case1":      "Choose model to use. ",
	"ch_network" : "choose network to work with. openai make calls to openai, localhost tries to access localai node installed alongside with bot",
	"help_command" : " input text to continue dialog, ask a question or use commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation",
}
