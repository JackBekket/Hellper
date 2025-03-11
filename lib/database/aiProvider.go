package database

// Returns a list of localAI service names that the user can use
func (s *Service) GetAIProvidersName(authMethod int64) ([]string, error) {
	rows, err := s.DBHandler.DB.Query("SELECT name FROM endpoints WHERE auth_method = $1", authMethod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, rows.Err()
}

// Returns all information for a specific service
func (s *Service) GetAIProvider(providerName string) (*Endpoint, error) {
	query := `SELECT id, name, url, auth_method FROM endpoints WHERE name = $1`
	row := s.DBHandler.DB.QueryRow(query, providerName)

	var provider Endpoint
	if err := row.Scan(&provider.ID, &provider.Name, &provider.BaseURL, &provider.AuthMethod); err != nil {
		return nil, err
	}
	return &provider, nil
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
			&endpoint.ID, &endpoint.Name, &endpoint.BaseURL, &endpoint.AuthMethod,
		); err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}
