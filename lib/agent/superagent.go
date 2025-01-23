// This will be prototype to superagent (autonomouse agent, which work with memory and have similar functionality to langchain chains.Run method)
package agent

import "log"



func RunSingle(prompt string) {


 call := OneShotRun(prompt)
 log.Println(call)




}

