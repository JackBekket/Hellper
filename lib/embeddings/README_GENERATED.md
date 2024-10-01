# Package: embeddings

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

The package provides two main functions: `Rag` and `SemanticSearch`.

The `Rag` function performs a retrieval-augmented generation (RAG) task using a language model (LLM) and a vector store. It takes the AI service URL, API token, question, number of results, and vector store as input. The function first creates an embeddings client using the provided AI service URL and API token. Then, it runs a retrieval-augmented generation chain using the LLM and the vector store to generate a response to the question. Finally, it returns the generated response and any errors encountered during the process.

The `SemanticSearch` function performs a semantic search using a vector store. It takes the search query, maximum number of results, vector store, and additional options as input. The function first retrieves the vector store (if not provided) and then performs a similarity search using the provided search query and maximum number of results. Finally, it returns the search results and any errors encountered during the process.

The package also includes a function called `LoadDocsToStore` that takes a list of documents and a vector store as input. It adds the documents to the vector store using the AddDocuments method. If there is an error during the process, it logs the error and panics.

The package provides a way to perform RAG tasks and semantic searches using a vector store. It also includes a function to load documents into a vector store.

Project package structure:

```
embeddings/
  - common.go
  - load.go
  - query.go
```

