package database

import (
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/llms"
)

func (s *Service) GetHistory(
	userID, endpointId, chatId, threadId int64, model string,
) ([]llms.MessageContent, error) {

	messages := []llms.MessageContent{}
	rows, err := s.DBHandler.DB.Query(`
        SELECT m.message_data
        FROM chat_messages m
        JOIN chat_sessions cs ON m.chat_session = cs.id
		WHERE tg_user_id = $1 AND
			endpoint = $2 AND
			chat_id = $3 AND
			thread_id = $4 AND
			model = $5
        ORDER BY m.created_at
		`, userID, endpointId, chatId, threadId, model)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		contentBytes := []byte{}
		err := rows.Scan(&contentBytes)
		if err != nil {
			return messages, err
		}
		content := llms.MessageContent{}
		err = content.UnmarshalJSON(contentBytes)
		if err != nil {
			return messages, err
		}
		messages = append(messages, content)
	}

	return messages, err
}

func (s *Service) DropHistory(userID, providerID, chatId, threadId int64, model string) error {
	res, err := s.DBHandler.DB.Exec(`
        DELETE FROM chat_messages
        WHERE chat_session IN (
            SELECT id
            FROM chat_sessions
            WHERE tg_user_id = $1 AND
				endpoint = $2 AND
				chat_id = $3 AND
				thread_id = $4 AND
				model = $5
        )
		`, userID, providerID, chatId, threadId, model)
	if err != nil {
		log.Warn().Err(err).Msg("failed to execute delete query")
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	log.Info().Int64("rowsAffected", rowsAffected).Msg("Query executed successfully")
	return nil
}

func (s *Service) UpdateHistory(
	userID, endpointId, chatId, threadId int64,
	model string, content llms.MessageContent,
) error {
	contentBytes, err := content.MarshalJSON()
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`WITH
		ChatSessionCheck AS (
			SELECT id
			FROM chat_sessions
			WHERE tg_user_id = $1 AND
				endpoint = $2 AND
				chat_id = $3 AND
				thread_id = $4 AND
				model = $5
		),
		InsertChatSession AS (
			INSERT INTO chat_sessions
			(tg_user_id, endpoint, chat_id, thread_id, model)
			SELECT $1, $2, $3, $4, $5
			WHERE NOT EXISTS (SELECT 1 FROM ChatSessionCheck)
			RETURNING id
		),
		ChatSessionID AS (
			SELECT id FROM ChatSessionCheck
			UNION ALL
			SELECT id FROM InsertChatSession
		)
		INSERT INTO chat_messages (chat_session, message_data)
		SELECT (SELECT id FROM ChatSessionID), $6
		`, userID, endpointId, chatId, threadId, model, contentBytes)
	return err
}
