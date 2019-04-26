package randutil

// Package randutil implements function for generation of random bytes slice.

import (
	"crypto/rand"
)

var randRead = rand.Read

// GenerateRandomBytes generates random byte array with length of size.
func GenerateRandomBytes(size int) ([]byte, error) {
	randomBytes := make([]byte, size)

	_, err := randRead(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}
