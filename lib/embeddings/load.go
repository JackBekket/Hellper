package embeddings

import (
	"context"
	"fmt"
	"net/http"

	"log"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

/*
func LoadDocs() error {

	fmt.Println("loading data from")

	store, err := GetVectorStore()

	if err != nil {
		return err
	}
	docs := getSampleDocs()

	fmt.Println("no. of documents to be loaded", len(docs))

	_, err = store.AddDocuments(context.Background(), docs)

	if err != nil {
		return err
	}

	fmt.Println("data successfully loaded into vector store")

	return nil
}
*/

/*
func getSampleDocs() ([]schema.Document){

		docs:= []schema.Document{
			{
				PageContent: "Tokyo",
				Metadata: map[string]any{
					"population": 38,
					"area":       2190,
				},
			},
			{
				PageContent: "Paris",
				Metadata: map[string]any{
					"population": 11,
					"area":       105,
				},
			},
			{
				PageContent: "London",
				Metadata: map[string]any{
					"population": 9.5,
					"area":       1572,
				},
			},
			{
				PageContent: "Santiago",
				Metadata: map[string]any{
					"population": 6.9,
					"area":       641,
				},
			},
			{
				PageContent: "Buenos Aires",
				Metadata: map[string]any{
					"population": 15.5,
					"area":       203,
				},
			},
			{
				PageContent: "Rio de Janeiro",
				Metadata: map[string]any{
					"population": 13.7,
					"area":       1200,
				},
			},
			{
				PageContent: "Sao Paulo",
				Metadata: map[string]any{
					"population": 22.6,
					"area":       1523,
				},
			},
		{
		  PageContent: "Moscow",
		  Metadata: map[string]any{
					"population": 22.6,
					"area":       1524,
				},
		},
		{
		  PageContent: "Novosibirsk",
		  Metadata: map[string]any{
					"population": 22.6,
					"area":       1525,
				},
		},
		{
		  PageContent: "Ufa, Russia",
		  Metadata: map[string]any{
					"population": 22.6,
					"area":       1526,
				},
		},
		{
			PageContent: "Krasnoyarsk",
			Metadata: map[string]any{
				"population": 22.6,
				"area":       1527,
			},
		},
		{
			PageContent: "Petrozavodsk",
			Metadata: map[string]any{
				"population": 22.6,
				"area":       1528,
			},
		},
		{
			PageContent: "Belgorod",
			Metadata: map[string]any{
				"population": 22.6,
				"area":       1529,
			},
		}}

		return docs
}
*/


func LoadDocsToStore(docs []schema.Document)  {
	fmt.Println("loading data from")

	store, err := GetVectorStore()

	if err != nil {
		log.Panic(err)
	}
	//docs := getSampleDocs()

	fmt.Println("no. of documents to be loaded", len(docs))

	_, err = store.AddDocuments(context.Background(), docs)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("data successfully loaded into vector store")

	log.Println(err)
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