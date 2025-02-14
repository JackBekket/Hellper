package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
)

type AISession struct {
	Model    *string
	Endpoint *Endpoint
}

type Endpoint struct {
	ID         int64
	Name       string
	URL        string
	AuthMethod int64
}

type Usage struct {
	CompletionTokens       int
	PromptTokens           int
	TotalTokens            int
	ReasoningTokens        int
	TimingTokenGeneration  float64
	TimingPromptProcessing float64
}

func (s *Service) CreateTables() error {
	_, err := s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS auth_methods (
			id SERIAL PRIMARY KEY
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS endpoints (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			url TEXT NOT NULL,
			auth_method INT REFERENCES auth_methods(id)
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS auth (
			id SERIAL PRIMARY KEY,
			tg_user_id INT NOT NULL,
			auth_method INT REFERENCES auth_methods(id),
			token TEXT NOT NULL
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS ai_sessions (
			tg_user_id INT PRIMARY KEY,
			model TEXT,
			endpoint INT REFERENCES endpoints(id)	
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS chat_sessions (
			id SERIAL PRIMARY KEY,
			tg_user_id INT NOT NULL,
			model TEXT NOT NULL,
			endpoint INT NOT NULL REFERENCES endpoints(id),
			chat_id INT NOT NULL,
			thread_id INT NOT NULL
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS last_usage (
			id BIGSERIAL PRIMARY KEY,
			chat_session INT REFERENCES chat_sessions(id),
			completion_tokens INT,
			prompt_tokens INT,
			total_tokens INT,
			reasoning_tokens INT,
			timing_token_generation REAL,
			timing_token_processing REAL,
			UNIQUE (chat_session)
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS usage (
			id BIGSERIAL PRIMARY KEY,
			chat_session INT REFERENCES chat_sessions(id),
			tg_user_id INT NOT NULL,
			completion_tokens INT,
			prompt_tokens INT,
			total_tokens INT,
			reasoning_tokens INT,
			timing_token_generation REAL,
			timing_token_processing REAL,
			UNIQUE (chat_session)
		);

		CREATE UNIQUE INDEX IF NOT EXISTS unique_user_usage_null_session ON usage (tg_user_id) WHERE chat_session IS NULL;
		`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS chat_messages (
			id BIGSERIAL PRIMARY KEY,
			chat_session INT NOT NULL REFERENCES chat_sessions(id),
			message_data BYTEA NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX IF NOT EXISTS idx_chat_session_timestamp ON chat_messages (chat_session, created_at);
		`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE OR REPLACE FUNCTION upsert_chat_session_and_usage(
		    p_user_id BIGINT,
		    p_endpoint_id INT,
		    p_chat_id INT,
		    p_thread_id INT,
		    p_model TEXT,
		    p_usage_completion_tokens INT,
		    p_usage_prompt_tokens INT,
		    p_usage_total_tokens INT,
		    p_usage_reasoning_tokens INT,
		    p_usage_timing_token_generation REAL,
		    p_usage_timing_token_processing REAL
		)
		RETURNS VOID AS $$
		DECLARE
		    v_chat_session_id INT;
		BEGIN
		    SELECT id INTO v_chat_session_id
		    FROM chat_sessions
		    WHERE
		    	tg_user_id = p_user_id AND
		    	endpoint = p_endpoint_id AND
		    	chat_id = p_chat_id AND
		    	thread_id = p_thread_id;
		
		    IF NOT FOUND THEN
		        INSERT INTO chat_sessions
		        	(tg_user_id, model, endpoint, chat_id, thread_id)
		        VALUES
		        	(p_user_id, p_model, p_endpoint_id, p_chat_id, p_thread_id)
		        RETURNING id INTO v_chat_session_id;
		    END IF;
		
		    INSERT INTO usage
		    	(chat_session, tg_user_id, completion_tokens, prompt_tokens,
		    	total_tokens, reasoning_tokens,
		    	timing_token_generation, timing_token_processing)
		    VALUES
		    	(v_chat_session_id, p_user_id, p_usage_completion_tokens, p_usage_prompt_tokens,
		    	p_usage_total_tokens, p_usage_reasoning_tokens,
		    	p_usage_timing_token_generation, p_usage_timing_token_processing)
		    ON CONFLICT (chat_session) DO UPDATE SET
		        completion_tokens = usage.completion_tokens + EXCLUDED.completion_tokens,
		        prompt_tokens = usage.prompt_tokens + EXCLUDED.prompt_tokens,
		        total_tokens = usage.total_tokens + EXCLUDED.total_tokens,
		        reasoning_tokens = usage.reasoning_tokens + EXCLUDED.reasoning_tokens,
		        timing_token_generation = usage.timing_token_generation + EXCLUDED.timing_token_generation,
		        timing_token_processing = usage.timing_token_processing + EXCLUDED.timing_token_processing;

		    INSERT INTO last_usage
		    	(chat_session, completion_tokens, prompt_tokens,
		    	total_tokens, reasoning_tokens,
		    	timing_token_generation, timing_token_processing)
		    VALUES
		    	(v_chat_session_id, p_usage_completion_tokens, p_usage_prompt_tokens,
		    	p_usage_total_tokens, p_usage_reasoning_tokens,
		    	p_usage_timing_token_generation, p_usage_timing_token_processing)
		    ON CONFLICT (chat_session) DO UPDATE SET
		        completion_tokens = EXCLUDED.completion_tokens,
		        prompt_tokens = EXCLUDED.prompt_tokens,
		        total_tokens = EXCLUDED.total_tokens,
		        reasoning_tokens = EXCLUDED.reasoning_tokens,
		        timing_token_generation = EXCLUDED.timing_token_generation,
		        timing_token_processing = EXCLUDED.timing_token_processing;
		
		    INSERT INTO usage
		    	(chat_session, tg_user_id, completion_tokens, prompt_tokens,
		    	total_tokens, reasoning_tokens,
		    	timing_token_generation, timing_token_processing)
		    VALUES
		    	(NULL, p_user_id, p_usage_completion_tokens, p_usage_prompt_tokens,
		    	p_usage_total_tokens, p_usage_reasoning_tokens,
		    	p_usage_timing_token_generation, p_usage_timing_token_processing)
		    ON CONFLICT (tg_user_id) WHERE chat_session IS NULL DO UPDATE SET 
		        completion_tokens = usage.completion_tokens + EXCLUDED.completion_tokens,
		        prompt_tokens = usage.prompt_tokens + EXCLUDED.prompt_tokens,
		        total_tokens = usage.total_tokens + EXCLUDED.total_tokens,
		        reasoning_tokens = usage.reasoning_tokens + EXCLUDED.reasoning_tokens,
		        timing_token_generation = usage.timing_token_generation + EXCLUDED.timing_token_generation,
		        timing_token_processing = usage.timing_token_processing + EXCLUDED.timing_token_processing;
		
		END;
		$$ LANGUAGE plpgsql;
    `)
	return err
}

func (s *Service) UpdateHistory(
	userId, endpointId, chatId, threadId int64,
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
		`, userId, endpointId, chatId, threadId, model, contentBytes)
	return err
}

func (s *Service) UpdateUsage(
	userId, endpointId, chatId, threadId int64,
	model string, usage map[string]any,
) error {
	var usageStruct Usage
	if val, ok := usage["CompletionTokens"]; ok {
		usageStruct.CompletionTokens = val.(int)
	}
	if val, ok := usage["PromptTokens"]; ok {
		usageStruct.PromptTokens = val.(int)
	}
	if val, ok := usage["TotalTokens"]; ok {
		usageStruct.TotalTokens = val.(int)
	}
	if val, ok := usage["ReasoningTokens"]; ok {
		usageStruct.ReasoningTokens = val.(int)
	}
	if val, ok := usage["TimingPromptProcessing"]; ok {
		usageStruct.TimingPromptProcessing = val.(float64)
	}
	if val, ok := usage["TimingTokenGeneration"]; ok {
		usageStruct.TimingTokenGeneration = val.(float64)
	}

	_, err := s.DBHandler.DB.Exec(`
		SELECT upsert_chat_session_and_usage($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`,
		userId, endpointId, chatId, threadId, model,
		usageStruct.CompletionTokens, usageStruct.PromptTokens,
		usageStruct.TotalTokens, usageStruct.ReasoningTokens,
		usageStruct.TimingPromptProcessing, usageStruct.TimingTokenGeneration,
	)
	return err
}

func (s *Service) DropUsage(
	userId, endpointId, chatId, threadId int64, model string,
) error {
	_, err := s.DBHandler.DB.Exec(`
    DELETE FROM usage
    WHERE chat_session IN (
        SELECT id
        FROM chat_sessions
        WHERE tg_user_id = $1 AND
        	endpoint = $2 AND
        	chat_id = $3 AND
        	thread_id = $4 AND
        	model = $5
    )`, userId, endpointId, chatId, threadId, model)
	return err
}

func (s *Service) GetUsage(
	userId, endpointId, chatId, threadId int64, model string,
) (*Usage, *Usage, *Usage, error) {
	row := s.DBHandler.DB.QueryRow(`
    	SELECT completion_tokens, prompt_tokens,
    		total_tokens, reasoning_tokens,
    		timing_token_generation, timing_token_processing
    	FROM usage
    	WHERE tg_user_id = $1 AND chat_session IS NULL
	`, userId)
	var globalUsage Usage
	err := row.Scan(
		&globalUsage.CompletionTokens, &globalUsage.PromptTokens,
		&globalUsage.TotalTokens, &globalUsage.TotalTokens,
		&globalUsage.TimingTokenGeneration, &globalUsage.TimingPromptProcessing)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, nil, err
	}

	row = s.DBHandler.DB.QueryRow(`
    	SELECT u.completion_tokens, u.prompt_tokens,
    		u.total_tokens, u.reasoning_tokens,
    		u.timing_token_generation, u.timing_token_processing
    	FROM usage u
    	JOIN chat_sessions cs ON u.chat_session = cs.id
    	WHERE cs.tg_user_id = $1
    	  AND cs.endpoint = $2
    	  AND cs.chat_id = $3
    	  AND cs.thread_id = $4
    	  AND cs.model = $5
	`, userId, endpointId, chatId, threadId, model) // Pass parameters
	var sessionUsage Usage
	err = row.Scan(
		&sessionUsage.CompletionTokens, &sessionUsage.PromptTokens,
		&sessionUsage.TotalTokens, &sessionUsage.TotalTokens,
		&sessionUsage.TimingTokenGeneration, &sessionUsage.TimingPromptProcessing)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, nil, err
	}

	row = s.DBHandler.DB.QueryRow(`
    	SELECT u.completion_tokens, u.prompt_tokens,
    		u.total_tokens, u.reasoning_tokens,
    		u.timing_token_generation, u.timing_token_processing
    	FROM last_usage u
    	JOIN chat_sessions cs ON u.chat_session = cs.id
    	WHERE cs.tg_user_id = $1
    	  AND cs.endpoint = $2
    	  AND cs.chat_id = $3
    	  AND cs.thread_id = $4
    	  AND cs.model = $5
	`, userId, endpointId, chatId, threadId, model) // Pass parameters
	var lastUsage Usage
	err = row.Scan(
		&lastUsage.CompletionTokens, &lastUsage.PromptTokens,
		&lastUsage.TotalTokens, &lastUsage.TotalTokens,
		&lastUsage.TimingTokenGeneration, &lastUsage.TimingPromptProcessing)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil, nil, err
	}
	return &globalUsage, &sessionUsage, &lastUsage, nil
}

func (s *Service) GetHistory(
	userId, endpointId, chatId, threadId int64, model string,
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
		`, userId, endpointId, chatId, threadId, model)
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

// TODO: debug
func (s *Service) DropHistory(
	userId, endpointId, chatId, threadId int64, model string,
) error {
	endpointId = 1 //TODO: this variable is not set anywhere in user object, so it equals to 0. In our db we use 1, that's why query fails.
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
		`, userId, endpointId, chatId, threadId, model)
	rowsAffected, _ := res.RowsAffected()
	log.Printf("Query executed successfully. Rows affected: %d", rowsAffected)
	return err
}

func (s *Service) UpdateModel(userId int64, model *string) error {
	var modelValue interface{}
	if model != nil {
		modelValue = *model
	} else {
		modelValue = nil
	}
	_, err := s.DBHandler.DB.Exec(`INSERT INTO ai_sessions
			(tg_user_id, model)
		VALUES
			($1, $2)
		ON CONFLICT(tg_user_id) DO UPDATE SET
			model = $2
		`, userId, modelValue)
	return err
}

// also create new session?
func (s *Service) UpdateEndpoint(userId int64, endpointId *int64) error {
	var endpointIdValue interface{}
	if endpointId != nil {
		endpointIdValue = *endpointId
	} else {
		endpointIdValue = nil
	}
	_, err := s.DBHandler.DB.Exec(`INSERT INTO ai_sessions
			(tg_user_id, endpoint)
		VALUES
			($1, $2)
		ON CONFLICT(tg_user_id) DO UPDATE SET
			endpoint = $2
		`, userId, endpointIdValue)
	return err
}

func (s *Service) GetSession(userId int64) (AISession, error) {
	var model, endpointName, endpointURL sql.NullString
	var endpointId, endpointAuthMethod sql.NullInt64
	var session AISession

	err := s.DBHandler.DB.QueryRow(`SELECT
			lt.model,
			rt.id,
			rt.name,
			rt.url,
			rt.auth_method
		FROM
			ai_sessions lt
		LEFT JOIN
			endpoints rt ON lt.endpoint = rt.id
		WHERE tg_user_id = $1`, userId).Scan(
		&model, &endpointId, &endpointName,
		&endpointURL, &endpointAuthMethod,
	)

	if model.Valid {
		session.Model = &model.String
	}
	if endpointId.Valid &&
		endpointName.Valid &&
		endpointURL.Valid &&
		endpointAuthMethod.Valid {
		var endpoint Endpoint

		endpoint.ID = endpointId.Int64
		endpoint.Name = endpointName.String
		endpoint.URL = endpointURL.String
		endpoint.AuthMethod = endpointAuthMethod.Int64

		session.Endpoint = &endpoint
	}

	return session, err
}

// check if session exists
func (s *Service) CheckSession(userId int64) bool {
	_, err := s.GetSession(userId)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}

func (s *Service) CreateLSession(userId int64, model string, endpoint int8) error {
	log.Printf("CreateLSession called with userId: %d, model: %s, endpoint: %d", userId, model, endpoint)

	res, err := s.DBHandler.DB.Exec(`
		INSERT INTO ai_sessions (tg_user_id, model, endpoint)
		VALUES ($1, $2, $3)
		ON CONFLICT(tg_user_id) DO UPDATE SET
		model = $2
		RETURNING tg_user_id
	`, userId, model, endpoint)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	log.Printf("Query executed successfully. Rows affected: %d", rowsAffected)

	return nil
}

func (s *Service) DeleteLSession(userId int64) error {
	_, err := s.DBHandler.DB.Exec(`
    DELETE FROM ai_sessions
    WHERE tg_user_id = $1
    `, userId)
	return err
}

func (s *Service) GetToken(userId, authMethod int64) (string, error) {
	token := ""
	err := s.DBHandler.DB.QueryRow(`SELECT token FROM auth
		WHERE
			tg_user_id = $1 AND
			auth_method = $2`,
		userId, authMethod).Scan(&token)
	return token, err
}

func (s *Service) InsertToken(userId, authMethod int64, token string) error {
	res, err := s.DBHandler.DB.Exec(`INSERT INTO auth
		(tg_user_id, auth_method, token)
		VALUES ($1, $2, $3)`, userId, authMethod, token)
	if err != nil {
		fmt.Println("Error inserting token:", err)
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("Insert token fired. Rows affected:", rowsAffected)
	return nil
}

func (s *Service) DeleteToken(userId, authMethod int64) error {
	_, err := s.DBHandler.DB.Exec(`DELETE FROM auth
		WHERE
			tg_user_id = $1 AND
			auth_method = $2
		`, userId, authMethod)
	return err
}

func (s *Service) GetEndpoints() ([]Endpoint, error) {
	rows, err := s.DBHandler.DB.Query(`SELECT id, name, url, auth_method FROM endpoints`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	endpoints := []Endpoint{}
	for rows.Next() {
		endpoint := Endpoint{}
		if err := rows.Scan(
			&endpoint.ID, &endpoint.Name, &endpoint.URL, &endpoint.AuthMethod,
		); err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}
