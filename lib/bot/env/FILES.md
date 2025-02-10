# lib/bot/env/newEvn.go  
**Package Name:** env  
  
**Imports:**  
  
* `errors`  
* `log`  
* `strconv`  
* `github.com/joho/godotenv`  
  
**External Data/Inputs:**  
  
* `.env` file (used to load environment settings)  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Load Function  
  
The `Load` function reads the `.env` file and populates the `env` map. If there's an error during the read process, it returns an error.  
  
### LoadAdminData Function  
  
This function loads admin data from the `env` map. It iterates over the map and parses the values as IDs and GPT keys. It returns a map of admin data.  
  
### LoadTGToken, LoadLocalPD, LoadLocalAI_Endpoint, and GetAdminToken Functions  
  
These functions return environment variables from the `env` map. They check if the variables exist and return an error if they don't.  
  
**  
  
