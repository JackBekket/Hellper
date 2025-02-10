# embeddings
## Summary
The embeddings package provides functions for loading and querying vector stores. It uses environment variables such as `PG_HOST`, `PG_USER`, `PG_PASSWORD`, `PG_DB`, and `API_KEY` for configuration. The package has three main files: `common.go`, `load.go`, and `query.go`.

## Environment Variables and Flags
* `PG_HOST`: the host of the PostgreSQL database
* `PG_USER`: the user of the PostgreSQL database
* `PG_PASSWORD`: the password of the PostgreSQL database
* `PG_DB`: the name of the PostgreSQL database
* `API_KEY`: the API key for the OpenAI service

## Command Line Arguments
The package can be launched with the following command line arguments:
* `ai_url`: the URL of the AI service
* `api_token`: the token for the AI service
* `db_link`: the link to the database
* `name`: the name of the vector store

## Edge Cases
The package can be launched in the following ways:
* With a database link and AI URL
* With a database link, AI URL, and API token
* With a database link, AI URL, API token, and name

## Project Package Structure
The project package structure is as follows:
* `embeddings`
	+ `common.go`
	+ `load.go`
	+ `query.go`
* `lib`
	+ `embeddings`
		- `common.go`
		- `load.go`
		- `query.go`

## Relations Between Code Entities
The `common.go` file provides functions for getting a vector store from a database. The `load.go` file provides functions for loading documents into a vector store. The `query.go` file provides functions for querying a vector store.

## Code Explanation
The `GetVectorStore` function in `common.go` takes `ai_url`, `api_token`, and `db_link` as parameters and returns a `vectorstores.VectorStore` and an error. The `LoadDocsToStore` function in `load.go` loads a slice of `schema.Document` objects into a `vectorstores.VectorStore` object. The `Rag` function in `query.go` takes in several parameters, including the AI URL, API token, question, number of results, and vector store, and returns the result and an error.

