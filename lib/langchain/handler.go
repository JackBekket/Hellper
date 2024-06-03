package langchain

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/JackBekket/hellper/lib/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	//""
)


type ChainCallbackHandler struct{}



// HandleAgentAction implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleAgentAction(ctx context.Context, action schema.AgentAction) {
	//panic("unimplemented")
}

// HandleAgentFinish implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleAgentFinish(ctx context.Context, finish schema.AgentFinish) {
	//panic("unimplemented")
}

// HandleChainEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainEnd(ctx context.Context, outputs map[string]any) {
	//panic("unimplemented")
}

// HandleChainError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleChainStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleChainStart(ctx context.Context, inputs map[string]any) {
	//panic("unimplemented")
}

// HandleLLMError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleLLMGenerateContentStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent) {
	//panic("unimplemented")
}

// HandleLLMStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleLLMStart(ctx context.Context, prompts []string) {
	//panic("unimplemented")
}

// HandleRetrieverEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document) {
	//panic("unimplemented")
}

// HandleRetrieverStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleRetrieverStart(ctx context.Context, query string) {
	//panic("unimplemented")
}

// HandleStreamingFunc implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleStreamingFunc(ctx context.Context, chunk []byte) {
	//panic("unimplemented")
}

// HandleToolEnd implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolEnd(ctx context.Context, output string) {
	//panic("unimplemented")
}

// HandleToolError implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolError(ctx context.Context, err error) {
	//panic("unimplemented")
}

// HandleToolStart implements callbacks.Handler.
func (h *ChainCallbackHandler) HandleToolStart(ctx context.Context, input string) {
	//panic("unimplemented")
}

func (h *ChainCallbackHandler) HandleText(ctx context.Context, text string) {
	// Implement this method if needed
}

// Implement other methods...

func (h *ChainCallbackHandler) HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse) {

	LogResponseContentChoice(ctx,res)	
}

func LogResponseContentChoice(ctx context.Context,resp *llms.ContentResponse) {
	//choice *llms.ContentChoice
	choice := resp.Choices[0]
	log.Println("Content: ", choice.Content)
	log.Println("Stop Reason: ", choice.StopReason)

	log.Println("Context: ", ctx)
	
	//Get user from context
	user, ok := ctx.Value("user").(db.User)
	if !ok {
	  log.Println("No user in context")
	  //log.Println
	  return
	}
	//chatID := user.ID



	// GenerationInfo is a map that could contain complex/nested structures,
	// so we'll marshal it into a JSON string for a cleaner log message.
	// This step is optional and depends on your preference for log clarity.
	genInfo, err := json.Marshal(choice.GenerationInfo)
	if err != nil {
		log.Println("Error marshaling GenerationInfo: ", err)
		return
	}
	log.Println("Generation Info: ", string(genInfo))

	// If you have specific fields you expect in GenerationInfo, you can log them individually:
	// Example: log.Println("Some specific gen info: ", choice.GenerationInfo["someKey"])
	log.Println("Promt tokens: ", choice.GenerationInfo["PromptTokens"])
	log.Println("Completion tokens: ", choice.GenerationInfo["CompletionTokens"])
	log.Println("Total tokens: ", choice.GenerationInfo["TotalTokens"])

	// type any
	promt_tokens_str := choice.GenerationInfo["PromptTokens"]
	completion_tokens_str := choice.GenerationInfo["CompletionTokens"]
	total_tokens_str := choice.GenerationInfo["TotalTokens"]


	// type assertion (string)
	pt, ok := promt_tokens_str.(int)
	if !ok {
  	log.Println("Error: value is not a string")
  	return
	}
	ct, ok := completion_tokens_str.(int)
	tt, ok := total_tokens_str.(int)
	
		  // Update the user's usage information.
		  usage := map[string]int{
			"Total": tt,
			"Promt": pt,
			"Completion": ct,
		  }
		
		  // Save the user usage back to the database -- it's will not update user info, but stored it in separate structure to avoide race condition
		  db.UpdateSessionUsage(user.ID,usage)



	// Note: Since FuncCall is a pointer to a schema.FunctionCall, ensure you check for nil to avoid panics.
	if choice.FuncCall != nil {
		// Assuming FuncCall has fields you want to log, replace 'FieldName' with actual fields.
		log.Printf("Function Call: %+v\n", choice.FuncCall)
		// For specific field: log.Println("FuncCall field: ", choice.FuncCall.FieldName)
	} else {
		log.Println("No Function Call requested.")
	}
}