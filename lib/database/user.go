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
	localAIToken := u.AiSession.AIToken
	ai_endpoint := os.Getenv("AI_ENDPOINT")
	//log.Println("ai endpoint is: ", ai_endpoint)
	dbLink := os.Getenv("EMBEDDINGS_DB_URL")

	vectorStore, err := e.GetVectorStoreWithOptions(ai_endpoint, localAIToken, dbLink, collectionName)
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

func (u *User) FlushMemory(ds *Service) error {
	if err := ds.DropHistory(u.ID, int64(u.AiSession.ProviderID), u.ID, u.ID, u.AiSession.GptModel); err != nil {
		return fmt.Errorf("flushMemory: %w", err)
	}
	return nil
}

func (u *User) Kill(ds *Service) {
	u.FlushThread()
	u.FlushMemory(ds) // dublicate
	ds.DeleteToken(u.ID, 1)
	ds.DeleteAISession(u.ID)
	delete(UsersMap, u.ID)
}

func (u *User) DropSession(ds *Service) {
	ds.DeleteAISession(u.ID)
	u.FlushThread()
	u.FlushMemory(ds) // dublicate

	//ds.DropHistory(u.ID,int64(u.AiSession.AIType),u.ID,u.ID,u.AiSession.GptModel)
}
