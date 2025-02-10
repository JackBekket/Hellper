# Package: env

### Imports:

* errors
* log
* strconv
* github.com/joho/godotenv

### External Data, Input Sources:

* .env file: Contains environment variables that are loaded into the `env` map.

### TODOs:

* None found.

### Summary:

The `env` package provides functions to load and access environment variables from a .env file. It defines a `Load()` function that reads the .env file and populates the `env` map with the environment variables. The package also provides functions to load specific environment variables, such as `LoadAdminData()`, which returns a map of admin data based on the environment variables in the .env file. Other functions include `LoadTGToken()`, `LoadLocalPD()`, `LoadLocalAI_Endpoint()`, and `GetAdminToken()`, which return specific environment variables as strings.

The package also includes a `AdminData` struct that represents the admin data, which includes an ID and a GPT key. The `LoadAdminData()` function parses the environment variables for admin data and returns a map of admin data based on the parsed values.

### Environment Variables:

* None specified.

### Flags, Cmdline Arguments:

* None specified.

### Files and Paths for Configuration:

* .env file: Contains environment variables that are loaded into the `env` map.

### Edge Cases for Launching Application:

* None specified.

### Project Package Structure:

- env/
    - newEvn.go
    - lib/bot/env/newEvn.go

