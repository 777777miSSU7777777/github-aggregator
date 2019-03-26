// Package aes contains an implentation of CryptoService interface.
package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

// AES is an implementation of Advanced Encryption Standart also known as Rjindael.
type AES struct {
	// key is a private field for secret key.
	// Correct key is byte array with length 16, 24 or 32.
	key []byte

	// initiliazationVector is private field for IV.
	initializationVector []byte
}

// SetKey sets secret key for AES instance.
// Key should be presented as byte array.
func (a *AES) SetKey(key []byte) {
	a.key = key
}

// SetIV sets IV for AES instance.
// IV should be presented as byte array.
func (a *AES) SetIV(IV []byte) {
	a.initializationVector = IV
}

// Encrypt encrypts data using AES with CTR block mode.
// Data should be presented as byte array.
// Encrypted data is presented as byte array.
// If aes.NewCipher occurs any error, this will be returned.
func (a AES) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, a.initializationVector)
	stream.XORKeyStream(data, data)

	return data, nil
}

// Decrypt decrypts encrypted data using AES with CTR block mode.
// Encrypted data should be presented as byte array.
// Decrypted data is presented as byte array.
// If aes.NewCipher occurs any error, this will be returned.
func (a AES) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, a.initializationVector)
	stream.XORKeyStream(data, data)

	return data, nil
}
