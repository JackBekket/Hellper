# lib/embeddings/common.go  
**Package Name:** embeddings  
  
**Imports:**  
  
* `context`  
* `fmt`  
* `log`  
* `github.com/tmc/langchaingo/embeddings`  
* `github.com/tmc/langchaingo/llms/openai`  
* `github.com/tmc/langchaingo/vectorstores`  
* `github.com/tmc/langchaingo/vectorstores/pgvector`  
* `github.com/jackc/pgx/v5/pgxpool`  
  
**External Data/Inputs:**  
  
* `ai_url` (string)  
* `api_token` (string)  
* `db_link` (string)  
* `name` (string)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Overview  
  
The `embeddings` package provides functionality for loading and interacting with vector stores. It appears to be a part of a larger system that utilizes OpenAI's API and a PostgreSQL database.  
  
### LoadEnv Function  
  
The `LoadEnv` function is currently empty and does not perform any actions.  
  
### GetVectorStore and GetVectorStoreWithOptions Functions  
  
These two functions are identical, except for the addition of a `name` parameter in `GetVectorStoreWithOptions`. They load a vector store from a PostgreSQL database using the `pgxpool` library. The vector store is created using the `pgvector` library, and an OpenAI client is initialized using the `openai` library. The `embeddings` library is used to create an embedder, which is then used to create the vector store.  
  
### External Dependencies  
  
The package relies on several external dependencies, including OpenAI's API, a PostgreSQL database, and the `pgxpool` library.  
  
**  
  
# lib/embeddings/load.go  
**Package Name:** embeddings  
  
**Imports:**  
  
* `context`  
* `fmt`  
* `net/http`  
* `log`  
* `github.com/tmc/langchaingo/documentloaders`  
* `github.com/tmc/langchaingo/schema`  
* `github.com/tmc/langchaingo/textsplitter`  
* `github.com/tmc/langchaingo/vectorstores`  
* `github.com/tmc/langchaingo/vectorstores/pgvector`  
  
**External Data and Input Sources:**  
  
* `fileData` (used in `GetTextDocs` function)  
* `source` (used in `getDocs` function)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Load Docs to Store  
  
The `LoadDocsToStore` function loads a list of documents (`docs`) into a vector store (`store`). It first prints a message indicating the start of the loading process, then attempts to add the documents to the store. If an error occurs, it panics with the error message. Finally, it prints a success message and closes the store.  
  
### Get Text Docs  
  
This function is not implemented and is marked as a comment (`/* */`). It appears to be intended to load text documents from an unknown source (`fileData`).  
  
### Get Docs  
  
This function retrieves documents from a URL (`source`) using the `http.Get` method. It then uses the `documentloaders` package to load and split the HTML content into documents. If an error occurs, it returns an error. Otherwise, it returns the loaded documents.  
  
**  
  
# lib/embeddings/query.go  
**Package Name:** embeddings  
  
**Imports:**  
  
* `context`  
* `fmt`  
* `log`  
* `github.com/tmc/langchaingo/chains`  
* `github.com/tmc/langchaingo/llms/openai`  
* `github.com/tmc/langchaingo/schema`  
* `github.com/tmc/langchaingo/vectorstores`  
* `github.com/tmc/langchaingo/vectorstores/pgvector`  
  
**External Data/Inputs:**  
  
* `ai_url` (string)  
* `api_token` (string)  
* `question` (string)  
* `numOfResults` (int)  
* `store` (vectorstores.VectorStore)  
* `options` (vectorstores.Option)  
* `searchQuery` (string)  
* `maxResults` (int)  
  
**TODOs:**  
  
* None found in the provided code  
  
**Summary:**  
  
### Functions  
  
#### Rag  
The `Rag` function takes in several inputs, including an AI URL, API token, question, number of results, and a vector store. It creates an OpenAI client using the provided URL and token, and then runs a retrieval QA from the LLM using the provided question and vector store. The result is printed to the console and returned.  
  
#### SemanticSearch  
The `SemanticSearch` function takes in a search query, maximum number of results, and a vector store. It performs a similarity search on the store using the provided query and returns the results. The results are then printed to the console, along with their scores.  
  
### Defer Functions  
Both `Rag` and `SemanticSearch` functions have a defer function that closes the vector store if it implements `pgvector.Store`.  
  
**  
  
