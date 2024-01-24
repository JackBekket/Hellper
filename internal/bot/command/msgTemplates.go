package command

var msgTemplates = map[string]string{
	"hello":      "Hey, this bot is OpenAI chatGPT.",
	"case0":      "Input your openAI API key. It can be created at https://platform.openai.com/account/api-keys  ..in case you are try to access localai node you should have local pwd for it",
	"await":      "Awaiting",
	"case1":      "Choose model to use. ",
	"codex_help": "``` # describe your task in comments like this or put your lines of code you need to autocomplete ```",
	"ch_network" : "choose network to work with. openai make calls to openai, localhost tries to access localai node installed alongside with bot",
}
