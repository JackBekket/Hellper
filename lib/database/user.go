package database

import (
	"log"
	"os"

	e "github.com/JackBekket/hellper/lib/embeddings"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)


func (u *User) SetContext (collectionName string) error{
	
		_ = godotenv.Load()
		api_token := u.AiSession.GptKey
		ai_endpoint := os.Getenv("AI_ENDPOINT")
		//log.Println("ai endpoint is: ", ai_endpoint)
		db_link := os.Getenv("EMBEDDINGS_DB_URL")
	
		vectorStore, err := e.GetVectorStoreWithOptions(ai_endpoint, api_token, db_link, collectionName)
		if err != nil {
			log.Println("error getting vectorstore")
			return err
		}
		u.VectorStore = vectorStore
	
		defer func() {
			var pgvStore pgvector.Store
			pgvStore, ok := vectorStore.(pgvector.Store)
			if !ok {
			  log.Fatalf("store does not implement pgvector.Store")
			}
			pgvStore.Close()
		  }()
	
	
		return nil
}



func (u *User)ClearContext() {
	u.VectorStore = nil
}
