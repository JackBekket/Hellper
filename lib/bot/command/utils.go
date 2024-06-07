package command

import (
	"fmt"
	"log"
	"os"

	db "github.com/JackBekket/hellper/lib/database"
	"github.com/JackBekket/hellper/lib/embeddings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)


func (c *Commander) HelpCommandMessage(updateMessage *tgbotapi.Message)  {
	chatID := updateMessage.From.ID
	user := db.UsersMap[chatID]
	msg := tgbotapi.NewMessage(user.ID, msgTemplates["help_command"])
	c.bot.Send(msg)
}

func (c *Commander) SearchDocuments(chatID int64, promt string, maxResults int) {
	//chatID := updateMessage.From.ID
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

	conn_pg_link := os.Getenv("PG_LINK")
	base_url := os.Getenv("AI_BASEURL")
	db_conn := conn_pg_link
	api_token := user.AiSession.GptKey
	store,err := embeddings.GetVectorStore(base_url,api_token,db_conn)
	if err != nil {
		//return nil, err
		msg := tgbotapi.NewMessage(user.ID, "error occured when getting store: " + err.Error())
		c.bot.Send(msg)
	}

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



