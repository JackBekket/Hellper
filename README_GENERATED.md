# ChatBot

This project is a chatbot application built using Go and various libraries for natural language processing, embeddings, and database management.

## Configuration

Environment variables:

- `OPENAI_API_KEY`: Your OpenAI API key.
- `PGX_URL`: URL to your PostgreSQL database.
- `PGX_USER`: Username for your PostgreSQL database.
- `PGX_PASSWORD`: Password for your PostgreSQL database.
- `LOCAL_AI_ENDPOINT`: URL to your local AI endpoint (optional).

Flags:

- `-model`: GPT model to use (e.g., "gpt-3.5-turbo").
- `-language`: Language of the chatbot (e.g., "en").

## Run Instructions

1. Set the environment variables mentioned above.
2. Build the project using `go build`.
3. Run the application using `./chatbot`.

## Files and Paths

- `config.json`: Configuration file for the chatbot (optional).

## Additional Information

- The project uses libraries like LangChain, OpenAI, PGVector, and PGX.
- The chatbot can send video messages to users.
- The application is containerized using Docker.
- The project utilizes a CI/CD pipeline for automated build, test, and deployment.

