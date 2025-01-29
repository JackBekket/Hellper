# lib/localai/audioRecognition/stt.go  
**Package Name:** stt  
  
**Imports:**  
  
* `fmt`  
* `io`  
* `log`  
* `net/http`  
* `os`  
* `path/filepath`  
* `github.com/go-telegram-bot-api/telegram-bot-api/v5` (tgbotapi)  
  
**External Data/Inputs:**  
  
* `updateMessage` ( Telegram message object)  
* `bot` (Telegram bot API object)  
* Environment variables:  
	+ `AI_ENDPOINT`  
	+ `VOICE_RECOGNITION_SUFFIX`  
	+ `VOICE_RECOGNITION_MODEL`  
  
**TODO Comments:**  
  
* None found  
  
**Summary:**  
  
### Main Functionality  
  
The `stt` package provides a function `HandleVoiceMessage` that handles voice messages sent to a Telegram bot. It retrieves the file URL of the voice message, downloads the file, and returns the local file path.  
  
### File Retrieval and Download  
  
The `GetFileURL` function retrieves the file URL of a Telegram file given its ID. It constructs the URL by combining the bot token and the file path.  
  
The `DownloadFile` function downloads a file from a given URL to a local file path.  
  
### Environment Variables  
  
The `GetEnvsForSST` function retrieves environment variables for SST (Speech-to-Text) functionality. It sets default values for `VOICE_RECOGNITION_SUFFIX` and `VOICE_RECOGNITION_MODEL` if they are not set.  
  
  
  
