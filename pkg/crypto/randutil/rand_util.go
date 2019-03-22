package randutil

import (
	"crypto/rand"
)


// GenerateRandomBytes generates random byte array with length of size.
// Returns error if any occurs while the byte array filling.
func GenerateRandomBytes(size int)([]byte,error){
	randomBytes := make([]byte, size)

	_, err := rand.Read(randomBytes); if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

