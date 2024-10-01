# main.go  
## Package: hellper/lib/bot  
  
### Imports:  
  
- context  
- log  
- os  
- strconv  
- github.com/JackBekket/hellper/lib/bot/command  
- github.com/JackBekket/hellper/lib/bot/dialog  
- github.com/JackBekket/hellper/lib/bot/env  
- github.com/JackBekket/hellper/lib/database  
- github.com/go-telegram-bot-api/telegram-bot-api/v5  
- github.com/joho/godotenv  
  
### External Data, Input Sources:  
  
- OPENAI_API_KEY (local key for localai)  
- PG_LINK (not used in the code)  
- TG_KEY (telegram bot token)  
- ADMIN_ID (admin user ID)  
- AI_ENDPOINT (local AI endpoint)  
  
### Code Summary:  
  
#### Initialization:  
  
1. Loads environment variables using godotenv.Load() and retrieves values for OPENAI_API_KEY, TG_KEY, ADMIN_ID, and AI_ENDPOINT.  
2. Creates a new Telegram bot instance using the retrieved TG_KEY.  
3. Initializes a map of admin data, including the admin ID and GPT key.  
  
#### Database and Commander:  
  
1. Initializes a database for storing user information.  
2. Creates a new command commander instance, which handles incoming commands and updates the database accordingly.  
  
#### Update Handling:  
  
1. Sets up a channel for handling incoming updates from the Telegram bot.  
2. Starts a goroutine to handle updates using the dialog.HandleUpdates function.  
3. Iterates through incoming updates and checks if the user is new. If so, adds the user to the database.  
  
#### Main Loop:  
  
1. Continuously listens for new updates from the Telegram bot.  
2. For each update, checks if the user is new and adds them to the database if necessary.  
3. Sends the update to the update handler goroutine for processing.  
  
