package openaibot

import gogpt "github.com/sashabaranov/go-openai"

/*
// GPT-3.5

	func createSimpleChatRequest(input string) gogpt.ChatCompletionRequest {
		return gogpt.ChatCompletionRequest{
			Model:     gogpt.GPT3Dot5Turbo,
			MaxTokens: 3000,
			Messages: []gogpt.ChatCompletionMessage{
				{
					Role:    gogpt.ChatMessageRoleUser,
					Content: input,
				},
			}}
	}
*/

func createComplexChatRequest(input string, modelName string) gogpt.ChatCompletionRequest {
	return gogpt.ChatCompletionRequest{
		Model:     modelName,
		MaxTokens: 3000,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    gogpt.ChatMessageRoleUser,
				Content: input,
			},
		}}
}

// for code generation
func createCodexRequest(input string) gogpt.CompletionRequest {
	return gogpt.CompletionRequest{
		Model:     gogpt.CodexCodeDavinci002,
		MaxTokens: 6000,
		Prompt:    input,
		Echo:      true,
	}
}

func createImageRequest(input string) gogpt.ImageRequest {
	return gogpt.ImageRequest{
		Prompt:         input,
		Size:           gogpt.CreateImageSize1024x1024,
		ResponseFormat: gogpt.CreateImageResponseFormatURL,
		N:              1,
	}
}

/*
// used for GPT-3
func CreateSimpleTextRequest(input string) (gogpt.CompletionRequest){
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Dot5Turbo,
		MaxTokens: 2048,
		Prompt:    input,
		Echo: true,
	}
	return req
}
*/

/*
// model should be gogpt.GPT3TextDavinci003 or gogpt.CodexCodeDavinci002
// WARN -- deprecated!
func CreateComplexRequest (input string, modelName string) (gogpt.CompletionRequest) {
	req := gogpt.CompletionRequest{
		Model: modelName,
		MaxTokens: 2048,
		Prompt: input,
		Echo: true,
	}
	return req
}
*/
