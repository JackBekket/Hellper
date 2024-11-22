package command

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	db "github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/embeddings"
	"github.com/JackBekket/hellper/lib/localai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func (c *Commander) HelpCommandMessage(updateMessage *tgbotapi.Message)  {
	chatID := updateMessage.Chat.ID
	user := db.UsersMap[chatID]
	msg := tgbotapi.NewMessage(user.ID, msgTemplates["help_command"])
	c.bot.Send(msg)
}

func (c *Commander) SearchDocuments(chatID int64, promt string, maxResults int) {
	
	_ = godotenv.Load()

	conn_pg_link := os.Getenv("PG_LINK")
	base_url := os.Getenv("AI_BASEURL")
	db_conn := conn_pg_link
	user := db.UsersMap[chatID]
	api_token := user.AiSession.GptKey
	store,err := embeddings.GetVectorStore(base_url,api_token,db_conn)
	if err != nil {
		//return nil, err
		msg := tgbotapi.NewMessage(user.ID, "error occured: " + err.Error())
		c.bot.Send(msg)
	}


	results, err := embeddings.SemanticSearch(promt,maxResults,store)
	if err != nil {
		//return nil, err
		msg := tgbotapi.NewMessage(user.ID, "error occured: " + err.Error())
		c.bot.Send(msg)
	}

	for i, result := range results {
		content := result.PageContent
		msg := tgbotapi.NewMessage(user.ID, "result number: " + fmt.Sprint(i))
		c.bot.Send(msg)
		msg = tgbotapi.NewMessage(user.ID, "page content: " + content)
		c.bot.Send(msg)

		score := result.Score
		score_string := fmt.Sprintf("%f", score)

		msg = tgbotapi.NewMessage(user.ID, "score: " + score_string)
		c.bot.Send(msg)
	}

}

// Retrival-Augmented Generation
func (c *Commander) RAG(chatID int64, promt string, maxResults int) {
	user := db.UsersMap[chatID]
	_ = godotenv.Load()

	//conn_pg_link := os.Getenv("PG_LINK")
	base_url := os.Getenv("AI_BASEURL")
	//db_conn := conn_pg_link
	api_token := user.AiSession.GptKey
	store := user.VectorStore
	//store,err := embeddings.GetVectorStore(base_url,api_token,db_conn)
	/*
	if err != nil {
		//return nil, err
		msg := tgbotapi.NewMessage(user.ID, "error occured when getting store: " + err.Error())
		c.bot.Send(msg)
	}
		*/

	result, err := embeddings.Rag(base_url,api_token,promt,maxResults,store)
	if err != nil {
		msg := tgbotapi.NewMessage(user.ID, "error occured when calling RAG: " + err.Error())
		c.bot.Send(msg)
	}
	msg := tgbotapi.NewMessage(user.ID, result)
	c.bot.Send(msg)
}


// Get usage for user 
func (c *Commander) GetUsage(chatID int64)  {
	user := db.UsersMap[chatID]
	log.Println("user", user)
	promt_tokens := user.AiSession.Usage["Promt"]
	completion_tokens := user.AiSession.Usage["Completion"]
	total_tokens := user.AiSession.Usage["Total"]

	pt_str := fmt.Sprint(promt_tokens)
	ct_str := fmt.Sprint(completion_tokens)
	tt_str := fmt.Sprint(total_tokens)

	msg := tgbotapi.NewMessage(user.ID, "Promt tokens: " + pt_str)
	c.bot.Send(msg)
	msg = tgbotapi.NewMessage(user.ID, "Completion tokens: " + ct_str)
	c.bot.Send(msg)
	msg = tgbotapi.NewMessage(user.ID, "Total tokens: " + tt_str)
	c.bot.Send(msg)
}



func (c *Commander) SendMediaHelper(chatID int64) {

		// Send helper video error
		// Get a list of all files in the media directory
		//files, err := os.ReadDir("../../media/")
		files, err := os.ReadDir("./media/")
		if err != nil {
		  log.Println("Could not read media directory:", err)
		  return
		}
	
		if len(files) == 0 {
			log.Println("No files in media directory")
			return  
		}


		// Select a random file
		//rand.Seed(time.Now().UnixNano())
		randomFile := files[rand.Intn(len(files))]
	
	  // Open the video file
	  videoFile, err := os.Open(filepath.Join("./media/", randomFile.Name()))
	  if err != nil {
		log.Println("Could not open video file:", err)
		return
	  }
	  defer videoFile.Close()
	
	  // Create a new video message
	  videoMsg := tgbotapi.NewVideo(chatID, tgbotapi.FileReader{
		Name: randomFile.Name(),
		Reader: videoFile,
		//Size: -1, // Let the tgbotapi package determine the size
	  })
	
	  // Send the video message
	  _, err = c.bot.Send(videoMsg)
	  if err != nil {
		log.Println("Could not send video message:", err)
	  }

}

func sendImage(bot *tgbotapi.BotAPI, chatID int64, path string) {

	auth := os.Getenv("OPENAI_API_KEY")

	fileName, err := getImage(path, auth)
	if err != nil {
		fmt.Errorf("getImageFail: %w", err)
	}
	filePath := filepath.Join("tmp", "generated", "images", fileName)
	photoBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	bot.Send(tgbotapi.NewPhoto(int64(chatID), photoFileBytes))
	DeleteFile(filePath)
}

func getImage(imageURL, authHeader string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+authHeader)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch the image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch image, status code: %d", resp.StatusCode)
	}

	fileName := transformURL(imageURL)

	dir := filepath.Join("tmp", "generated", "images")
	filePath := filepath.Join(dir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	return fileName, nil

}

func DeleteFile(fileName string) {
	os.Remove(fileName)
}

func transformURL(inputURL string) string {
	// Replace "http://localhost:8080" with "/tmp" using strings.Replace
	parsedURL, _ := url.Parse(inputURL)

	// Use path.Base to get the filename from the URL path
	fileName := path.Base(parsedURL.Path)
	return fileName
}


// stable diffusion
func (c *Commander) GenerateNewImageLAI_SD(promt, url string, chatID int64, bot *tgbotapi.BotAPI) {
	size := "256x256"
	model := os.Getenv("IMAGE_GENERATION_MODEL")
	if model == "" {
		model = "stablediffusion"
	}
	urlSuffix := os.Getenv("IMAGE_GENERATION_SUFFIX")
	if urlSuffix == "" {
		urlSuffix = "/v1/images/generations"
	}
	url += urlSuffix

	filepath, err := localai.GenerateImageStableDiffusion(promt, size, url, model)
	if err != nil {
		//return nil, err
		log.Println(err)
	}
	log.Println("url_path: ", filepath)

	sendImage(bot, chatID, filepath)
}


