package helper

import (
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func GetElasticIdFromResponse(response *esapi.Response) (string, error) {
	responseMap := make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&responseMap); err != nil {
		return "", err
	}

	hits := responseMap["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) > 0 {
		return hits[0].(map[string]interface{})["_id"].(string), nil
	}
	return "", errors.New("no hits found")
}
