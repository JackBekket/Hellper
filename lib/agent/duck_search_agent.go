package agent

import (
	"context"
	"encoding/json"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools/duckduckgo"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/langgraphgo/graph"
)

func MainDuckSearch() {

  model, err := openai.New(openai.WithModel("gpt-4o"))
  if err != nil {
    panic(err)
  }

  intialState := []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeSystem, "You are an agent that has access to a Duck Duck go search engine. Please provide the user with the information they are looking for by using the search tool provided."),
  }

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
  }

  agent := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
    response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))
    if err != nil {
      return state, err
    }

    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

    if len(response.Choices[0].ToolCalls) > 0 {
      for _, toolCall := range response.Choices[0].ToolCalls {
        msg.Parts = append(msg.Parts, toolCall)
      }
    }

    state = append(state, msg)
    return state, nil
  }

  search := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
    lastMsg := state[len(state)-1]

    for _, part := range lastMsg.Parts {
      toolCall, ok := part.(llms.ToolCall)

      if ok && toolCall.FunctionCall.Name == "search" {
        var args struct {
          Query string `json:"query"`
        }

        if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
          return state, err
        }

        search, err := duckduckgo.New(1, duckduckgo.DefaultUserAgent)
        if err != nil {
          log.Printf("search error: %v", err)
          return state, err
        }

        toolResponse, err := search.Call(ctx, args.Query)
        if err != nil {
          log.Printf("search error: %v", err)
          return state, err
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

  shouldSearch := func(ctx context.Context, state []llms.MessageContent) string {
    lastMsg := state[len(state)-1]
    for _, part := range lastMsg.Parts {
      toolCall, ok := part.(llms.ToolCall)

      if ok && toolCall.FunctionCall.Name == "search" {
        log.Printf("agent should use search")
        return "search"
      }
    }

    return graph.END
  }

  workflow := graph.NewMessageGraph()

  workflow.AddNode("agent", agent)
  workflow.AddNode("search", search)

  workflow.SetEntryPoint("agent")
  workflow.AddConditionalEdge("agent", shouldSearch)
  workflow.AddEdge("search", "agent")

  app, err := workflow.Compile()
  if err != nil {
    log.Printf("error: %v", err)
    return
  }

  intialState = append(
    intialState,
    llms.TextParts(llms.ChatMessageTypeHuman, "Who won the last FIFA World Cup?"),
  )

  response, err := app.Invoke(context.Background(), intialState)
  if err != nil {
    log.Printf("error: %v", err)
    return
  }

  lastMsg := response[len(response)-1]
  log.Printf("last msg: %v", lastMsg.Parts[0])
}
