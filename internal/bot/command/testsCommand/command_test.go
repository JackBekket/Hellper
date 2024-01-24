package command_test

import (
	"context"
	"testing"

	"github.com/JackBekket/telegram-gpt/internal/bot/command"
	"github.com/JackBekket/telegram-gpt/internal/bot/env"
	"github.com/JackBekket/telegram-gpt/internal/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
)

var (
	bot     = &tgbotapi.BotAPI{}
	usersDb = make(map[int64]database.User)
	ctx     = context.Background()
)

func TestCommander_AddAdminToMap_and_AddUserToMap(t *testing.T) {
	commander := command.NewCommander(bot, usersDb, ctx)

	t.Run("Correct addition of admin data to the database", func(t *testing.T) {
		// Вызываем функцию AddAdminToMap
		inputMessage := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:        123,
				UserName:  "kokos",
				FirstName: "Kowka",
				LastName:  "Kokosovaya",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}

		adminKey := "admin123"
		commander.AddAdminToMap(adminKey, inputMessage)

		// The admin has been added to the database with the expected values.
		expectedAdmin := database.User{
			ID:           123,
			Username:     "kokos",
			DialogStatus: 2,
			Admin:        true,
			AiSession: database.AiSession{
				GptKey: adminKey,
			},
		}
		if _, ok := usersDb[123]; !ok {
			t.Error("Admin not added to database")
		} else if usersDb[123] != expectedAdmin {
			t.Errorf("Incorrect admin data in database. Expected: %+v, Got: %+v", expectedAdmin, usersDb[123])
		}
	})

	t.Run("Correct addition of user data to the database", func(t *testing.T) {
		inputMessage := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:       456,
				UserName: "biba",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}
		commander.AddNewUserToMap(inputMessage)

		//  The user has been added to the database with the expected values.
		expectedUser := database.User{
			ID:           456,
			Username:     "biba",
			DialogStatus: 0,
			Admin:        false,
		}
		if _, ok := usersDb[456]; !ok {
			t.Error("User not added to database")
		} else if usersDb[456] != expectedUser {
			t.Errorf("Incorrect user data in database. Expected: %+v, Got: %+v", expectedUser, usersDb[123])
		}
	})

}

func TestCommander_CheckAdmin(t *testing.T) {
	commander := command.NewCommander(bot, usersDb, ctx)
	t.Run("Checks correct assigment to the database", func(t *testing.T) {
		//Emulation of paged data from .env file
		adminData := map[string]env.AdminData{
			"ADMIN_ID": {
				ID:     1111,
				GPTKey: "admin_key",
			},
			"MINTY_ID": {
				ID:     2222,
				GPTKey: "minty_key",
			},
			"OK_ID": {
				ID:     3333,
				GPTKey: "ok_key",
			},
			"MURS_ID": {
				ID:     4444,
				GPTKey: "murs_key",
			},
		}
		// Emulate an incoming message
		inputMsgWithAdminID := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:       1111,
				UserName: "Admin",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}

		inputMsgWithOkID := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:       3333,
				UserName: "ok_Admin",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}

		inputMsgWithUserID := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:        5555,
				UserName:  "randomUser",
				FirstName: "Random",
				LastName:  "User",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}
		commander.CheckAdmin(adminData, inputMsgWithAdminID)
		commander.CheckAdmin(adminData, inputMsgWithOkID)
		// check regular user
		commander.CheckAdmin(adminData, inputMsgWithUserID)

		expectedAdmins := map[int64]database.User{
			1111: {
				ID:           1111,
				Username:     "Admin",
				DialogStatus: 2,
				Admin:        true,
				AiSession: database.AiSession{
					GptKey: "admin_key",
				},
			},
			3333: {
				ID:           3333,
				Username:     "ok_Admin",
				DialogStatus: 2,
				Admin:        true,
				AiSession: database.AiSession{
					GptKey: "ok_key",
				},
			},
		}
		for id, expectedAdmin := range expectedAdmins {
			if _, ok := usersDb[id]; !ok {
				t.Errorf("Admin with ID %d not added to database", id)
			} else if usersDb[id] != expectedAdmin {
				t.Errorf("Incorrect admin data in database. Expected: %+v, Got: %+v", expectedAdmin, usersDb[id])
			}
		}

	})

	t.Run("Correct addition of user data to the database with an empty key", func(t *testing.T) {
		inputMsgWithAdminID := &tgbotapi.Message{
			Text: "/start",
			From: &tgbotapi.User{
				ID:       1111,
				UserName: "AdminWithEmptyKey",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}
		adminData := map[string]env.AdminData{
			"OK_ID": {
				ID:     1111,
				GPTKey: "", // Пустой ключ
			},
			"ADMIN_ID": {
				ID:     2222,
				GPTKey: "",
			},
			"MINTY_ID": {
				ID:     3333,
				GPTKey: "minty_key",
			},
		}
		commander.CheckAdmin(adminData, inputMsgWithAdminID)
		// The method must add an admin with an empty key as a regular user
		result := map[int64]database.User{
			1111: {
				ID:           1111,
				Username:     "AdminWithEmptyKey",
				DialogStatus: 0,
				Admin:        false,
			},
		}

		for id, user := range result {
			if _, ok := usersDb[id]; !ok {
				t.Errorf("User with ID %d not added to database", id)
			} else if usersDb[id] != user {
				t.Errorf("Incorrect user data in database. Expected: %+v, Got: %+v", user, usersDb[id])
			}
		}

	})

}

func TestCommander_Cases(t *testing.T) {
	commander := command.NewCommander(bot, usersDb, ctx)

	inputMsg := &tgbotapi.Message{
		Text: "/start",
		From: &tgbotapi.User{
			ID:       12345,
			UserName: "randomUser",
		},
		Chat: &tgbotapi.Chat{
			UserName: "testchat",
		},
	}

	commander.AddNewUserToMap(inputMsg)

	t.Run("Test InputYourAPIKey, updates DB", func(t *testing.T) {
		// Обновление DialogStatus
		commander.InputYourAPIKey(inputMsg)
		result := map[int64]database.User{
			12345: {
				ID:           12345,
				Username:     "randomUser",
				DialogStatus: 1,
				Admin:        false,
			},
		}
		for id, user := range result {
			if _, ok := usersDb[id]; !ok {
				t.Errorf("User with ID %d not added to database", id)
			} else if usersDb[id] != user {
				t.Errorf("Incorrect user data in database. Expected: %+v, Got: %+v", user, usersDb[id])
			}
		}
	})

	t.Run("Test ChooseModel, updates DB", func(t *testing.T) {
		inputMsg := &tgbotapi.Message{
			// tracks a message with a GPT key from the user
			Text: "key",
			From: &tgbotapi.User{
				ID:       12345,
				UserName: "randomUser",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}
		// updates DialogStatus and GPTkey
		commander.ChooseModel(inputMsg)
		result := map[int64]database.User{
			12345: {
				ID:           12345,
				Username:     "randomUser",
				DialogStatus: 2,
				Admin:        false,
				AiSession: database.AiSession{
					GptKey: "key",
				},
			},
		}
		for id, user := range result {
			if _, ok := usersDb[id]; !ok {
				t.Errorf("User with ID %d not added to database", id)
			} else if usersDb[id] != user {
				t.Errorf("Incorrect user data in database. Expected: %+v, Got: %+v", user, usersDb[id])
			}
		}
	})
	t.Run("Test ModelGPTDOT5, updates DB", func(t *testing.T) {
		inputMsg := &tgbotapi.Message{
			Text: "GPT-3.5",
			From: &tgbotapi.User{
				ID:       4567,
				UserName: "randomUser",
			},
			Chat: &tgbotapi.Chat{
				UserName: "testchat",
			},
		}
		// updates DialogStatus and GPTkey
		commander.ModelGPT3DOT5(inputMsg)
		result := map[int64]database.User{
			4567: {
				DialogStatus: 3,
				Admin:        false,
				AiSession: database.AiSession{
					GptModel: openai.GPT3Dot5Turbo,
				},
			},
		}
		for id, user := range result {
			if _, ok := usersDb[id]; !ok {
				t.Errorf("User with ID %d not added to database", id)
			} else if usersDb[id] != user {
				t.Errorf("Incorrect user data in database. Expected: %+v, Got: %+v", user, usersDb[id])
			}
		}
	})

}
