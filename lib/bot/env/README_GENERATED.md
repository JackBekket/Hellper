# env
The provided code is for a package named `env` that handles environment variables and external data sources. 
It has the following environment variables, flags, cmdline arguments, files and their paths that can be used for configuration: 
- `envLoc` constant for the location of the `.env` file, 
- `env` map to store environment variables.
The package can be launched as a cmd/cli/main package with the following edge cases: 
- loading environment variables from the `.env` file, 
- loading tokens such as Telegram token, local password, AI endpoint, and admin token.
The project package structure is as follows:
- newEvn.go
- lib/bot/env/newEvn.go
The relations between code entities are as follows: 
- the `Load` function reads environment variables from the `.env` file, 
- the `LoadTGToken`, `LoadLocalPD`, `LoadLocalAI_Endpoint`, and `GetAdminToken` functions load specific tokens from the `env` map, 
- the `AdminData` struct stores admin data with `ID` and `GPTKey` fields.
