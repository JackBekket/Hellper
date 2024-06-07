package embeddings

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
)



func Rag(ai_url string,api_token string,question string, numOfResults int,store vectorstores.VectorStore) (result string,err error) {

	//base_url := os.Getenv("AI_BASEURL")
	base_url := ai_url

	// Create an embeddings client using the. 
	llm, err := openai.New(
		//openai.WithBaseURL("http://localhost:8080/v1/"),
		openai.WithBaseURL(base_url),
		openai.WithAPIVersion("v1"),
		openai.WithToken(api_token),
    	openai.WithModel("wizard-uncensored-13b"),
    	openai.WithEmbeddingModel("text-embedding-ada-002"),
	)
	if err != nil {
		return "",err
	}

	result, err = chains.Run(
		context.Background(),
		chains.NewRetrievalQAFromLLM(
			llm,
			vectorstores.ToRetriever(store, numOfResults),
		),
		question,
		chains.WithMaxTokens(2048),
	)
	if err != nil {
		return "",err
	}

	fmt.Println("====final answer====\n", result)

	return result,nil

}

func SemanticSearch(searchQuery string, maxResults int, store vectorstores.VectorStore, options ...vectorstores.Option) (searchResults []schema.Document, err error) {
	//var store vectorstores.VectorStore


	/*
	store, err := GetVectorStore()
	if err != nil {
		return nil,err
	}
	*/

	searchResults, err = store.SimilaritySearch(context.Background(), searchQuery, maxResults, options...)

	if err != nil {
		return nil,err
	}

	fmt.Println("============== similarity search results ==============")

	for _, doc := range searchResults {
		fmt.Println("similarity search info -", doc.PageContent)
		fmt.Println("similarity search score -", doc.Score)
		fmt.Println("============================")

	}

	return searchResults,nil

}

