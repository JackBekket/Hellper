# embeddings

## Summary

This code package provides functions for loading environment variables and creating vector stores for embedding text data. It utilizes the LangChain library for embedding text data and the OpenAI API for generating embeddings. The package also includes functions for creating vector stores using PostgreSQL with the pgvector library.

The `GetVectorStore` function takes the AI URL, API token, and database link as input and returns a vector store. It first initializes the OpenAI client using the provided API token and base URL. Then, it creates an embedding client using the OpenAI API and a vector store using the pgvector library. The vector store is configured to use the provided database connection and embedding client.

The `GetVectorStoreWithOptions` function is similar to `GetVectorStore` but allows for specifying a collection name for the vector store. This function also takes the AI URL, API token, database link, and collection name as input and returns a vector store.

In summary, this code package provides a way to create vector stores for embedding text data using the OpenAI API and PostgreSQL. It includes functions for loading environment variables, initializing the OpenAI client, creating an embedding client, and creating a vector store with the pgvector library. The package also allows for specifying a collection name for the vector store.



## Summary

This code package provides functions for loading and managing documents in a vector store. It includes functions for loading documents from various sources, such as HTML files or URLs, and storing them in a vector store. The package also includes utilities for splitting text into documents and extracting metadata from documents.

The `LoadDocsToStore` function takes a list of documents and a vector store as input and loads the documents into the vector store. It first prints the number of documents to be loaded and then calls the `AddDocuments` method of the vector store to add the documents. If an error occurs during the process, it logs the error and panics.

The `getDocs` function takes a source URL as input and returns a list of documents and an error. It first sends an HTTP GET request to the source URL and then uses the `documentloaders.NewHTML` function to load and split the HTML content into documents. The `textsplitter.NewRecursiveCharacter` function is used to split the text into documents based on character boundaries.

In summary, this code package provides a set of functions for loading and managing documents in a vector store, including utilities for splitting text into documents and extracting metadata from documents. The package also includes functions for loading documents from various sources, such as HTML files or URLs.



## Summary

This code package provides functions for semantic search and question answering using large language models (LLMs) and vector stores. The `Rag` function takes an AI URL, API token, question, number of results, and a vector store as input. It creates an embeddings client using the provided information and runs a retrieval-based question answering chain using the LLM and vector store. The function returns the final answer and any errors encountered.

The `SemanticSearch` function performs a similarity search on a given search query using a vector store. It takes the search query, maximum number of results, and a vector store as input. The function returns a list of documents with their scores and any errors encountered.

Both functions utilize vector stores to efficiently store and retrieve embeddings of text data. This allows for fast and accurate semantic search and question answering. The package also includes options for customizing the search process, such as specifying the maximum number of results and additional vector store options.

In summary, this code package provides a robust and flexible framework for semantic search and question answering using LLMs and vector stores. It offers a simple and efficient way to perform these tasks, allowing users to leverage the power of LLMs for various applications.



