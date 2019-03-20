package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

type AES struct {
	key []byte
	initializationVector []byte
}


func (a *AES) SetKey(key []byte){
	a.key = key
}

func (a *AES) SetIV(IV []byte){
	a.initializationVector = IV
}


func (a AES) Encrypt(data []byte)([]byte,error){
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, a.initializationVector)
	stream.XORKeyStream(data,data)

	return data, nil
}


func (a AES) Decrypt(data []byte)([]byte,error){
	block, err := aes.NewCipher(a.key)

	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, a.initializationVector)
	stream.XORKeyStream(data,data)

	return data, nil
}






