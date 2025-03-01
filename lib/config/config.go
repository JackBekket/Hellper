package config

// This package is designed for centralized storage of configurations required for service operation

// Configuration for working with various AI APIs
type AIConfig struct {
	AI_endpoint            string
	OpenAI_APIKey          string
	BaseURL                string
	ImageGenerationModel   string
	ImageGenerationSuffix  string
	ImageRecognitionModel  string
	ImageRecognitionSuffix string
	VoiceRecognitionModel  string
	VoiceRecognitionSuffix string
}
