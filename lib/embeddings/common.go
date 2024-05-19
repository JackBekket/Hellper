package embeddings

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)


func LoadEnv() {

}

func GetVectorStore(api_token string, db_link string) (vectorstores.VectorStore, error) {

	/*
	_ = godotenv.Load()
	api_token = os.Getenv("OPENAI_API_KEY")	// this is not openai key actually, it's local key for localai
	conn_pg_link := os.Getenv("PG_LINK")
  	*/

	


	/*
	host := os.Getenv("PG_HOST")
	if host == "" {
		log.Fatal("missing PG_HOST")
	}

	user := os.Getenv("PG_USER")
	if user == "" {
		log.Fatal("missing PG_USER")
	}

	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		log.Fatal("missing PG_PASSWORD")
	}

	dbName := os.Getenv("PG_DB")
	if dbName == "" {
		log.Fatal("missing PG_DB")
	}

	connURLFormat := "postgres://%s:%s@%s:5432/%s?sslmode=disable"

	pgConnURL := fmt.Sprintf(connURLFormat, user, url.QueryEscape(password), host, dbName)
	*/
	pgConnURL := db_link


	/*
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	*/

	// Create an embeddings client using the OpenAI API. Requires environment variable API_KEY to be set.
	llm, err := openai.New(
		openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithAPIVersion("v1"),
    	//openai.WithModel("wizard-uncensored-13b"),
    	openai.WithEmbeddingModel("text-embedding-ada-002"),
		openai.WithToken(api_token),
	)
	if err != nil {
		log.Fatal(err)
	}

	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return nil, err
	}

	store, err := pgvector.New(
		context.Background(),
		//pgvector.WithPreDeleteCollection(true),
		pgvector.WithConnectionURL(pgConnURL),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("vector store ready")

	return store, nil
}