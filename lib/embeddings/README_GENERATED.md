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
    - `AI_BASEURL`: Base URL for the AI service (e.g., "http://localhost:8080/v1/").
    - `OPENAI_API_KEY`: API key for the OpenAI service.
    - `PG_HOST`: Hostname or IP address of the PostgreSQL database.
    - `PG_USER`: Username for the PostgreSQL database.
    - `PG_PASSWORD`: Password for the PostgreSQL database.
    - `PG_DB`: Database name for the PostgreSQL database.
2. Command-line arguments:
    - `ai_url`: URL for the AI service.
    - `api_token`: API token for the AI service.
    - `db_link`: Connection string for the PostgreSQL database.

### Code Summary:

#### LoadEnv() function:

- This function is not implemented in the provided code.

#### GetVectorStore() function:

- This function creates a vector store using the PostgreSQL database and the OpenAI API.
- It first retrieves the necessary environment variables or command-line arguments for the AI service, API token, and database connection.
- Then, it creates a connection pool for the PostgreSQL database using the provided connection string.
- Next, it creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and API token.
- An embedder is created using the embeddings client.
- Finally, a vector store is created using the PostgreSQL database connection, embedder, and collection name.

#### GetVectorStoreWithOptions() function:

- This function is similar to GetVectorStore() but allows specifying a collection name for the vector store.
- It takes the same input parameters as GetVectorStore() and adds an additional parameter for the collection name.
- The rest of the logic is the same as GetVectorStore(), except for the collection name being passed to the vector store creation.



lib/embeddings/load.go
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



lib/embeddings/query.go
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



