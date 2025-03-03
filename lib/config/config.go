package config

// This package is designed for centralized storage of configurations required for service operation

// Configuration for working with various AI APIs
type AIConfig struct {
	AIEndpoint             string
	BaseURL                string
	ModelsListSuffix       string
	ImageGenerationModel   string
	ImageGenerationSuffix  string
	ImageRecognitionModel  string
	ImageRecognitionSuffix string
	VoiceRecognitionModel  string
	VoiceRecognitionSuffix string
}
