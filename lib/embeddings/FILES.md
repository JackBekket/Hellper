# lib/embeddings/common.go  
## Package: embeddings  
  
### Imports:  
  
```  
"context"  
"fmt"  
"log"  
"github.com/tmc/langchaingo/embeddings"  
"github.com/tmc/langchaingo/llms/openai"  
"github.com/tmc/langchaingo/vectorstores"  
"github.com/tmc/langchaingo/vectorstores/pgvector"  
"github.com/jackc/pgx/v5/pgxpool"  
```  
  
### External Data, Input Sources:  
  
- `ai_url`: AI url (localhost or openai or docker)  
- `api_token`: AI token  
- `db_link`: database link  
  
### TODOs:  
  
- None found  
  
### Summary:  
  
The `embeddings` package provides functions for loading environment variables and creating vector stores. The `GetVectorStore` function takes an AI url, API token, and database link as input and returns a vector store. It first parses the database link to create a connection pool. Then, it creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and token. An embedder is created using the embeddings client. Finally, a vector store is created using the connection pool and embedder.  
  
The `GetVectorStoreWithOptions` function is similar to `GetVectorStore` but also takes a collection name as input. It creates a vector store with the specified collection name.  
  
# lib/embeddings/load.go  
## Package: embeddings  
  
### Imports:  
  
- context  
- fmt  
- net/http  
- log  
- github.com/tmc/langchaingo/documentloaders  
- github.com/tmc/langchaingo/schema  
- github.com/tmc/langchaingo/textsplitter  
- github.com/tmc/langchaingo/vectorstores  
- github.com/tmc/langchaingo/vectorstores/pgvector  
  
### External Data, Input Sources:  
  
- The code uses the `http.Get` function to fetch data from a given URL.  
- It also uses the `documentloaders.NewHTML` function to load and split HTML content.  
  
### TODOs:  
  
- There is a commented-out function `GetTextDocs` that seems to be intended for loading text documents from a data source called `fileData`. This function needs to be implemented.  
  
### Summary:  
  
The code in this package focuses on loading and storing documents in a vector store. The `LoadDocsToStore` function takes a slice of `schema.Document` objects and a `vectorstores.VectorStore` as input. It first prints the number of documents to be loaded and then calls the `AddDocuments` method of the vector store to add the documents. If there is an error during this process, the function panics. After the documents are loaded, the function closes the vector store using a defer statement.  
  
The `getDocs` function is responsible for fetching documents from a given URL and loading them into a vector store. It first uses `http.Get` to retrieve the content from the URL and then uses `documentloaders.NewHTML` to load and split the HTML content. The resulting documents are then returned as a slice of `schema.Document` objects.  
  
# lib/embeddings/query.go  
## Package: embeddings  
  
### Imports:  
  
```  
"context"  
"fmt"  
"log"  
"github.com/tmc/langchaingo/chains"  
"github.com/tmc/langchaingo/llms/openai"  
"github.com/tmc/langchaingo/schema"  
"github.com/tmc/langchaingo/vectorstores"  
"github.com/tmc/langchaingo/vectorstores/pgvector"  
```  
  
### External Data, Input Sources:  
  
- `ai_url`: URL of the AI service (e.g., OpenAI API).  
- `api_token`: API token for the AI service.  
- `question`: The question to be answered.  
- `numOfResults`: Number of results to return.  
- `store`: A vector store to store and retrieve embeddings.  
- `option`: Additional options for the vector store.  
- `searchQuery`: The query for semantic search.  
- `maxResults`: Maximum number of results to return for semantic search.  
  
### TODOs:  
  
- None found.  
  
### Code Summary:  
  
#### Rag Function:  
  
This function performs a retrieval-augmented generation (RAG) task using a language model and a vector store. It first creates an embeddings client using the provided AI service URL and API token. Then, it runs a retrieval-augmented question answering chain using the embeddings client and the vector store. The chain retrieves relevant documents from the vector store based on the input question and uses the language model to generate a response. Finally, it returns the generated response and any errors encountered during the process.  
  
#### SemanticSearch Function:  
  
This function performs a semantic search using a vector store. It takes a search query, maximum number of results, and a vector store as input. It then uses the vector store to find documents that are similar to the search query and returns the top `maxResults` documents along with their scores. The function also closes the vector store connection after the search is complete.  
  
