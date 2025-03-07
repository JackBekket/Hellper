package database

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
			tg_user_id BIGINT NOT NULL,
			auth_method INT UNIQUE REFERENCES auth_methods(id),
			token TEXT NOT NULL
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS ai_sessions (
			tg_user_id BIGINT PRIMARY KEY,
			model TEXT,
			endpoint INT REFERENCES endpoints(id)	
		)`)
	if err != nil {
		return err
	}
	_, err = s.DBHandler.DB.Exec(`
		CREATE TABLE IF NOT EXISTS chat_sessions (
			id SERIAL PRIMARY KEY,
			tg_user_id BIGINT NOT NULL,
			model TEXT NOT NULL,
			endpoint INT NOT NULL REFERENCES endpoints(id),
			chat_id BIGINT NOT NULL,
			thread_id BIGINT NOT NULL
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
			tg_user_id BIGINT NOT NULL,
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
		    p_chat_id BIGINT,
		    p_thread_id BIGINT,
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
