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

	// iv is private field for IV.
	iv []byte
}

// SetKey sets secret key for AES instance.
// Key should be presented as byte array.
func (a *AES) SetKey(key []byte) {
	a.key = key
}

// SetIV sets IV for AES instance.
// IV should be presented as byte array.
func (a *AES) SetIV(IV []byte) {
	a.iv = IV
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
	
	enData := make([]byte, len(data))
	stream := cipher.NewCTR(block, a.iv)
	stream.XORKeyStream(enData, data)

	return enData, nil
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

	deData := make([]byte, len(data))
	stream := cipher.NewCTR(block, a.iv)
	stream.XORKeyStream(deData, data)

	return deData, nil
}
