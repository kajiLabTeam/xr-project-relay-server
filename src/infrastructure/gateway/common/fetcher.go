package common_gateway

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostFetcher[RequestType, ResponseType any](requestStruct RequestType, endpoint string) (*ResponseType, error) {
	jsonData, err := json.Marshal(requestStruct)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseBody ResponseType
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
