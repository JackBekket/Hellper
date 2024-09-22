# embeddings

This package provides functions for loading documents into a vector store, querying the store for relevant documents, and performing semantic search. It utilizes the OpenAI API for embeddings and the pgvector library for storing and retrieving vectors.

## File Structure

```
lib/embeddings/
  common.go
  load.go
  query.go
```

## Code Summary

### common.go

- The `LoadEnv()` function is not implemented.
- The `GetVectorStore()` function takes an AI service URL, API token, and database link as input. It creates a connection to the database, an embeddings client using the OpenAI API, and a vector store using the pgvector library. The vector store is connected to the database and uses the embeddings client. The function returns the vector store and any errors encountered.
- The `GetVectorStoreWithOptions()` function is similar to `GetVectorStore()`, but it also takes a collection name as input. It creates a vector store using the pgvector library, connecting to the database, using the embeddings client, and specifying the collection name. The function returns the vector store and any errors encountered.

### load.go

- The `LoadDocsToStore()` function takes a list of documents and a vector store as input. It prints the number of documents to be loaded, adds the documents to the vector store, and prints a success message. It also logs any errors encountered.
- The `getDocs()` function takes a source string as input. It fetches data from the given source, loads and splits the data into documents, and returns the documents and any errors.

### query.go

- The `Rag()` function takes an AI service URL, API token, question, number of results, and a vector store as input. It sets the base URL to the AI service URL, creates an embeddings client using the provided parameters, runs a retrieval QA chain using the embeddings client and the provided store, prints the final answer, and returns the result and any error.
- The `SemanticSearch()` function takes a search query, maximum number of results, a vector store, and optional options as input. It performs a similarity search using the provided store and options, prints the similarity search results, and returns the search results and any error.

