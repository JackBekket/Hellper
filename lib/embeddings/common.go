package embeddings

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadEnv() {

}

// Get vector store from db. ai_url is AI url (localhost or openai or docker), api_token is AI token, db_link is database link
func GetVectorStore(ai_url string, api_token string, db_link string) (vectorstores.VectorStore, error) {

	base_url := ai_url

	//_ = godotenv.Load()


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


	config, err := pgxpool.ParseConfig(pgConnURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Create an embeddings client using the OpenAI API. Requires environment variable API_KEY to be set.
	llm, err := openai.New(
		//openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithBaseURL(base_url),
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
		pgvector.WithConn(pool),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("vector store ready")

	defer func() {
		pgvStore := store
		pgvStore.Close()
	  }()

	return store, nil

	
}

func GetVectorStoreWithOptions(ai_url string, api_token string, db_link string, name string) (vectorstores.VectorStore, error) {

	base_url := ai_url

	pgConnURL := db_link

	config, err := pgxpool.ParseConfig(pgConnURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	llm, err := openai.New(
		//openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithBaseURL(base_url),
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
		pgvector.WithCollectionName(name),
		//pgvector.WithPreDeleteCollection(true),
		pgvector.WithConn(pool),
		pgvector.WithEmbedder(e),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("vector store ready")

	defer func() {
		pgvStore := store
		pgvStore.Close()
	  }()

	return store, nil
}
