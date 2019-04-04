package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
)

var aesInstance *AES

func init() {
	aesInstance = &AES{}
}

func tearUp() {
	secretKey, _ := randutil.GenerateRandomBytes(16)

	aesInstance.key = secretKey

	IV, _ := randutil.GenerateRandomBytes(16)

	aesInstance.iv = IV
}

func tearDown() {
	aesInstance.key = nil
	aesInstance.iv = nil
}

func TestSetKey__SameKey__Equal(t *testing.T) {
	tearUp()

	key := aesInstance.key

	aesInstance.SetKey(key)

	assert.Equal(t, key, aesInstance.key)

	tearDown()
}

func TestSetKey__DifferentKeys__NotEqual(t *testing.T) {
	tearUp()

	key := aesInstance.key

	anotherKey, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(anotherKey)

	assert.NotEqual(t, key, aesInstance.key)

	tearDown()
}

func TestSetIV__SameIV__Equal(t *testing.T) {
	tearUp()

	iv := aesInstance.iv

	aesInstance.SetIV(iv)

	assert.Equal(t, iv, aesInstance.iv)

	tearDown()
}

func TestSetIV__DifferentIV__NotEqual(t *testing.T) {
	tearUp()

	iv := aesInstance.iv

	anotherIV, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetIV(anotherIV)

	assert.NotEqual(t, iv, aesInstance.iv)

	tearDown()
}

func TestEncrypt__SameKeyAndIV__Equal(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	block, _ := aes.NewCipher(aesInstance.key)

	stream := cipher.NewCTR(block, aesInstance.iv)

	decryptedData := make([]byte, len(encryptedData))

	stream.XORKeyStream(decryptedData, encryptedData)

	assert.Equal(t, testData, decryptedData)

	tearDown()
}

func TestEncrypt__AnotherKeyAndSameIV__NotEqual(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	anotherKey, _ := randutil.GenerateRandomBytes(16)

	block, _ := aes.NewCipher(anotherKey)

	stream := cipher.NewCTR(block, aesInstance.iv)

	decryptedData := make([]byte, len(encryptedData))

	stream.XORKeyStream(decryptedData, encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestEncrypt__SameKeyAndAnotherIV__Equal(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	block, _ := aes.NewCipher(aesInstance.key)

	anotherIV, _ := randutil.GenerateRandomBytes(16)

	stream := cipher.NewCTR(block, anotherIV)

	decryptedData := make([]byte, len(encryptedData))

	stream.XORKeyStream(decryptedData, encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestEncrypt__DifferentKeyAndIV__NotEqual(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	anotherKey, _ := randutil.GenerateRandomBytes(16)

	block, _ := aes.NewCipher(anotherKey)

	anotherIV, _ := randutil.GenerateRandomBytes(16)

	stream := cipher.NewCTR(block, anotherIV)

	decryptedData := make([]byte, len(encryptedData))

	stream.XORKeyStream(decryptedData, encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestEncrypt__NilKeyAndCorrectIV__Error(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(nil)

	_, err := aesInstance.Encrypt(testData)

	assert.Error(t, err)

	tearDown()
}

func TestEncrypt__CorrectKeyAndNilIV__Panics(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetIV(nil)

	assert.Panics(t, func() {
		_, _ = aesInstance.Encrypt(testData)
	})

	tearDown()
}

func TestEncrypt__NilKeyAndNilIV__Error(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(nil)
	aesInstance.SetIV(nil)

	_, err := aesInstance.Encrypt(testData)

	assert.Error(t, err)

	tearDown()
}

func TestDecrypt__SameKeyAndIV__Equal(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	decryptedData, _ := aesInstance.Decrypt(encryptedData)

	assert.Equal(t, testData, decryptedData)

	tearDown()
}

func TestDecrypt__AnotherKeyAndSameIV__NotEqual(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	anotherKey, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(anotherKey)

	decryptedData, _ := aesInstance.Decrypt(encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestDecrypt__SameKeyAndAnotherIV__NotEqual(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	anotherIV, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetIV(anotherIV)

	decryptedData, _ := aesInstance.Decrypt(encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestDecrypt__DifferentKeyAndIV__NotEqual(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	encryptedData, _ := aesInstance.Encrypt(testData)

	anotherKey, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(anotherKey)

	anotherIV, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetIV(anotherIV)

	decryptedData, _ := aesInstance.Decrypt(encryptedData)

	assert.NotEqual(t, testData, decryptedData)

	tearDown()
}

func TestDecrypt__NilKeyAndCorrectIV__Error(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(nil)

	_, err := aesInstance.Decrypt(testData)

	assert.Error(t, err)

	tearDown()
}

func TestDecrypt__CorrectKeyAndNilIV__Panics(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetIV(nil)

	assert.Panics(t, func() {
		_, _ = aesInstance.Decrypt(testData)
	})

	tearDown()
}

func TestDecrypt__NilKeyAndNilIV__Error(t *testing.T) {
	tearUp()

	testData, _ := randutil.GenerateRandomBytes(16)

	aesInstance.SetKey(nil)
	aesInstance.SetIV(nil)

	_, err := aesInstance.Decrypt(testData)

	assert.Error(t, err)

	tearDown()
}
