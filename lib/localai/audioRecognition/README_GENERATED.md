# stt
The provided code is for a package named `stt` that handles voice messages and performs speech-to-text recognition using the Telegram Bot API and an external AI endpoint. 

The package uses the following environment variables:
* `AI_ENDPOINT`
* `VOICE_RECOGNITION_ENDPOINT` (optional, defaults to `/v1/audio/transcriptions`)
* `VOICE_RECOGNITION_MODEL` (optional, defaults to `whisper-1`)

The package can be launched as a command-line application with the following edge cases:
* Handling voice messages from Telegram Bot API
* Downloading audio files from Telegram Bot API
* Sending requests to the external AI endpoint for speech-to-text recognition

The project package structure is as follows:
- audioRecognition directory
  - stt.go
  - lib/localai/audioRecognition/stt.go

The code entities are related as follows:
* The `HandleVoiceMessage` function uses the `GetFileURL` and `DownloadFile` functions to download the audio file and then calls the `GetEnvsForSST` function to get the URL and model for speech-to-text recognition.
* The `GetFileURL` function uses the `bot` to get the file configuration and construct the file URL.
* The `DownloadFile` function uses the `url` and `localFilePath` to download the file.
* The `GetEnvsForSST` function uses the environment variables to get the URL and model for speech-to-text recognition.

