package database

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func (s *Service) GetUserTokensWithAuthMethods(userID int64) (map[int64]string, error) {
	query := `SELECT auth_method, token FROM auth WHERE tg_user_id = $1`
	rows, err := s.DBHandler.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := make(map[int64]string)
	for rows.Next() {
		var authMethod int64
		var token string
		if err := rows.Scan(&authMethod, &token); err != nil {
			return nil, err
		}
		tokens[authMethod] = token
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *Service) GetToken(userID, authMethod int64) (string, error) {
	token := ""
	err := s.DBHandler.DB.QueryRow(`SELECT token FROM auth
		WHERE
			tg_user_id = $1 AND
			auth_method = $2`,
		userID, authMethod).Scan(&token)
	return token, err
}

func (s *Service) ExistInAuth(userID int64) bool {
	var exists bool
	err := s.DBHandler.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM auth WHERE tg_user_id = $1)",
		userID,
	).Scan(&exists)
	if err != nil {
		log.Warn().Err(err).Msg("failed to check existence in auth")
		return false
	}
	return exists
}

func (s *Service) CheckToken(userID, authMethod int64) bool {
	token, err := s.GetToken(userID, authMethod)
	if err != nil {
		fmt.Println(err)
	}
	tokenPresent := false
	if token != "" {
		tokenPresent = true
	}
	return tokenPresent
}

func (s *Service) InsertToken(userID, authMethod int64, token string) error {
	res, err := s.DBHandler.DB.Exec(`
		INSERT INTO auth (tg_user_id, auth_method, token)
		VALUES ($1, $2, $3)
	`, userID, authMethod, token)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting token")
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Warn().Err(err).Msg("Error getting rows affected")
	} else {
		log.Info().Int64("rowsAffected", rowsAffected).Msg("Token inserted successfully")
	}

	return nil
}

func (s *Service) DeleteToken(userID, authMethod int64) error {
	_, err := s.DBHandler.DB.Exec(`DELETE FROM auth
		WHERE
			tg_user_id = $1 AND
			auth_method = $2
		`, userID, authMethod)
	return err
}
