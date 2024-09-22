# env

This package provides functions to load environment variables from a .env file and parse them into useful data structures. It also provides functions to access specific environment variables, such as the Telegram token, local password, and local AI endpoint.

## Environment variables
- .env

## Files
- lib/bot/env/newEvn.go

## Code summary
The package defines a map called `env` to store environment variables. It also defines a struct called `AdminData` to store admin data, which includes an ID and a GPT key.

The `Load()` function reads environment variables from the .env file and returns an error if reading fails. The `LoadAdminData()` function creates a map to store admin data and iterates through environment variables to parse admin ID and GPT key for each admin. It returns a map with admin data.

The `LoadTGToken()`, `LoadLocalPD()`, and `LoadLocalAI_Endpoint()` functions return specific environment variables from the .env file, such as the Telegram token, local password, and local AI endpoint, respectively.

The `GetAdminToken()` function returns the admin key from the .env file.

