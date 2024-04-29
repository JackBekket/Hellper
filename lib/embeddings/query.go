package embeddings

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
)

func RagSearch(question string, numOfResults int) error {

	store, err := GetVectorStore()

	if err != nil {
		return err
	}


	// Create an embeddings client using the OpenAI API. Requires environment variable OPENAI_API_KEY to be set.
	llm, err := openai.New(
		openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithAPIVersion("v1"),
    	openai.WithModel("wizard-uncensored-13b"),
    //openai.WithEmbeddingModel("text-embedding-ada-002"),
	)
	if err != nil {
		return err
	}

	result, err := chains.Run(
		context.Background(),
		chains.NewRetrievalQAFromLLM(
			llm,
			vectorstores.ToRetriever(store, numOfResults),
		),
		question,
		chains.WithMaxTokens(2048),
	)
	if err != nil {
		return err
	}

	fmt.Println("====final answer====\n", result)

	return nil

}

func SemanticSearch(searchQuery string, maxResults int) error {

	store, err := GetVectorStore()
	if err != nil {
		return err
	}

	searchResults, err := store.SimilaritySearch(context.Background(), searchQuery, maxResults)

	if err != nil {
		return err
	}

	fmt.Println("============== similarity search results ==============")

	for _, doc := range searchResults {
		fmt.Println("similarity search info -", doc.PageContent)
		fmt.Println("similarity search score -", doc.Score)
		fmt.Println("============================")

	}

	return nil

}