## Package: env

### Imports:

```
errors
log
strconv
github.com/joho/godotenv
```

### External Data, Input Sources:

- `.env` file: Contains environment variables that are loaded and used by the package.

### Code Summary:

The package provides functions to load and retrieve environment variables from the `.env` file. It also includes functions to parse and store admin data, such as admin IDs and GPT keys, as well as functions to retrieve the Telegram token, local password, local AI endpoint, and admin token.

#### Load() function:

- Loads environment variables from the `.env` file using the `godotenv` package.
- Returns an error if there is an issue loading the environment variables.

#### LoadAdminData() function:

- Creates a map of AdminData structs, where each key is an admin identifier (e.g., "ADMIN_ID", "MINTY_ID") and the value is an AdminData struct containing the admin ID and GPT key.
- Iterates through the loaded environment variables and parses the values for each admin identifier.
- Parses the admin ID as an integer and the GPT key as a string.
- Stores the parsed data in the AdminData struct and adds it to the map.

#### LoadTGToken() function:

- Retrieves the Telegram token from the environment variables.
- Returns an error if the Telegram token is not found in the `.env` file.

#### LoadLocalPD() function:

- Retrieves the local password from the environment variables.

#### LoadLocalAI_Endpoint() function:

- Retrieves the local AI endpoint from the environment variables.

#### GetAdminToken() function:

- Retrieves the admin token from the environment variables.

### Edge Cases:

- If the `.env` file is not found or is empty, the package will return an error when attempting to load environment variables.
- If any of the required environment variables are missing, the corresponding functions will return an error.

### Project Package Structure:

```
lib/bot/env/
├── newEvn.go
```

