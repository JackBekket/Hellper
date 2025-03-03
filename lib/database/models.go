package database

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIDataObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

type OpenAIModelsResponse struct {
	Data []OpenAIDataObject `json:"data"`
}

func (s *Service) GetModelsList(url, localAIToken string) ([]string, error) {
	modelsList := []string{}

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+localAIToken)
	if err != nil {
		return modelsList, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return modelsList, err
	}

	fmt.Println(resp.Status)

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
