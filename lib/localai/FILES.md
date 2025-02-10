# lib/localai/localai.go  
**Package Name:** localai  
  
**Imports:**  
  
* `bytes`  
* `encoding/json`  
* `fmt`  
* `io`  
* `io/ioutil`  
* `log`  
* `mime/multipart`  
* `os`  
* `path/filepath`  
* `strings`  
  
**External Data/Inputs:**  
  
* `prompt` (string) in `main` function  
* `modelName` (string) in `main` function  
* `url` (string) in `main` function  
* `size` (string) in `GenerateImageStableDiffusion` function  
* `path` (string) in `TranscribeWhisper` function  
* `OPENAI_API_KEY` (environment variable) in `GenerateImageStableDiffusion` and `TranscribeWhisper` functions  
  
**TODO Comments:**  
  
* None found in the provided code  
  
**Summary:**  
  
### Main Function  
The `main` function demonstrates the usage of the `localai` package by generating a chat completion using the `GenerateCompletion` function.  
  
### GenerateCompletion Function  
This function sends a POST request to the specified `url` with a JSON body containing a `ChatRequest` struct. It then parses the JSON response and returns the first choice's message content.  
  
### GenerateImageStableDiffusion Function  
This function generates an image using the Stable Diffusion model. It creates a multipart form and adds a file field with the specified `path`. It then sends a POST request to the specified `url` with the form data and returns the generated image URL.  
  
### TranscribeWhisper Function  
This function transcribes an audio file using the Whisper model. It creates a multipart form, adds a file field with the specified `path`, and sends a POST request to the specified `url` with the form data. It then unmarshals the JSON response and returns the transcribed text.  
  
**  
  
