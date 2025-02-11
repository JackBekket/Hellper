//package main

package langchain

import (
	"context"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"io/fs"

	"github.com/JackBekket/hellper/lib/agent"
	"github.com/JackBekket/hellper/lib/database"
	db "github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	//TODO: investigate why meme videos with helper are not sent by this func!
	// Notifies the user that an error occurred while creating the request.
	// "An error has occured. In order to proceed we need to recreate client and initialize new session"
	// Removes a user from the database (тemporary solution).
	"sort"
)

func errorMessage(err error, bot *tgbotapi.BotAPI, user db.User) {
	log.Println("error :", err)
	msg := tgbotapi.NewMessage(user.ID, err.Error())
	bot.Send(msg)
	msg = tgbotapi.NewMessage(user.ID, "an error has occured. In order to proceed we need to recreate client and initialize new session")
	bot.Send(msg)

	// Send helper video error
	// Get a list of all files in the media directory
	files, err := func() ([]fs.FileInfo, error) {
		f, err := os.Open("../../media/")
		if err != nil {
			return nil, err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, err
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name() < list[j].Name()
		})
		return list, nil
	}()
	if err != nil {
		log.Println("Could not read media directory:", err)
		return
	}

	// Select a random file
	//rand.Seed(time.Now().UnixNano())
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
		Name:   randomFile.Name(),
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

}


func StartDialogSequence(bot *tgbotapi.BotAPI, chatID int64, promt string, ctx context.Context, ai_endpoint string, ds *database.Service) {
	mu.Lock()
	defer mu.Unlock()

	user := db.UsersMap[chatID]

	gptModel := user.AiSession.GptModel
	log.Printf(
		"GPT model: %s,\npromt: %s\n",
		gptModel,
		promt,
	)
	api_key := user.AiSession.GptKey
	base_url := ai_endpoint

	thread := user.AiSession.DialogThread

	post_session, resp, err := ContinueAgent(api_key, gptModel, base_url, promt, &thread)
	if err != nil {
		errorMessage(err, bot, user)
	} else {

		log.Println("AI response: ", resp)
		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ParseMode = "MARKDOWN"
		bot.Send(msg)

		user.DialogStatus = 6
		usage := db.GetSessionUsage(user.ID)
		user.AiSession.Usage = usage
		//db.UsersMap[chatID] = user

		//log.Println(history)
		total_turns := len(thread.ConversationBuffer)
		log.Println("total number of turns: ", total_turns)
		// Iterate over each message and print
		/*
			log.Println("Printing messages:")
			for _, msg := range history {
				log.Println(msg.GetContent())
			}
		*/

		user.AiSession.DialogThread = *post_session
		buffer := post_session.ConversationBuffer
		last_msg := buffer[len(buffer)-1]
		db.UsersMap[chatID] = user	//save in cash
		//TODO: here we should save user conversation to the db?
		h := agent.CreateMessageContentHuman(promt)
		ds.UpdateHistory(chatID,1,chatID,chatID,gptModel,h[0])
		ds.UpdateHistory(chatID,1,chatID,chatID,gptModel,last_msg)
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
