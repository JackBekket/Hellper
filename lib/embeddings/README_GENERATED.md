# Package: embeddings

### Imports:

```
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
```

### External Data, Input Sources:

- `ai_url`: AI URL (localhost or openai or docker)
- `api_token`: AI token
- `db_link`: Database link
- `name`: Collection name for vector store

### Code Summary:

#### LoadEnv()

This function is not used in the provided code.

#### GetVectorStore()

This function creates a vector store from the database. It takes the AI URL, API token, and database link as input. It first parses the database link and creates a connection pool. Then, it creates an embeddings client using the OpenAI API and an embedder using the embeddings client. Finally, it creates a vector store using the pgvector library, which uses the connection pool and embedder.

#### GetVectorStoreWithOptions()

This function is similar to GetVectorStore() but also takes a collection name as input. It creates a vector store with the specified collection name.

The code provides two functions to create a vector store from the database. The first function takes the AI URL, API token, and database link as input, while the second function also takes a collection name as input. Both functions create a vector store using the pgvector library, which uses the connection pool and embedder to store and retrieve embeddings.

lib/embeddings/load.go
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

### External Data, Input Sources:

- The code uses an external vector store to store the documents.
- It also uses an external source to retrieve the documents, which can be a URL or any other data source.

### Code Summary:

#### LoadDocsToStore Function:

This function takes a list of documents and a vector store as input. It first prints the number of documents to be loaded and then adds the documents to the vector store using the AddDocuments method. If there is an error during the process, it logs the error and panics.

#### getDocs Function:

This function takes a URL as input and retrieves the documents from the specified source. It uses the http.Get function to fetch the content from the URL and then uses the documentloaders.NewHTML function to load and split the content into documents. The textsplitter.NewRecursiveCharacter function is used to split the content into individual documents. If there is an error during the process, it returns an error.

#### Other Code:

There are some commented-out code sections that were not included in the final code. These sections include a function to get text documents from a file and a function to get a vector store.



lib/embeddings/query.go
## Package: embeddings

### Imports:

```
"context"
"fmt"
"github.com/tmc/langchaingo/chains"
"github.com/tmc/langchaingo/llms/openai"
"github.com/tmc/langchaingo/schema"
"github.com/tmc/langchaingo/vectorstores"
```

### External Data, Input Sources:

- `ai_url`: URL of the AI service (e.g., OpenAI API).
- `api_token`: API token for the AI service.
- `question`: The question to be answered.
- `numOfResults`: The number of results to return.
- `store`: A vector store to store and retrieve embeddings.
- `searchQuery`: The query for semantic search.
- `maxResults`: The maximum number of results to return for semantic search.
- `options`: Additional options for the vector store.

### Code Summary:

#### Rag Function:

The `Rag` function performs a retrieval-augmented generation (RAG) task using a language model (LLM) and a vector store. It takes the AI service URL, API token, question, number of results, and vector store as input. The function first creates an embeddings client using the provided AI service URL and API token. Then, it runs a retrieval-augmented generation chain using the LLM and the vector store to generate a response to the question. Finally, it returns the generated response and any errors encountered during the process.

#### SemanticSearch Function:

The `SemanticSearch` function performs a semantic search using a vector store. It takes the search query, maximum number of results, vector store, and additional options as input. The function first retrieves the vector store (if not provided) and then performs a similarity search using the provided search query and maximum number of results. Finally, it returns the search results and any errors encountered during the process.

