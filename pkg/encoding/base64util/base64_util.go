// Package base64util implements functions for string decoding and bytes encoding in base64 encoding.
package base64util

import (
	"encoding/base64"
)

// Encode encodes byte array into string with base64 encoding.
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Decode decodes string with base64 encoding into byte array.
func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return data, nil
}
