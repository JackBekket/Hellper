# hellper
## Short Summary
The provided code is for a Telegram bot written in Go, utilizing the `tgbotapi` package. It initializes the bot, sets up a database, and handles incoming updates.

## Environment Variables and Flags
* `TG_KEY`: Telegram Bot API key
* `GODOTENV`: environment variables loaded using `godotenv`

## Command Line Arguments
None

## Files and Paths
* `configuration/README.MD`: configuration README file
* `docker-compose.yaml`: Docker Compose file
* `hf_models_config/`: directory containing model configuration files
* `lib/agent/`: directory containing agent-related code
* `lib/bot/`: directory containing bot-related code
* `lib/database/`: directory containing database-related code
* `media/`: directory containing media files
* `models/`: directory containing model files
* `prompt-templates/`: directory containing prompt templates
* `tmp/`: directory containing temporary files

## Edge Cases
The application can be launched using the `go run main.go` command.

## Project Package Structure
```
hellper/
|-- .dockerignore
|-- .envExample
|-- .gitignore
|-- .vscode/
|   |-- launch.json
|-- Dockerfile
|-- configuration/
|   |-- README.MD
|-- docker-compose.yaml
|-- drafts/
|   |-- deepseeker.yaml
|-- go.mod
|-- go.sum
|-- hf_models_config/
|   |-- animagine-xl.yaml
|   |-- bert.yaml
|   |-- ...
|-- img/
|   |-- helper.jpg
|   |-- local_ai.png
|-- lib/
|   |-- agent/
|   |   |-- duck_search_agent.go
|   |   |-- semantic_search_agent.go
|   |   |-- ...
|   |-- bot/
|   |   |-- command/
|   |   |   |-- addNewUsertoMap.go
|   |   |   |-- cases.go
|   |   |   |-- ...
|   |-- database/
|   |   |-- newUserDataBase.go
|   |   |-- user.go
|   |-- ...
|-- main.go
|-- media/
|   |-- error_10.mp4
|   |-- error_11.mp4
|   |-- ...
|-- models/
|   |-- animagine-xl.yaml
|   |-- bert.yaml
|   |-- ...
|-- prompt-templates/
|   |-- alpaca.tmpl
|   |-- getting_started.tmpl
|   |-- ...
|-- tmp/
|   |-- audio/
|   |   |-- transcriptions_folder.txt
|   |-- generated/
|   |   |-- images/
|   |   |   |-- generated_images_folder.txt
|   |-- images/
|   |   |-- images_folder.txt
|-- token_speed.txt
```

## Relations Between Code Entities
The code entities are related as follows:
* The `main.go` file initializes the Telegram Bot API and sets up the database and commander.
* The `lib/agent/` package contains code for agents, such as the `duck_search_agent` and `semantic_search_agent`.
* The `lib/bot/` package contains code for the bot, including command handling and dialog management.
* The `lib/database/` package contains code for the database, including user data management.

## Unclear Places or Dead Code
There are no unclear places or dead code in the provided code.

