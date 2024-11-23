package agent_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools/duckduckgo"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/embeddings"
	"github.com/JackBekket/langgraphgo/graph"
)

func Test_Run (t *testing.T) {


  model_name := "tiger-gemma-9b-v1-i1"
  _ = godotenv.Load()
          ai_url := os.Getenv("AI_ENDPOINT")          //TODO: should be global?
          api_token := os.Getenv("ADNIN_KEY")
          //db_link := os.Getenv("EMBEDDINGS_DB_URL")

  model, err := openai.New(
    openai.WithToken(api_token),
    //openai.WithBaseURL("http://localhost:8080"),
    openai.WithBaseURL(ai_url),
    openai.WithModel(model_name),
    openai.WithAPIVersion("v1"),
  )
  if err != nil {
    log.Fatal(err)
  }


  intialState := []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeSystem, "You are an agent that has access to a Document Search. Please provide the user with the information they are looking for by using the semantic_search tool provided."),
  }

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
            "searchQuery": map[string]any{
              "type":        "string",
              "description": "The search query",
            },
            "maxResults": map[string]any{
              "type":        "integer",
              "description": "Maximum number of results",
            },
            "store": map[string]any{
              "type":        "object",
              "description": "Vector store",
            },
            "options": map[string]any{
              "type":        "array",
              "description": "Optional parameters for the search",
            },
          },
        },
      },
    },
  }

  //TODO: REWORK


// AGENT NODE
  agent := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
    response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))
    if err != nil {
      return state, err
    }
    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

    if len(response.Choices[0].ToolCalls) > 0 {
      for _, toolCall := range response.Choices[0].ToolCalls {
        if toolCall.FunctionCall.Name == "semantic_search" {

          msg.Parts = append(msg.Parts, toolCall)

        }
      }
    }
    state = append(state, msg)
    return state, nil
  }


// TOOL FUNCTIONS
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

  // Custom semantic search function
  semanticSearch := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
      lastMsg := state[len(state)-1]

      for _, part := range lastMsg.Parts {
        toolCall, ok := part.(llms.ToolCall)
        if ok && toolCall.FunctionCall.Name == "semantic_search" {

          // TODO: Extract query and store parameters from the arguments
          // ... (logic to extract necessary values for SemanticSearch call)
          var args struct {
            Query string `json:"query"`
            //Store string `json:"store"`
            //Options []map[string]any `json:"options"`
            Name string `json:"name"`
          }
          if err := json.Unmarshal([]byte(toolCall.FunctionCall.Arguments), &args); err != nil {
            // Handle any errors in deserializing the arguments
            return state, err
          }
          // Extract query from the args structure
          searchQuery := args.Query

          //get env
          _ = godotenv.Load()
          ai_url := os.Getenv("AI_ENDPOINT")          //TODO: should be global?
          api_token := os.Getenv("ADNIN_KEY")
          db_link := os.Getenv("EMBEDDINGS_DB_URL")

          // Retrieve your vector store based on the store value in the args
          // You'll likely need to have a method for getting the vector store based
          // on the store string ("store" value in the args)
          store, err := embeddings.GetVectorStoreWithOptions(ai_url,api_token,db_link,args.Name) // Implement this method
          if err != nil {
            // Handle errors in retrieving the vector store
            return state, err
          }

          maxResults := 10 // Set your desired maxResults here
          //options := args.Options // Pass in any additional options as needed

          // Call your SemanticSearch function here
          searchResults, err := embeddings.SemanticSearch(
            searchQuery, 
            maxResults,
            store, // Pass in your vector store
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



//CONDITIONS

  /*
  //should use duck_duck_go search engine
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
  */

  // should use semantc search tool we defined earlier
  shouldSearchDocuments := func(ctx context.Context, state []llms.MessageContent) string {
    lastMsg := state[len(state)-1]
    for _, part := range lastMsg.Parts {
      toolCall, ok := part.(llms.ToolCall)

      if ok && toolCall.FunctionCall.Name == "semantic_search" {
        log.Printf("agent should use SemanticSearch (embeddings similarity search aka DocumentsSearch)")
        return "semantic_search"
      }
    }

    return graph.END
  }



  workflow := graph.NewMessageGraph()

  workflow.AddNode("agent", agent)
  workflow.AddNode("search", search)
  workflow.AddNode("semantic_search", semanticSearch)

  workflow.SetEntryPoint("agent")
  workflow.AddConditionalEdge("agent", shouldSearchDocuments)
  workflow.AddEdge("semantic_search", "agent")

  app, err := workflow.Compile()
  if err != nil {
    log.Printf("error: %v", err)
    return
  }

  intialState = append(
    intialState,  //TODO: check if we can somehow set collection name in initial state
    llms.TextParts(llms.ChatMessageTypeHuman, "Collection Name: Hellper Query: How does embeddings package works?"),
  )

  response, err := app.Invoke(context.Background(), intialState)
  if err != nil {
    log.Printf("error: %v", err)
    return
  }

  lastMsg := response[len(response)-1]
  log.Printf("last msg: %v", lastMsg.Parts[0])
}
