package config

// This package is designed for centralized storage of configurations required for service operation

// Configuration for working with various AI APIs
type AIConfig struct {
	ModelsListEndpoint       string
	ImageGenerationModel     string
	ImageGenerationEndpoint  string
	ImageRecognitionModel    string
	ImageRecognitionEndpoint string
	VoiceRecognitionModel    string
	VoiceRecognitionEndpoint string
}
