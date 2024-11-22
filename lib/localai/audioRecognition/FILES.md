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
  
This function retrieves the URL of a file given its file ID and the Telegram bot API. It uses the `tgbotapi.FileConfig` struct to specify the file ID and then calls the `bot.GetFile` method to retrieve the file information. The function then constructs the file URL using the bot token and file path and returns it along with any errors encountered.  
  
#### DownloadFile:  
  
This function downloads a file from a given URL to a local file path. It first creates the local file and then uses the `http.Get` method to retrieve the file data from the URL. The file data is then written to the local file using the `io.Copy` function. Finally, the function returns any errors encountered during the process.  
  
#### GetEnvsForSST:  
  
This function retrieves the AI endpoint URL and voice recognition model from environment variables. It first retrieves the AI endpoint URL from the `AI_ENDPOINT` environment variable and appends the voice recognition suffix from the `VOICE_RECOGNITION_SUFFIX` environment variable. If the suffix is not set, it defaults to `/v1/audio/transcriptions`. The function then retrieves the voice recognition model from the `VOICE_RECOGNITION_MODEL` environment variable and defaults to "whisper-1" if not set. Finally, it returns the AI endpoint URL and voice recognition model.  
  
  
  
