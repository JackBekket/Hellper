# localai
The localai package provides functionality for audio and image recognition, as well as text generation using OpenAI models. The package is written in Go and utilizes various libraries such as `net/http` and `encoding/json`.

## Environment Variables and Configuration
The package uses the following environment variables:
- `OPENAI_API_KEY`
File paths and URLs can be used for configuration, such as:
- Audio file path for transcription
- URL for API requests (e.g., `http://localhost:8080/v1/chat/completions`)

## Edge Cases for Launching the Application
The application can be launched in the following ways:
- Using the `localai.go` file as the main entry point
- Using the `audioRecognition/stt.go` file for audio transcription
- Using the `imageRecognition/imageRecognition.go` file for image recognition

## Project Package Structure
The project package structure is as follows:
- localai/
  - audioRecognition/
    - stt.go
  - imageRecognition/
    - imageRecognition.go
  - localai.go
- lib/
  - localai/
    - localai.go

## Relations Between Code Entities
The package defines several data structures, such as `ChatRequest` and `Message`, which are used in functions like `GetModelsList` and `GenerateCompletion`. The `TranscribeWhisper` function is used for audio transcription, while the `GenerateImageStableDiffusion` function is used for image generation. The `cleanText` function is used to clean the text output from transcription.

The package also defines several other functions, including `GenerateCompletion` and `GetModelsList`, which are used to interact with OpenAI models. The `OpenAIDataObject` and `OpenAIModelsResponse` data structures are used to represent OpenAI data and responses.

Overall, the package provides a range of functionality for audio and image recognition, as well as text generation using OpenAI models.
