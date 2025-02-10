# Package: stt

### Imports:

- fmt
- io
- log
- net/http
- os
- path/filepath
- github.com/go-telegram-bot-api/telegram-bot-api/v5

### External Data, Input Sources:

- os.Getenv("AI_ENDPOINT") - AI endpoint URL
- os.Getenv("VOICE_RECOGNITION_SUFFIX") - URL suffix for voice recognition
- os.Getenv("VOICE_RECOGNITION_MODEL") - Voice recognition model

### TODOs:

- None

### Code Summary:

#### HandleVoiceMessage:

This function handles voice messages received from Telegram. It first extracts the file ID of the voice message and then uses the GetFileURL function to retrieve the URL of the voice file. The file is then downloaded to a local file using the DownloadFile function. Finally, the function returns the local file path and any error encountered during the process.

#### GetFileURL:

This function retrieves the URL of a Telegram file given its file ID and the bot API. It uses the bot API to get the file information and constructs the URL using the bot token and file path. The function returns the file URL and any error encountered during the process.

#### DownloadFile:

This function downloads a file from a given URL to a local file path. It first creates the local file and then uses the http.Get function to retrieve the file data. The file data is then written to the local file using the io.Copy function. The function returns any error encountered during the process.

#### GetEnvsForSST:

This function retrieves the AI endpoint URL and voice recognition model from environment variables. If the environment variables are not set, it uses default values for the URL and model. The function returns the AI endpoint URL and the voice recognition model.

