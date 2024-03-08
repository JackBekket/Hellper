package main

//package langchain

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/JackBekket/uncensoredgpt_tgbot/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tmc/langchaingo/llms"
)

// Notifies the user that an error occurred while creating the request.
// "An error has occured. In order to proceed we need to recreate client and initialize new session"
// Removes a user from the database (тemporary solution).
func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User) {
	log.Println("error :", err)
	msg := tgbotapi.NewMessage(user.ID, err.Error())
	bot.Send(msg)
	msg = tgbotapi.NewMessage(user.ID, "an error has occured. In order to proceed we need to recreate client and initialize new session")
	bot.Send(msg)

	userDb := db.UsersMap
	delete(userDb, user.ID)
	// updateDb := userDatabase[ID]
	// updateDb.Dialog_status = 0
	// userDatabase[ID] = updateDb
}

func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string) {
	mu.Lock()
	defer mu.Unlock()

	user := db.UsersMap[chatID]

	gptModel := user.AiSession.GptModel
	log.Printf(
		"GPT model: %s,\npromt: %s\n",
		gptModel,
		promt,
	)

	/*
	req := createComplexChatRequest(promt, gptModel)
	c := user.AiSession.GptClient
	*/

	resp, err := GenerateContentLAI(promt,gptModel,ai_endpoint)
	if err != nil {
		errorMessage(err, bot, user)
	} else {
		LogResponseContentChoice(resp)
		respText := resp.Choices[0].Content
		msg := tgbotapi.NewMessage(chatID, respText)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 4
		db.UsersMap[chatID] = user
	}

}

/*
func LogResponse(resp *llms.ContentResponse) {
	log.Println("full response obj log: ", resp)
	log.Println("created: ", resp.Choices[0].GenerationInfo)
	log.Println("resp id: ",resp.ID)
	log.Println("resp model: ", resp.Model)
	log.Println("resp object: ",resp.Object)
	log.Println("resp Choices[0]: ", resp.Choices[0])
	log.Println("resp_usage promt tokens: ", resp.Usage.PromptTokens)
	log.Println("resp_usage completion tokens: ", resp.Usage.CompletionTokens)
	log.Println("resp_usage total tokens: ",resp.Usage.TotalTokens)
}
*/


func LogResponseContentChoice(resp *llms.ContentResponse) {
	//choice *llms.ContentChoice
	choice := resp.Choices[0]
	log.Println("Content: ", choice.Content)
	log.Println("Stop Reason: ", choice.StopReason)
  
	// GenerationInfo is a map that could contain complex/nested structures,
	// so we'll marshal it into a JSON string for a cleaner log message.
	// This step is optional and depends on your preference for log clarity.
	genInfo, err := json.Marshal(choice.GenerationInfo)
	if err != nil {
	  log.Println("Error marshaling GenerationInfo: ", err)
	  return
	}
	log.Println("Generation Info: ", string(genInfo))
  
	// If you have specific fields you expect in GenerationInfo, you can log them individually:
	// Example: log.Println("Some specific gen info: ", choice.GenerationInfo["someKey"])
	
	// Note: Since FuncCall is a pointer to a schema.FunctionCall, ensure you check for nil to avoid panics.
	if choice.FuncCall != nil {
	  // Assuming FuncCall has fields you want to log, replace 'FieldName' with actual fields.
	  log.Printf("Function Call: %+v\n", choice.FuncCall)
	  // For specific field: log.Println("FuncCall field: ", choice.FuncCall.FieldName)
	} else {
	  log.Println("No Function Call requested.")
	}
  }
