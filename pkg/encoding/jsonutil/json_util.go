package jsonutil

import (
	"encoding/json"
)

// BytesToMap transforms jsonBytes to map.
// Map stores interface{} values.
// If json.Unmarshal occurs any error, this will be returned.
func BytesToMap(jsonBytes []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &jsonMap)

	if err != nil {
		return nil, err
	}

	return jsonMap, err
}
