package cryptofactory

import (
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/aes"
	"github.com/stretchr/testify/assert"
)

func TestNew__aes_String__AES(t *testing.T) {
	testCryptoService := New("aes")
	aesInstance := aes.AES{}

	assert.ObjectsAreEqual(aesInstance, testCryptoService)
}

func TestNew__Default__AES(t *testing.T) {
	testCryptoService := New("123")
	aesInstance := aes.AES{}

	assert.ObjectsAreEqual(aesInstance, testCryptoService)
}
