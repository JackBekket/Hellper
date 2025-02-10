# env Package
The `env` package is responsible for loading environment variables from a `.env` file and providing functions to access these variables. 
The package has the following imports: 
* `errors`
* `log`
* `strconv`
* `github.com/joho/godotenv`.
The package uses the following environment variables: 
* `ADMIN_ID`
* `ADMIN_KEY`
* `MINTY_ID`
* `MINTY_KEY`
* `OK_ID`
* `OK_KEY`
* `MURS_ID`
* `MURS_KEY`
* `TG_KEY`
* `LOCALHOST_PWD`
* `AI_ENDPOINT`.
The package has the following file structure:
- newEvn.go
- lib/bot/env/newEvn.go
The package has the following functions:
* `Load`: reads environment variables from the `.env` file
* `LoadAdminData`: loads admin data from environment variables
* `LoadTGToken`: returns the Telegram token from environment variables
* `LoadLocalPD`: returns the local password from environment variables
* `LoadLocalAI_Endpoint`: returns the AI endpoint from environment variables
* `GetAdminToken`: returns the admin token from environment variables.
