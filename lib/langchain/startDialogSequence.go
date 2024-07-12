//package main

package langchain

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	db "github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

		// Send helper video error
		// Get a list of all files in the media directory
		files, err := ioutil.ReadDir("../../media/")
		if err != nil {
		  log.Println("Could not read media directory:", err)
		  return
		}
	
		  // Select a random file
		  rand.Seed(time.Now().UnixNano())
		  randomFile := files[rand.Intn(len(files))]
	
	  // Open the video file
	  videoFile, err := os.Open(filepath.Join("../../media/", randomFile.Name()))
	  if err != nil {
		log.Println("Could not open video file:", err)
		return
	  }
	  defer videoFile.Close()
	
	  // Create a new video message
	  videoMsg := tgbotapi.NewVideo(user.ID, tgbotapi.FileReader{
		Name: randomFile.Name(),
		Reader: videoFile,
		//Size: -1, // Let the tgbotapi package determine the size
	  })
	
	  // Send the video message
	  _, err = bot.Send(videoMsg)
	  if err != nil {
		log.Println("Could not send video message:", err)
	  }










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

	thread := user.AiSession.DialogThread

	resp,post_session, err := ContinueChatWithContextNoLimit(ctx,&thread,promt)
	if err != nil {
		errorMessage(err,bot,user)
	} else {

		log.Println("AI response: ", resp)
		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 6
		usage := db.GetSessionUsage(user.ID)
		user.AiSession.Usage = usage
		//db.UsersMap[chatID] = user

		
		//log.Println("check if it's stored in messages, printing messages:")
		history, err := thread.ConversationBuffer.ChatHistory.Messages(ctx)
		if err != nil {
			log.Println(err)
		}
		
		//log.Println(history)
		total_turns := len(history)
		log.Println("total number of turns: ", total_turns)
		// Iterate over each message and print
		/*
		log.Println("Printing messages:")
		for _, msg := range history {
			log.Println(msg.GetContent())
		}
		*/

		user.AiSession.DialogThread = *post_session
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



