package base64util

import (
	"encoding/base64"
)


func Encode(b []byte)(string){
	return base64.StdEncoding.EncodeToString(b)
}


func Decode(s string)([]byte, error){
	data, err := base64.StdEncoding.DecodeString(s); if err != nil {
		return nil, err
	}

	return data, nil
}