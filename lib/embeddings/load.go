package embeddings

import (
	"context"
	"fmt"
	"net/http"

	"log"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)




func LoadDocsToStore(docs []schema.Document, store vectorstores.VectorStore)  {
	fmt.Println("loading data from")

	fmt.Println("no. of documents to be loaded", len(docs))

	_, err := store.AddDocuments(context.Background(), docs)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("data successfully loaded into vector store")
	log.Println(err)

	defer func() {
		var pgvStore pgvector.Store
		pgvStore, ok := store.(pgvector.Store)
		if !ok {
		  log.Fatalf("store does not implement pgvector.Store")
		}
		pgvStore.Close()
	  }()
}

/*
func GetTextDocs() {
	var docs []schema.Document

    for _, data := range fileData {
        doc := schema.Document{
            PageContent: data.Content,
            Metadata: map[string]interface{}{
                "date": data.Date,
            },
        }
        docs = append(docs, doc)
    }

    return docs
}
*/

func getDocs(source string) ([]schema.Document, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	docs, err := documentloaders.NewHTML(resp.Body).LoadAndSplit(context.Background(), textsplitter.NewRecursiveCharacter())

	if err != nil {
		return nil, err
	}

	return docs, nil
}