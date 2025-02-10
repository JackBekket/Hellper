# lib/bot/env/newEvn.go  
# Package Name and Imports  
The package name is `env`. The imports are:  
* `errors`  
* `log`  
* `strconv`  
* `github.com/joho/godotenv`  
  
## External Data and Input Sources  
The external data sources are:  
* `.env` file, which contains environment settings  
* The following environment variables are loaded from the `.env` file:  
	+ `ADMIN_ID`  
	+ `ADMIN_KEY`  
	+ `MINTY_ID`  
	+ `MINTY_KEY`  
	+ `OK_ID`  
	+ `OK_KEY`  
	+ `MURS_ID`  
	+ `MURS_KEY`  
	+ `TG_KEY`  
	+ `LOCALHOST_PWD`  
	+ `AI_ENDPOINT`  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Load Function  
The `Load` function reads the environment variables from the `.env` file using the `godotenv` package. If there is an error, it returns the error.  
  
### LoadAdminData Function  
The `LoadAdminData` function loads the admin data from the environment variables and returns a map of admin data. It parses the admin IDs and GPT keys from the environment variables and stores them in the `AdminData` struct.  
  
### LoadTGToken Function  
The `LoadTGToken` function returns the Telegram token from the environment variables. If the token is not found, it returns an error.  
  
### LoadLocalPD Function  
The `LoadLocalPD` function returns the local password from the environment variables.  
  
### LoadLocalAI_Endpoint Function  
The `LoadLocalAI_Endpoint` function returns the AI endpoint from the environment variables.  
  
### GetAdminToken Function  
The `GetAdminToken` function returns the admin token from the environment variables.  
  
  
  
