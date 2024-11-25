# lib/localai/audioRecognition/stt.go  
## Package: stt  
  
### Imports:  
  
```  
fmt  
io  
log  
net/http  
os  
path/filepath  
github.com/go-telegram-bot-api/telegram-bot-api/v5  
```  
  
### External Data, Input Sources:  
  
1. Environment variables:  
    - `AI_ENDPOINT`: URL of the AI endpoint for speech-to-text conversion.  
    - `VOICE_RECOGNITION_SUFFIX`: Suffix to append to the AI endpoint URL for voice recognition. Defaults to `/v1/audio/transcriptions` if not set.  
    - `VOICE_RECOGNITION_MODEL`: Model to use for speech-to-text conversion. Defaults to "whisper-1" if not set.  
  
2. Telegram bot API:  
    - `tgbotapi.BotAPI`: Used to interact with the Telegram bot API.  
  
### Code Summary:  
  
#### HandleVoiceMessage:  
  
This function handles voice messages received from the Telegram bot. It first extracts the file ID of the voice message and then uses the `GetFileURL` function to retrieve the URL of the voice file. The file is then downloaded to a local file using the `DownloadFile` function. Finally, the function returns the local file path and any errors encountered during the process.  
  
#### GetFileURL:  
  
This function takes the file ID of a Telegram file and the bot API as input. It uses the bot API to retrieve the file information and constructs the URL to access the file from Telegram's servers. The function returns the file URL and any errors encountered during the process.  
  
#### DownloadFile:  
  
This function takes a URL and a local file path as input. It creates a new file at the specified local path and downloads the data from the given URL to the file. The function returns any errors encountered during the process.  
  
#### GetEnvsForSST:  
  
This function retrieves the necessary environment variables for speech-to-text conversion. It first retrieves the AI endpoint URL from the `AI_ENDPOINT` environment variable. If the `VOICE_RECOGNITION_SUFFIX` environment variable is not set, it defaults to `/v1/audio/transcriptions`. The function also retrieves the voice recognition model from the `VOICE_RECOGNITION_MODEL` environment variable, defaulting to "whisper-1" if not set. Finally, the function returns the constructed AI endpoint URL and the voice recognition model.  
  
  
  
