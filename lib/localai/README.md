# localai
The localai package provides a set of functions for text and image recognition, including chat completion, image generation, and audio transcription. The package uses environment variables, file paths, and URLs for API requests to configure its behavior.

## Environment Variables
* **OPENAI_API_KEY**: API key for OpenAI requests

## Flags and Cmdline Arguments
None specified

## Files and Paths
* Audio files for transcription (e.g., audioRecognition/stt.go)
* Image files for recognition (e.g., imageRecognition/imageRecognition.go)
* localai.go: main entry point of the program
* lib/localai/localai.go: library implementation of localai functions

## Edge Cases for Launch
The application can be launched in the following ways:
* Running the main function in localai.go
* Using the GenerateCompletion function with a prompt, model name, and URL
* Using the GenerateCompletionWithPWD function with a prompt, model name, URL, and password
* Using the GenerateImageStableDiffusion function with a prompt, size, URL, and model
* Using the TranscribeWhisper function with an audio file

## Project Package Structure
* localai/
	+ audioRecognition/
		- stt.go
	+ imageRecognition/
		- imageRecognition.go
	+ localai.go
	+ lib/
		- localai/
			- localai.go

## Relations between Code Entities
The package defines several data structures, including ChatRequest, Message, ChatResponse, Choice, UsageStatistics, GenerationResponse, GenerationData, GenerationUsage, and WrongPwdError. These data structures are used by the functions in the package, such as main, GenerateCompletion, GenerateCompletionWithPWD, GenerateImageStableDiffusion, TranscribeWhisper, and cleanText. The functions use environment variables, file paths, and URLs to configure their behavior.

## Unclear Places or Dead Code
None found

