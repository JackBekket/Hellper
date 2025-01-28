package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"

	//"github.com/tmc/langgraphgo/graph"
	"github.com/JackBekket/hellper/lib/embeddings"
	"github.com/JackBekket/langgraphgo/graph"
)

/** My current vision of this mechanism is a graph. So each agent can be represented as graph. Each node is usually single action in <turn_of_dialog>. Graphs is connected with themselves through edges, which represent
  relations whithin graphs. Each graph can be conditional or direct. If we need to reorder graph we can simply alter entry_point instead of rewriting code of dialog itseelf every time.
  Each graph can also be represented graphically.


    This is OneShot agent example
    It does not have memory by itself, but memory (history of previouse messages) can be passed as optional parameter

    So let's say at high level this code package is a graph (or supergraph), so the main logic of this package is a workflow graph.
    This graph ensures that messages stack (dialog) is processed as intended.
    Graph have multiple nodes, which starts with entry_points, and those nodes are connected via edges.
    There may be direct edges or a conditional edges

    Graph must end in sometime to be able to compile (doesn't really work you have to make shure there is no loophole)

    Nodes are basically could be tool nodes, agents and conditions

    Agent is basically main thinking or decision-making algorythm, it is responsible to call tools, process user input, process tool responses, etc.

    Conditonal node is basically algorythm that check if Agent is calling some tools and if it is -- call this tools. Note, that agent is *calling* the tool, but conditional node *handle this tool call*

    If your task is to create a documentation for this package -- start with describing workflow, what are the nodes in general, how does they connect with each other in general, how workflow works in general.
    Then explain/describe how agent node works, what tools does it have. Then describe how each tool works.
    Then make a general summary for this package


*/

// global var
var Model openai.LLM

var Tools []llms.Tool



func OneShotRun(prompt string, model openai.LLM,history_state ...llms.MessageContent) string{

  
  // Operation with message STATE stack

  agentState := []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeSystem, "You are helpful agent that has access to a semanticSearch tool. Use this tool if user ask to retrive some information from database/collection to provide user with information he/she looking for."),
  }
  intialState := []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeSystem, "Below a current conversation between user and helpful AI assistant. You (assistant) should help user in any task he/she ask you to do."),
  }

  if len(history_state) > 0 {                   // if there are previouse message state then we first load it into message state
    // Access the first element of the slice
    history := history_state
    // ... use the history variable as needed
    for _, message := range history {
      intialState = append(intialState, message)  // load history as initial state
    }
    intialState = append(
      intialState,   
      agentState...,        // append agent system prompt     //something wrong here
    )
    intialState = append(
      intialState,  
      llms.TextParts(llms.ChatMessageTypeHuman, prompt),  //append user input (!)
    )
  } else {
    intialState = agentState    //history is empty -- load agentState as initial_state and append user prompt
    intialState = append(
      intialState,  
      llms.TextParts(llms.ChatMessageTypeHuman, prompt),
    )
    
    
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
            "query": map[string]any{
              "type":        "string",
              "description": "The search query",
            },
            "collection": map[string]any{                     //TODO: there should NOT exist arguments which called NAME cause it cause COLLISION with actual function name.    .....more like confusion then collision so there are no error
              "type":        "string",
              "description": "name of collection store in which we perform the search",
            },
          }, 
        },
      },
    },
  }

  Tools = tools
  Model = model


// AGENT NODE
/** We are telling agent, that it should response withTools, giving it function signatures defined earlier. 
    if agent get response from conditional edge like 'yes, use x function with this signatures and this json object as input parameters -- it will match with predefined pointer to semanticSearch function and it will make a toolCall
    then it will append toolCall to message state.
    Note, that agent can call few toolCalls and all of them can be append here. toolCalls may be done parallel (I guess) 
  agent := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
   

    consideration_query := []llms.MessageContent{
      llms.TextParts(llms.ChatMessageTypeSystem, "Your task is to determine whether or not to call semanticSearch function based on human input. You should ONLY return 'true' or 'false'."),
    }

    lastMsg := state[len(state)-1]
        if lastMsg.Role == "tool" {   // If we catch response from tool then we use this response
          state = append(state, lastMsg)
          response, err := model.GenerateContent(ctx, state)
          if err != nil {
            return state, err
          }
          msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)
          state = append(state, msg)
          return state,nil


  }   else {   // If it is not tool response 

    if lastMsg.Role == "human" {    //                                            any user request
      //state
      consideration_stack := append(consideration_query, state...)
      check, err := model.GenerateContent(ctx, consideration_stack)               // one punch which determine wheter or not call tools. this is hardcode and probably should be separate part of the graph.
      if err != nil {
        return state, err
      }
      check_txt := fmt.Sprintf(check.Choices[0].Content) 
      log.Println("check result: ", check_txt)

      if check_txt == "true" {                                                      // tool call required by one-shot agent
        response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))   // AI call tool function.. in this step it just put call in messages stack
        if err != nil {
          return state, err
        }
    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

    if len(response.Choices[0].ToolCalls) > 0 {
      for _, toolCall := range response.Choices[0].ToolCalls {
        if toolCall.FunctionCall.Name == "semanticSearch" {       // AI catch that there is a function call in messages, so *now* it actually calls the function.

          msg.Parts = append(msg.Parts, toolCall) // Add result to messages stack

        }
      }
      state = append(state, msg)  
      return state, nil
    }
     } else {                                                     // proceed without tools
      response, err := model.GenerateContent(ctx, state)
      if err != nil {
        return state, err
      }
      msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)
      state = append(state, msg)
      return state,nil
     }
        // end if_human
    }
 

  


    // if we get to this point, then our input is not tool result and not a human... I have no idea how it can reach this point

    response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))   // AI call tool function.. in this step it just put call in messages stack
    if err != nil {
      return state, err
    }
    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

    if len(response.Choices[0].ToolCalls) > 0 {
      for _, toolCall := range response.Choices[0].ToolCalls {
        if toolCall.FunctionCall.Name == "semanticSearch" {       // AI catch that there is a function call in messages, so *now* it actually calls the function.

          msg.Parts = append(msg.Parts, toolCall) // Add result to messages stack

        }
      }
    }
    state = append(state, msg)  
    return state, nil
  }
}
*/

  
/*
// TOOL FUNCTIONS
  // Custom semantic search function for working with vector-store information
  semanticSearch := func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
    lastMsg := state[len(state)-1]

    for _, part := range lastMsg.Parts {
      toolCall, ok := part.(llms.ToolCall)
      if ok && toolCall.FunctionCall.Name == "semanticSearch" {

        // TODO: Extract query and store parameters from the arguments
        // ... (logic to extract necessary values for SemanticSearch call)
        var args struct {
          Query string `json:"query"`
          //Store string `json:"store"`
          //Options []map[string]any `json:"options"`
          Collection string `json:"collection"`   //TODO: ALWAYS CHECK THIS JSON REFERENCE WHEN ALTERING VARS
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
        ai_url := os.Getenv("AI_ENDPOINT")          //TODO: should be global?   .. there are global, there might be resetting.
        api_token := os.Getenv("ADNIN_KEY")
        db_link := os.Getenv("EMBEDDINGS_DB_URL")

        log.Println("Collection Name: ", args.Collection)
        log.Println("db_link: ", db_link)


        // Retrieve your vector store based on the store value in the args
        // You'll likely need to have a method for getting the vector store based
        // on the store string ("store" value in the args)
        store, err := embeddings.GetVectorStoreWithOptions(ai_url,api_token,db_link,args.Collection) // TODO: changed argument 'Name' to 'CollectionName' or something like that
        if err != nil {
          // Handle errors in retrieving the vector store
    log.Println("error getting store")
          return state, err
        }

        log.Println("store:", store)

        maxResults := 1 // Set your desired maxResults here
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
*/


/*
  //CONDITIONS funcs
  // condition function, which defines whether or not to use semanticSearch tool. we have access to semanticSearch itself in main thread through a pointer to this function. So if llm says 'yes, use this function with x signatures` -- it will match to a pointer and x function will be called.`
  shouldSearchDocuments := func(ctx context.Context, state []llms.MessageContent) string {
  
    // this function (I suppose) can be reworked to work with a *set* of a functions, not just one func.

    lastMsg := state[len(state)-1]
    for _, part := range lastMsg.Parts {
      toolCall, ok := part.(llms.ToolCall)

      if ok && toolCall.FunctionCall.Name == "semanticSearch"  {
        log.Printf("agent should use SemanticSearch (embeddings similarity search aka DocumentsSearch)")
        return "semanticSearch"
      }
    }

    return graph.END  // never reach this point, should be removed?
  }
*/



// MAIN WORKFLOW
  workflow := graph.NewMessageGraph()

  workflow.AddNode("agent", agent)
  //workflow.AddNode("search", search)
  workflow.AddNode("semanticSearch", semanticSearch)

  workflow.SetEntryPoint("agent")
  workflow.AddConditionalEdge("agent", shouldSearchDocuments)
  workflow.AddEdge("semanticSearch", "agent")

  app, err := workflow.Compile()
  if err != nil {
    log.Printf("error: %v", err)
    return fmt.Sprintf("error :%v", err)
  }

  response, err := app.Invoke(context.Background(), intialState)
  if err != nil {
    log.Printf("error: %v", err)
    return fmt.Sprintf("error :%v", err)
  }

  lastMsg := response[len(response)-1]
  log.Printf("last msg: %v", lastMsg.Parts[0]) 
  result := lastMsg.Parts[0]
  result_str := fmt.Sprintf("%v", result)
  return result_str
}


// AGENT NODE
/** We are telling agent, that it should response withTools, giving it function signatures defined earlier. 
    if agent get response from conditional edge like 'yes, use x function with this signatures and this json object as input parameters -- it will match with predefined pointer to semanticSearch function and it will make a toolCall
    then it will append toolCall to message state.
    Note, that agent can call few toolCalls and all of them can be append here. toolCalls may be done parallel (I guess) */
func agent(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
  
  
  model := Model  // global... should be .env or getting from user context I guess.
  tools := Tools
  
  consideration_query := []llms.MessageContent{
    llms.TextParts(llms.ChatMessageTypeSystem, "Your task is to determine whether or not to call semanticSearch function based on human input. You should ONLY return 'true' or 'false'."),
  }

  lastMsg := state[len(state)-1]
      if lastMsg.Role == "tool" {   // If we catch response from tool then we use this response
        state = append(state, lastMsg)
        response, err := model.GenerateContent(ctx, state)
        if err != nil {
          return state, err
        }
        msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)
        state = append(state, msg)
        return state,nil


}   else {   // If it is not tool response 

  if lastMsg.Role == "human" {    //                                            any user request
    //state
    consideration_stack := append(consideration_query, state...)
    check, err := model.GenerateContent(ctx, consideration_stack)               // one punch which determine wheter or not call tools. this is hardcode and probably should be separate part of the graph.
    if err != nil {
      return state, err
    }
    check_txt := fmt.Sprintf(check.Choices[0].Content) 
    log.Println("check result: ", check_txt)

    if check_txt == "true" {                                                      // tool call required by one-shot agent
      response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))   // AI call tool function.. in this step it just put call in messages stack
      if err != nil {
        return state, err
      }
  msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

  if len(response.Choices[0].ToolCalls) > 0 {
    for _, toolCall := range response.Choices[0].ToolCalls {
      if toolCall.FunctionCall.Name == "semanticSearch" {       // AI catch that there is a function call in messages, so *now* it actually calls the function.

        msg.Parts = append(msg.Parts, toolCall) // Add result to messages stack

      }
    }
    state = append(state, msg)  
    return state, nil
  }
   } else {                                                     // proceed without tools
    response, err := model.GenerateContent(ctx, state)
    if err != nil {
      return state, err
    }
    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)
    state = append(state, msg)
    return state,nil
   }
  } // end if human
  response, err := model.GenerateContent(ctx, state, llms.WithTools(tools))   // AI call tool function.. in this step it just put call in messages stack
    if err != nil {
      return state, err
    }
    msg := llms.TextParts(llms.ChatMessageTypeAI, response.Choices[0].Content)

    if len(response.Choices[0].ToolCalls) > 0 {
      for _, toolCall := range response.Choices[0].ToolCalls {
        if toolCall.FunctionCall.Name == "semanticSearch" {       // AI catch that there is a function call in messages, so *now* it actually calls the function.

          msg.Parts = append(msg.Parts, toolCall) // Add result to messages stack

        }
      }
    }
    state = append(state, msg)  
    return state, nil
  } // end if not tool response
  } 




func shouldSearchDocuments(ctx context.Context, state []llms.MessageContent)  string{
        // this function (I suppose) can be reworked to work with a *set* of a functions, not just one func.

        lastMsg := state[len(state)-1]
        for _, part := range lastMsg.Parts {
          toolCall, ok := part.(llms.ToolCall)
    
          if ok && toolCall.FunctionCall.Name == "semanticSearch"  {
            log.Printf("agent should use SemanticSearch (embeddings similarity search aka DocumentsSearch)")
            return "semanticSearch"
          }
        }
    
        return graph.END  // never reach this point, should be removed?
}

func semanticSearch(ctx context.Context, state []llms.MessageContent)  ([]llms.MessageContent, error){
  lastMsg := state[len(state)-1]

  for _, part := range lastMsg.Parts {
    toolCall, ok := part.(llms.ToolCall)
    if ok && toolCall.FunctionCall.Name == "semanticSearch" {

      // TODO: Extract query and store parameters from the arguments
      // ... (logic to extract necessary values for SemanticSearch call)
      var args struct {
        Query string `json:"query"`
        //Store string `json:"store"`
        //Options []map[string]any `json:"options"`
        Collection string `json:"collection"`   //TODO: ALWAYS CHECK THIS JSON REFERENCE WHEN ALTERING VARS
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
      ai_url := os.Getenv("AI_ENDPOINT")          //TODO: should be global?   .. there are global, there might be resetting.
      api_token := os.Getenv("ADNIN_KEY")
      db_link := os.Getenv("EMBEDDINGS_DB_URL")

      log.Println("Collection Name: ", args.Collection)
      log.Println("db_link: ", db_link)


      // Retrieve your vector store based on the store value in the args
      // You'll likely need to have a method for getting the vector store based
      // on the store string ("store" value in the args)
      store, err := embeddings.GetVectorStoreWithOptions(ai_url,api_token,db_link,args.Collection) // TODO: changed argument 'Name' to 'CollectionName' or something like that
      if err != nil {
        // Handle errors in retrieving the vector store
  log.Println("error getting store")
        return state, err
      }

      log.Println("store:", store)

      maxResults := 1 // Set your desired maxResults here
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
  
