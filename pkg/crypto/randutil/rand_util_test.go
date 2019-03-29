package randutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomBytes__Size16__Successful(t *testing.T) {
	randomBytes, _ := GenerateRandomBytes(16)

	assert.Equal(t, 16, len(randomBytes))
}

func TestGenerateRandomBytes__Size16__Filled(t *testing.T) {
	randomBytes, _ := GenerateRandomBytes(16)

	zeroBytes := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	assert.NotEqual(t, zeroBytes, randomBytes)
}

func TestGenerateRandomBytes__Size0__Successful(t *testing.T) {
	randomBytes, _ := GenerateRandomBytes(0)

	assert.Equal(t, 0, len(randomBytes))
}

func TestGenerateRandomBytes__Size0__Empty(t *testing.T) {
	randomBytes, _ := GenerateRandomBytes(0)

	emptyByteArray := []byte{}

	assert.Equal(t, emptyByteArray, randomBytes)
}

func TestGenerateRandomBytes__NegativeSize__Panics(t *testing.T) {
	assert.Panics(t, func() {
		GenerateRandomBytes(-16)

	})
}

func TestGenerateRandomBytes__MockerRead__Error(t *testing.T) {
	original := randRead

	defer func() {
		randRead = original
	}()

	randRead = func(b []byte) (int, error) {
		return 0, errors.New("Error")
	}

	_, err := GenerateRandomBytes(16)

	assert.Error(t, err)
}
