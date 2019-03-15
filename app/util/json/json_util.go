package json

import (
	"encoding/json"
)

//BytesToMap returns string : interface map from JSON bytes.
// Also returns error if happen.
func BytesToMap(jsonBytes []byte)(map[string]interface{}, error){
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, &jsonMap)
	
	if err != nil {
		return nil, err
	}

	return jsonMap, err
}