# Package: embeddings

### Imports:

```
context
fmt
log
github.com/tmc/langchaingo/documentloaders
github.com/tmc/langchaingo/schema
github.com/tmc/langchaingo/textsplitter
github.com/tmc/langchaingo/vectorstores
```

### External Data and Input Sources:

1. The code uses the `http.Get` function to fetch data from a given URL.
2. It utilizes the `documentloaders` package to load and split the fetched data into documents.
3. The `textsplitter` package is used to split the documents into smaller units.
4. The code also uses the `vectorstores` package to store the processed documents in a vector store.

### Code Summary:

#### LoadDocsToStore Function:

This function takes a list of documents and a vector store as input. It first prints a message indicating that data is being loaded from the given source. Then, it iterates through the list of documents and adds them to the vector store using the `AddDocuments` method. If any errors occur during the process, the function logs the error and panics. Finally, it prints a message confirming that the data has been successfully loaded into the vector store.

#### getDocs Function:

This function takes a URL as input and returns a list of documents and an error. It first uses the `http.Get` function to fetch the data from the given URL. Then, it uses the `documentloaders` package to load and split the fetched data into documents. Finally, it returns the list of documents and any errors that occurred during the process.

#### Other Code:

The code also includes a commented-out function called `GetTextDocs` which is not used in the current implementation. This function was likely intended to load documents from a local file or other data source, but it is not currently being used.

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

1. `ai_url`: URL of the AI service (e.g., OpenAI API).
2. `api_token`: API token for authentication with the AI service.
3. `question`: The query to be answered or searched for.
4. `numOfResults`: The number of results to return.
5. `store`: A vector store to store and retrieve embeddings.
6. `searchQuery`: The query for semantic search.
7. `maxResults`: The maximum number of results to return for semantic search.
8. `options`: Additional options for the vector store.

### Code Summary:

#### Rag Function:

The `Rag` function performs a retrieval-augmented generation (RAG) task using a language model (LLM) and a vector store. It takes the AI service URL, API token, question, number of results, and vector store as input. The function first creates an embeddings client using the provided AI service URL, API token, and model names. Then, it runs a retrieval-augmented generation chain using the LLM and vector store to generate a response to the question. Finally, it returns the generated response and any errors encountered during the process.

#### SemanticSearch Function:

The `SemanticSearch` function performs a semantic search using a vector store. It takes the search query, maximum number of results, vector store, and additional options as input. The function first retrieves the vector store (if not provided) and then performs a similarity search using the provided search query and maximum number of results. Finally, it returns the search results and any errors encountered during the process.



lib/embeddings/common.go
## Package: embeddings

### Imports:

```
context
fmt
log

github.com/tmc/langchaingo/embeddings
github.com/tmc/langchaingo/llms/openai
github.com/tmc/langchaingo/vectorstores
github.com/tmc/langchaingo/vectorstores/pgvector

github.com/jackc/pgx/v5/pgxpool
```

### External Data, Input Sources:

1. Environment variables:
    - `AI_BASEURL`: Base URL for the AI service (e.g., localhost, OpenAI, Docker).
    - `OPENAI_API_KEY`: API key for the OpenAI service.
    - `PG_HOST`: Hostname for the PostgreSQL database.
    - `PG_USER`: Username for the PostgreSQL database.
    - `PG_PASSWORD`: Password for the PostgreSQL database.
    - `PG_DB`: Database name for the PostgreSQL database.
2. Database connection string: `db_link` - A string containing the connection details for the PostgreSQL database.

### Code Summary:

#### LoadEnv() function:

- This function is not implemented in the provided code.

#### GetVectorStore() function:

1. Retrieves the base URL for the AI service from the `ai_url` parameter.
2. Retrieves the API token for the AI service from the `api_token` parameter.
3. Retrieves the database connection string from the `db_link` parameter.
4. Creates a PostgreSQL connection pool using the provided database connection string.
5. Creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and API token.
6. Creates an embeddings embedder using the OpenAI embeddings client.
7. Creates a vector store using the PostgreSQL connection pool and the embeddings embedder.
8. Returns the created vector store and any errors encountered during the process.

#### GetVectorStoreWithOptions() function:

1. Retrieves the base URL for the AI service from the `ai_url` parameter.
2. Retrieves the API token for the AI service from the `api_token` parameter.
3. Retrieves the database connection string from the `db_link` parameter.
4. Retrieves the name of the collection for the vector store from the `name` parameter.
5. Creates a PostgreSQL connection pool using the provided database connection string.
6. Creates an embeddings client using the OpenAI API, specifying the base URL, API version, embedding model, and API token.
7. Creates an embeddings embedder using the OpenAI embeddings client.
8. Creates a vector store using the PostgreSQL connection pool, embeddings embedder, and the specified collection name.
9. Returns the created vector store and any errors encountered during the process.



