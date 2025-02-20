package tools

import (
	"github.com/tmc/langchaingo/llms"
	// ... other necessary imports
)

type SemanticSearchTool struct {
    // ... tool-specific implementation
}

func  GetTools() ([]llms.Tool, error) {
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
		return tools,nil
}

/*
func (s *SemanticSearchTool) Execute(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
    // ... tool execution logic
}
*/