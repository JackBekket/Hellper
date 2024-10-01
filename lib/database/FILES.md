# lib/database/newAiSessionDataBase.go  
Package: database  
  
Imports:  
- gogpt  
  
External data, input sources:  
- AiSessionMap: A map that stores AiSession objects, where the key is an int64 and the value is an AiSession object.  
  
Summary:  
The code defines a map called AiSessionMap, which stores AiSession objects. Each AiSession object has three fields: GptKey, GptClient, and GptModel. The GptKey field is a string, the GptClient field is a pointer to a gogpt.Client object, and the GptModel field is a string. The AiSessionMap is initialized as an empty map.  
  
  
  
# lib/database/newUserDataBase.go  
## Package: database  
  
### Imports:  
- github.com/tmc/langchaingo/chains  
- github.com/tmc/langchaingo/memory  
  
### External Data, Input Sources:  
- UsersMap: A map that stores user data, where the key is the telegram user ID and the value is a User struct.  
- UsageMap: A map that stores session usage data, where the key is the session ID and the value is a SessionUsage struct.  
  
### Code Summary:  
#### User Struct:  
- Represents a user in the database.  
- Contains fields for user ID, username, dialog status, admin status, AI session, and network.  
  
#### SessionUsage Struct:  
- Represents a session usage record.  
- Contains fields for session ID and usage data.  
  
#### AiSession Struct:  
- Represents an AI session.  
- Contains fields for GPT key, GPT model, AI type, dialog thread, base URL, and usage data.  
  
#### ChatSession Struct:  
- Represents a chat session.  
- Contains fields for conversation buffer and dialog thread.  
  
#### AddUser Function:  
- Adds a new user to the UsersMap.  
  
#### UpdateUserUsage Function:  
- Updates the usage data for a specific user in the UsersMap.  
  
#### UpdateSessionUsage Function:  
- Updates the usage data for a specific session in the UsageMap.  
  
#### GetSessionUsage Function:  
- Retrieves the usage data for a specific session from the UsageMap.  
  
