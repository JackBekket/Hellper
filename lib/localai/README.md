# localai

The `localai` package provides a set of functions for generating chat completions, images, and transcribing audio files using various AI models.

**Configuration:**

* `OPENAI_API_KEY`: an environment variable required for generating images and transcribing audio files

**Usage:**

The package can be launched from the command line by running the `main` function with the following flags:
```
go run localai.go -prompt <prompt> -modelName <modelName> -url <url>
```
Alternatively, the package can be launched by importing the `localai` package and calling the `main` function programmatically.

**Project Structure:**
```
localai/
main.go
audioRecognition/
stt.go
imageRecognition/
imageRecognition.go
lib/
localai/
localai.go
```
**Summary:**

The `localai` package provides three main functions: `GenerateCompletion`, `GenerateImageStableDiffusion`, and `TranscribeWhisper`. The `main` function demonstrates the usage of these functions by generating a chat completion.

The `GenerateCompletion` function sends a POST request to a specified URL with a JSON body containing a `ChatRequest` struct, and returns the first choice's message content.

The `GenerateImageStableDiffusion` function generates an image using the Stable Diffusion model by sending a POST request to a specified URL with a multipart form containing a file field with a specified path.

The `TranscribeWhisper` function transcribes an audio file using the Whisper model by sending a POST request to a specified URL with a multipart form containing a file field with a specified path.

**Edge Cases:**

* The package does not handle errors or edge cases, such as invalid input or network errors.
* The `main` function does not handle command-line arguments or flags.

**Notes:**

* The code does not contain any TODO comments, indicating that the functionality is complete and does not require further development.
* The code does not contain any unclear or dead code.

