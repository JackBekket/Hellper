package database

import "database/sql"

func (s *Service) UpdateUsage(
	userID, endpointId, chatId, threadId int64,
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
		userID, endpointId, chatId, threadId, model,
		usageStruct.CompletionTokens, usageStruct.PromptTokens,
		usageStruct.TotalTokens, usageStruct.ReasoningTokens,
		usageStruct.TimingPromptProcessing, usageStruct.TimingTokenGeneration,
	)
	return err
}

func (s *Service) DropUsage(
	userID, endpointId, chatId, threadId int64, model string,
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
    )`, userID, endpointId, chatId, threadId, model)
	return err
}

func (s *Service) GetUsage(
	userID, endpointId, chatId, threadId int64, model string,
) (*Usage, *Usage, *Usage, error) {
	row := s.DBHandler.DB.QueryRow(`
    	SELECT completion_tokens, prompt_tokens,
    		total_tokens, reasoning_tokens,
    		timing_token_generation, timing_token_processing
    	FROM usage
    	WHERE tg_user_id = $1 AND chat_session IS NULL
	`, userID)
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
	`, userID, endpointId, chatId, threadId, model) // Pass parameters
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
	`, userID, endpointId, chatId, threadId, model) // Pass parameters
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
