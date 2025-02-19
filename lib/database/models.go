package database

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type OpenAIDataObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

type OpenAIModelsResponse struct {
	Data []OpenAIDataObject `json:"data"`
}

func (s *Service) GetModelsList(endpoint, token string) ([]string, error) {
	modelsList := []string{}
	urlPath, err := url.JoinPath(endpoint, "models")
	if err != nil {
		return modelsList, err
	}
	req, err := http.NewRequest("GET", urlPath, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return modelsList, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return modelsList, err
	}

	modelsResp := OpenAIModelsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&modelsResp)
	if err != nil {
		return modelsList, err
	}

	for _, obj := range modelsResp.Data {
		modelsList = append(modelsList, obj.ID)
	}

	return modelsList, nil
}
