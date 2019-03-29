package base64util

import (
	"encoding/base64"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
)

var stringData string
var byteData []byte

func tearUp() {
	byteData, _ = randutil.GenerateRandomBytes(16)
	stringData = Encode(byteData)
}

func tearDown() {
	stringData = ""
	byteData = nil
}

func TestEncode__SameString__Successful(t *testing.T) {
	tearUp()

	encoded := Encode(byteData)

	assert.Equal(t, stringData, encoded)

	tearDown()
}

func TestEncode__DifferentStrings__Failed(t *testing.T) {
	tearUp()

	encoded := Encode(byteData)

	anotherBytes, _ := randutil.GenerateRandomBytes(16)

	anotherString := Encode(anotherBytes)

	assert.NotEqual(t, anotherString, encoded)

	tearDown()
}

func TestDecode__SameBytes__Successful(t *testing.T) {
	tearUp()

	decoded, _ := Decode(stringData)

	assert.Equal(t, byteData, decoded)

	tearDown()
}

func TestDecode__DifferentBytes__Failed(t *testing.T) {
	tearUp()

	anotherDecoded, _ := base64.StdEncoding.DecodeString("YWJjMTIzIT8kKiYoKSctPUB+")

	decoded, _ := Decode(stringData)

	assert.NotEqual(t, anotherDecoded, decoded)

	tearDown()
}

func TestDecode__IncorrectString__Error(t *testing.T) {
	_, err := Decode("XXXXXaGVsbG8=")

	assert.Error(t, err)
}
