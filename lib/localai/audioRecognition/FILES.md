# lib/localai/audioRecognition/stt.go  
# Package Name and Imports  
The package name is `stt`. The imports are:  
* `fmt`  
* `io`  
* `log`  
* `net/http`  
* `os`  
* `path/filepath`  
* `tgbotapi` from `github.com/go-telegram-bot-api/telegram-bot-api/v5`  
  
## External Data and Input Sources  
The external data and input sources are:  
* `updateMessage` of type `*tgbotapi.Message`  
* `bot` of type `tgbotapi.BotAPI`  
* Environment variables:  
	+ `AI_ENDPOINT`  
	+ `VOICE_RECOGNITION_ENDPOINT` (optional, defaults to `/v1/audio/transcriptions`)  
	+ `VOICE_RECOGNITION_MODEL` (optional, defaults to `whisper-1`)  
  
## TODO Comments  
There are no TODO comments in the provided code.  
  
## Code Summary  
### HandleVoiceMessage Function  
The `HandleVoiceMessage` function takes an `updateMessage` and a `bot` as input, and returns the local file path of the downloaded audio file and an error. It gets the file URL from the `GetFileURL` function, downloads the file using the `DownloadFile` function, and returns the local file path.  
  
### GetFileURL Function  
The `GetFileURL` function takes a `fileID` and a `bot` as input, and returns the file URL and an error. It uses the `bot` to get the file configuration and then constructs the file URL.  
  
### DownloadFile Function  
The `DownloadFile` function takes a `url` and a `localFilePath` as input, and returns an error. It creates the file, gets the data from the URL, and writes the body to the file.  
  
### GetEnvsForSST Function  
The `GetEnvsForSST` function returns the URL and model for speech-to-text recognition. It gets the URL and model from environment variables, with default values if the variables are not set.  
  
  
  
