package database

import (
	"fmt"
	"log"
	"os"

	//"github.com/JackBekket/hellper/lib/database"

	e "github.com/JackBekket/hellper/lib/embeddings"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func (u *User) SetContext(collectionName string) error {

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

func (u *User) ClearContext() {
	u.VectorStore = nil
}

func (u *User) FlushThread() {
}

func (u *User) FlushMemory(ds *Service) {
	err := ds.DropHistory(u.ID, int64(u.AiSession.AI_Type), u.ID, u.ID, u.AiSession.GptModel)

	fmt.Println(err)
}

func (u *User) Kill(ds *Service) {
	u.FlushThread()
	u.FlushMemory(ds)
	ds.DeleteToken(u.ID, 1)
	ds.DeleteLSession(u.ID)
	delete(UsersMap, u.ID)
}

func (u *User) DropSession(ds *Service) {
	ds.DeleteLSession(u.ID)
	//ds.DropHistory(u.ID,int64(u.AiSession.AI_Type),u.ID,u.ID,u.AiSession.GptModel)
}
