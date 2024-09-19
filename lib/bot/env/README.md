# env

This package provides functions to load environment variables from a .env file and access them as needed. It also includes functions to load admin data, telegram token, local password, and local AI endpoint from environment variables.

## Project Package Structure

```
lib/bot/env/newEvn.go
```

## Code Summary

### Load()

This function reads environment variables from the .env file and returns an error if reading fails.

### LoadAdminData()

This function creates a map to store admin data and iterates through environment variables. It parses admin IDs and GPT keys from the environment variables and returns a map with admin data.

### LoadTGToken()

This function returns the telegram token from the environment variable. It returns an error if the token is not found.

### LoadLocalPD()

This function returns the localhost password from the environment variable.

### LoadLocalAI_Endpoint()

This function returns the local AI endpoint from the environment variable.

### GetAdminToken()

This function returns the admin key from the environment variable.

