package tools

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/JackBekket/hellper/lib/embeddings"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	// ... other necessary imports
)

type SemanticSearchTool struct {
	// ... tool-specific implementation
}

type Tool interface {
	Execute(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error)
}

func GetTools() ([]llms.Tool, error) {
	// ... tool init logic
	// toolS definition interfaces
	tools := []llms.Tool{
		{
			Type: "function",
			Function: &llms.FunctionDefinition{
				Name:        "search",
				Description: "Preforms Duck Duck Go web search",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"query": map[string]any{
							"type":        "string",
							"description": "The search query",
						},
					},
				},
			},
		},
		{
			Type: "function",
			Function: &llms.FunctionDefinition{
				Name:        "semanticSearch",
				Description: "Performs semantic search using a vector store",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"query": map[string]any{
							"type":        "string",
							"description": "The search query",
						},
						"collection": map[string]any{ //TODO: there should NOT exist arguments which called NAME cause it cause COLLISION with actual function name.    .....more like confusion then collision so there are no error
							"type":        "string",
							"description": "name of collection store in which we perform the search",
						},
					},
				},
			},
		},
	}
	return tools, nil
}

// Unsure if needed here
// REFACTOR
func (s *SemanticSearchTool) Execute(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
	lastMsg := state[len(state)-1]

	for _, part := range lastMsg.Parts {
		toolCall, ok := part.(llms.ToolCall)
		if ok && toolCall.FunctionCall.Name == "semanticSearch" {

			// TODO: Extract query and store parameters from the arguments
			// (logic to extract necessary values for SemanticSearch call)
			var args struct {
				Query string `json:"query"`
				//Store string `json:"store"`
				//Options []map[string]any `json:"options"`
				Collection string `json:"collection"` //TODO: ALWAYS CHECK THIS JSON REFERENCE WHEN ALTERING VARS
			}
			if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
				// Handle any errors in deserializing the arguments
				log.Println("error unmurshal json")
				return state, err
			}
			// Extract query from the args structure
			searchQuery := args.Query

			//get env
			_ = godotenv.Load()
			endpoint := os.Getenv("AI_ENDPOINT") // there are global, there might be resetting.
			localAIToken := os.Getenv("OPENAI_API_KEY")
			dbLink := os.Getenv("DB_LINK")

			log.Println("Collection Name: ", args.Collection)
			log.Println("dbLink: ", dbLink)

			// Retrieve your vector store based on the store value in the args
			store, err := embeddings.GetVectorStoreWithOptions(endpoint, localAIToken, dbLink, args.Collection)
			if err != nil {
				log.Println("error getting store")
				return state, err
			}

			log.Println("store:", store) // actually return empty store in case of error (!)

			maxResults := 1 // Set your desired maxResults here
			//options := args.Options // Pass in any additional options as needed

			// Call *real* SemanticSearch function
			searchResults, err := embeddings.SemanticSearch(
				searchQuery,
				maxResults,
				store,
				// options, // Pass in any additional options you need
			)

			if err != nil {
				log.Printf("semantic search error: %v", err)
				return state, err
			}

			// Format and return search results
			// ... (process and format search results from SemanticSearch)
			//toolResponse := []string{} // Initialize an empty slice to store extracted text
			toolResponse := ""
			for _, result := range searchResults {
				//toolResponse = append(toolResponse, result.PageContent)
				toolResponse += result.PageContent + "\n"

			}

			msg := llms.MessageContent{
				Role: llms.ChatMessageTypeTool,
				Parts: []llms.ContentPart{
					llms.ToolCallResponse{
						ToolCallID: toolCall.ID,
						Name:       toolCall.FunctionCall.Name,
						Content:    toolResponse,
					},
				},
			}
			state = append(state, msg)
		}
	}
	return state, nil
}
