package base64util

import (
	"encoding/base64"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
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

	assert.Equal(t, encoded, stringData)

	tearDown()
}

func TestEncode__DifferentStrings__Failed(t *testing.T) {
	tearUp()

	encoded := Encode(byteData)

	anotherBytes, err := randutil.GenerateRandomBytes(16)

	if err != nil {
		log.Error.Fatalln(err)
	}

	anotherString := Encode(anotherBytes)

	assert.NotEqual(t, encoded, anotherString)

	tearDown()
}

func TestDecode__SameBytes__Successful(t *testing.T) {
	tearUp()

	decoded, err := Decode(stringData)

	if err != nil {
		log.Error.Fatalln(err)
	}

	assert.Equal(t, decoded, byteData)

	tearDown()
}

func TestDecode__DifferentBytes__Failed(t *testing.T) {
	tearUp()

	anotherDecoded, err := base64.StdEncoding.DecodeString("YWJjMTIzIT8kKiYoKSctPUB+")

	if err != nil {
		log.Error.Fatalln(err)
	}

	decoded, err := Decode(stringData)

	if err != nil {
		log.Error.Fatalln(err)
	}

	assert.NotEqual(t, decoded, anotherDecoded)

	tearDown()
}

func TestDecode__IncorrectString__Error(t *testing.T) {
	_, err := Decode("XXXXXaGVsbG8=")

	assert.Error(t, err)
}
