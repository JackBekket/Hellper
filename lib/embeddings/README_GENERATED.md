# embeddings

This package provides functions for loading documents into a vector store, performing semantic search, and running a retrieval QA chain using an OpenAI model.

## Environment variables, flags, cmdline arguments, files and their paths that can be used for configuration

- `ai_url`: URL of the OpenAI API
- `api_token`: API token for the OpenAI API
- `db_link`: Database connection string for the vector store
- `name`: Name of the collection in the vector store

## Edge cases of how application can be launched

- No specific edge cases mentioned in the code.

## Project package structure

```
lib/embeddings/
    query.go
    load.go
    common.go
```

## Major code parts

### Rag function

- Takes an OpenAI API URL, API token, a question, the number of results to return, and a vector store as input.
- Initializes an OpenAI client using the provided API URL, token, and model.
- Creates a retrieval QA chain using the OpenAI client and the vector store.
- Runs the retrieval QA chain with the given question and a maximum of 2048 tokens.
- Returns the final answer and any error.

### SemanticSearch function

- Takes a search query, the maximum number of results to return, a vector store, and optional vector store options as input.
- Performs a similarity search on the vector store using the given search query and maximum results.
- Prints the similarity search results, including the page content and score for each document.
- Returns the search results and any error.

### LoadDocsToStore function

- Takes a slice of documents and a vector store as input.
- Prints the number of documents to be loaded.
- Adds the documents to the vector store.
- Prints a success message or any error.

### getDocs function

- Takes a source URL as input.
- Sends a GET request to the given source.
- Loads and splits the response body using document loaders and text splitters.
- Returns the loaded documents and any error.

### LoadEnv function

- Unimplemented.

### GetVectorStore function

- Takes an OpenAI API URL, API token, and database connection string as input.
- Creates a pgxpool.Pool from the database connection string.
- Creates an embeddings.Embedder using the OpenAI API with the given API token and base URL.
- Creates a pgvector.VectorStore using the Embedder and the pgxpool.Pool.
- Returns the VectorStore and any error.

### GetVectorStoreWithOptions function

- Takes an OpenAI API URL, API token, database connection string, and a collection name as input.
- Creates a pgxpool.Pool from the database connection string.
- Creates an embeddings.Embedder using the OpenAI API with the given API token and base URL.
- Creates a pgvector.VectorStore using the Embedder, the pgxpool.Pool, and the given collection name.
- Returns the VectorStore and any error.

