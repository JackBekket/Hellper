package command

var msgTemplates = map[string]string{
	"hello":      "Hey, this bot is working with local ai node.",
	"case0":      "If you choosed OpenAI then input your openai-key, if you choosed LocalAI then input password. DM me in telegram if you want password or buy key from openai",
	"await":      "Awaiting",
	"case1":      "Choose model to use. ",
	"ch_network" : "choose network to work with. openai make calls to openai, localhost tries to access localai node installed alongside with bot",
	"help_command" : "Authorize for additional commands: /help -- print this message, /restart -- restart session (if you want to switch between local-ai and openai chatGPT), /search_doc -- searching documents, /rag -- process Retrival-Augmented Generation, /instruct -- use system promt template instead of langchain (higher priority, see examples), /image -- generate image ....all funcs are experimental so bot can halt and catch fire",
}
