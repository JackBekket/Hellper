package database

import (
	"fmt"
	"os"

	//"github.com/JackBekket/hellper/lib/database"
	"github.com/rs/zerolog/log"

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
		//log.Println("error getting vectorstore")
		return err
	}
	u.VectorStore = vectorStore

	defer func() {
		var pgvStore pgvector.Store
		pgvStore, ok := vectorStore.(pgvector.Store)
		if !ok {
			//log.Fatalf("store does not implement pgvector.Store")
		}
		pgvStore.Close()
	}()

	return nil
}

func (u *User) ClearContext() {
	u.VectorStore = nil
}

func (u *User) FlushThread() error {
	return nil
}

func (u *User) FlushMemory(ds *Service) error {
	if err := ds.DropHistory(u.ID, int64(u.AiSession.ProviderID), u.ID, u.ID, u.AiSession.GptModel); err != nil {
		return fmt.Errorf("flushMemory: %w", err)
	}
	return nil
}

func (u *User) Kill(ds *Service) {
	if err := u.FlushThread(); err != nil {
		log.Error().Err(err).Caller().Msg("failed to flush thread")
	}
	if err := u.FlushMemory(ds); err != nil {
		log.Error().Err(err).Caller().Msg("failed to flush memory")
	}
	if err := ds.DeleteFromAuth(u.ID); err != nil {
		log.Error().Err(err).Caller().Msg("failed to delete from auth")
	}
	if err := ds.DeleteAISession(u.ID); err != nil {
		log.Error().Err(err).Caller().Msg("failed to delete AI session")
	}
}
func (u *User) DropSession(ds *Service) {
	if err := ds.DeleteAISession(u.ID); err != nil {
		log.Error().Err(err).Caller().Msg("failed to delete AI session")
	}
	if err := u.FlushThread(); err != nil {
		log.Error().Err(err).Caller().Msg("failed to flush thread")
	}
	if err := u.FlushMemory(ds); err != nil {
		log.Error().Err(err).Caller().Msg("failed to flush memory")
	}
	// Если нужно удалять историю, можно раскомментировать следующий код:
	// if err := ds.DropHistory(u.ID, int64(u.AiSession.AIType), u.ID, u.ID, u.AiSession.GptModel); err != nil {
	//     log.Error().Err(err).Caller().Msg("failed to drop history")
	// }
}
