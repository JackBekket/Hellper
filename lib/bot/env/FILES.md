# lib/bot/env/newEvn.go  
# Package Name and Imports  
The package name is `env`. The imports are:  
* `errors`  
* `github.com/joho/godotenv`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `.env` file located at `envLoc` constant  
* Environment variables stored in the `env` map  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### Load Function  
The `Load` function reads the environment variables from the `.env` file using the `godotenv.Read` function. If an error occurs, it returns the error.  
  
### Token Loading Functions  
The package provides several functions to load tokens:  
* `LoadTGToken`: loads the Telegram token from the `env` map and returns it as a string. If the token is not found, it returns an error.  
* `LoadLocalPD`: loads the local password from the `env` map and returns it as a string.  
* `LoadLocalAI_Endpoint`: loads the AI endpoint from the `env` map and returns it as a string.  
* `GetAdminToken`: loads the admin token from the `env` map and returns it as a string.  
  
### Data Structures  
The package defines a struct `AdminData` with two fields: `ID` and `GPTKey`.  
  
  
  
