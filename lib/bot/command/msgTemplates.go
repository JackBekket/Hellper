package command

var msgTemplates = map[string]string{
	"hello":      "Hey, this bot is working with local ai node.",
	"case0":      "input your password",
	"await":      "Awaiting",
	"case1":      "Choose model to use. ",
	"codex_help": "``` # describe your task in comments like this or put your lines of code you need to autocomplete ```",
	"ch_network" : "choose network to work with. openai make calls to openai, localhost tries to access localai node installed alongside with bot",
}
