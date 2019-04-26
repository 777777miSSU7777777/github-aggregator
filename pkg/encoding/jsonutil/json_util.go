// Package jsonutil implements function for unmarshalling data to map.
package jsonutil

import (
	"encoding/json"
)

// BytesToMap transforms jsonBytes to map.
func BytesToMap(jsonBytes []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &jsonMap)

	if err != nil {
		return nil, err
	}

	return jsonMap, err
}
