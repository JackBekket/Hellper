# lib/bot/env/newEvn.go  
## Package: env  
  
### Imports:  
  
* errors  
* log  
* strconv  
* github.com/joho/godotenv  
  
### External Data, Input Sources:  
  
* .env file: Contains environment variables that are loaded and used by the package.  
  
### Code Summary:  
  
#### Load() function:  
  
This function loads environment variables from the .env file using the godotenv library. It returns an error if there is an issue loading the environment variables.  
  
#### LoadAdminData() function:  
  
This function parses the loaded environment variables and creates a map of AdminData structs. Each AdminData struct contains an ID and a GPTKey. The function handles parsing the ID values from the environment variables and populates the AdminData structs accordingly.  
  
#### LoadTGToken() function:  
  
This function retrieves the Telegram token from the environment variables and returns it as a string. It returns an error if the Telegram token is not found in the .env file.  
  
#### LoadLocalPD() function:  
  
This function retrieves the Localhost password from the environment variables and returns it as a string.  
  
#### LoadLocalAI_Endpoint() function:  
  
This function retrieves the Local AI endpoint from the environment variables and returns it as a string.  
  
#### GetAdminToken() function:  
  
This function retrieves the Admin key from the environment variables and returns it as a string.  
  
  
  
