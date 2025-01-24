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
  
1. Environment variables:  
    - `PG_HOST`  
    - `PG_USER`  
    - `PG_PASSWORD`  
    - `PG_DB`  
    - `API_KEY`  
  
2. Function arguments:  
    - `ai_url`: AI URL (localhost, OpenAI, or Docker)  
    - `api_token`: AI token  
    - `db_link`: Database link  
    - `name`: Collection name for vector store  
  
### Code Summary:  
  
#### LoadEnv() function:  
  
This function is not implemented in the provided code.  
  
#### GetVectorStore() function:  
  
This function creates a vector store from a database using the provided AI URL, API token, and database link. It first parses the database link and creates a connection pool. Then, it creates an embeddings client using the OpenAI API and an embedder using the embeddings client. Finally, it creates a vector store using the pgvector library, which uses the connection pool and embedder.  
  
#### GetVectorStoreWithOptions() function:  
  
This function is similar to GetVectorStore() but allows specifying a collection name for the vector store. It takes the same arguments as GetVectorStore() plus an additional argument, `name`, which specifies the collection name. The rest of the logic is the same as GetVectorStore().  
  
  
  
# lib/embeddings/load.go  
## Package: embeddings  
  
### Imports:  
  
```  
"context"  
"fmt"  
"net/http"  
"log"  
"github.com/tmc/langchaingo/documentloaders"  
"github.com/tmc/langchaingo/schema"  
"github.com/tmc/langchaingo/textsplitter"  
"github.com/tmc/langchaingo/vectorstores"  
"github.com/tmc/langchaingo/vectorstores/pgvector"  
```  
  
### External Data, Input Sources:  
  
1. `source` string: This is used to fetch documents from a given URL.  
  
### Code Summary:  
  
#### LoadDocsToStore Function:  
  
This function takes a slice of `schema.Document` and a `vectorstores.VectorStore` as input. It first prints the number of documents to be loaded and then adds the documents to the vector store using the `AddDocuments` method. If there is an error during the process, it logs the error and panics. After the documents are loaded, it closes the vector store using the `Close` method.  
  
#### getDocs Function:  
  
This function takes a `source` string as input and returns a slice of `schema.Document` and an error. It first fetches the content from the given URL using `http.Get`. Then, it loads and splits the content into documents using the `documentloaders.NewHTML` and `textsplitter.NewRecursiveCharacter` functions. Finally, it returns the slice of documents and any error encountered during the process.  
  
  
  
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
  
1. `ai_url`: URL of the AI service (e.g., OpenAI API).  
2. `api_token`: API token for authentication with the AI service.  
3. `question`: The question to be answered or the query for semantic search.  
4. `numOfResults`: The number of results to return for the retrieval-based QA.  
5. `store`: A vector store to store and retrieve embeddings.  
6. `option`: Additional options for the vector store.  
7. `searchQuery`: The query for semantic search.  
8. `maxResults`: The maximum number of results to return for semantic search.  
  
### Code Summary:  
  
#### Rag Function:  
  
This function performs retrieval-augmented generation (RAG) using an LLM and a vector store. It first creates an embeddings client using the provided AI URL and API token. Then, it runs a retrieval-based QA chain using the embeddings client and the vector store. The chain takes the question as input and returns the answer. Finally, it prints the final answer and closes the vector store.  
  
#### SemanticSearch Function:  
  
This function performs semantic search using a vector store. It takes a search query, maximum number of results, and a vector store as input. It then performs a similarity search on the vector store using the provided query and returns the top `maxResults` results. The function also prints the similarity search results, including the page content and score for each result. Finally, it closes the vector store.  
  
  
  
