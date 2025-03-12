package database

const (
	AuthMethodLocalAI = 1
	AuthMethodOpenAI  = 2
)

type AISessionDB struct {
	Model string
	// AIProvider is a more appropriate name for the information stored in the Endpoint structure.
	// Read the comment for the Endpoint type
	AIProvider *Endpoint
}

// The table in the database was named incorrectly, as it does not store endpoints; it stores the base URLs of different services.
// It could have been named something like "entrypoints," "providers," or "services," but not "endpoints."
// Leaving the table name in the structure until the table name is changed
type Endpoint struct {
	ID         int64
	Name       string
	BaseURL    string
	AuthMethod int64
}

type Usage struct {
	CompletionTokens       int
	PromptTokens           int
	TotalTokens            int
	ReasoningTokens        int
	TimingTokenGeneration  float64
	TimingPromptProcessing float64
}
