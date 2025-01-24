// This will be prototype to superagent (autonomouse agent, which work with memory and have similar functionality to langchain chains.Run method)
package agent

import (
	"log"

	"github.com/tmc/langchaingo/llms"
)

// This function fire One-Shot agent without history context
func OnePunch(prompt string) {

 call := OneShotRun(prompt)
 log.Println(call)
}


// this function recive previouse history message state and append new user prompt, than run agent
func RunThread(prompt string, history ...llms.MessageContent) {
	call := OneShotRun("Collection Name: 'Hellper' Query: How does embeddings package works? Also do you remember what is my name?", history...)
	log.Println(call)
}