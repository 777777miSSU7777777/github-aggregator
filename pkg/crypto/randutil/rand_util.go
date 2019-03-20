package randutil

import (
	"crypto/rand"
)

func GenerateRandomBytes(size int)([]byte,error){
	randomBytes := make([]byte, size)

	_, err := rand.Read(randomBytes); if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

