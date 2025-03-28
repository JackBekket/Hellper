package embeddings

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func Rag(endpoint string, localAIToken string, question string, numOfResults int, store vectorstores.VectorStore, option ...vectorstores.Option) (result string, err error) {

	// Create an embeddings client using the.
	llm, err := openai.New(
		//openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithBaseURL(endpoint),
		openai.WithAPIVersion("v1"),
		openai.WithToken(localAIToken),
		openai.WithModel("tiger-gemma-9b-v1-i1"),
		openai.WithEmbeddingModel("text-embedding-ada-002"),
	)
	if err != nil {
		return "", err
	}

	result, err = chains.Run(
		context.Background(),
		chains.NewRetrievalQAFromLLM(
			llm,
			vectorstores.ToRetriever(store, numOfResults, option...),
		),
		question,
		chains.WithMaxTokens(8192),
	)
	if err != nil {
		return "", err
	}

	fmt.Println("====final answer====\n", result)

	defer func() {
		var pgvStore pgvector.Store
		pgvStore, ok := store.(pgvector.Store)
		if !ok {
			log.Fatalf("store does not implement pgvector.Store")
		}
		pgvStore.Close()
	}()

	return result, nil
}

func SemanticSearch(searchQuery string, maxResults int, store vectorstores.VectorStore, options ...vectorstores.Option) (searchResults []schema.Document, err error) {

	searchResults, err = store.SimilaritySearch(context.Background(), searchQuery, maxResults, options...)
	if err != nil {
		return nil, err
	}
	fmt.Println("============== similarity search results ==============")

	for _, doc := range searchResults {
		fmt.Println("similarity search info -", doc.PageContent)
		fmt.Println("similarity search score -", doc.Score)
		fmt.Println("============================")

	}

	defer func() {
		var pgvStore pgvector.Store
		pgvStore, ok := store.(pgvector.Store)
		if !ok {
			log.Fatalf("store does not implement pgvector.Store")
		}
		pgvStore.Close()
	}()
	return searchResults, nil
}
