package bodyutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
)

func TestReadBody__SameBody__Equal(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	body := ioutil.NopCloser(bytes.NewBuffer(randomBytes))

	readBytes, _ := readBody(body)

	assert.Equal(t, randomBytes, readBytes)
}

func TestReadBody__AnotherBytes__NotEqual(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	body := ioutil.NopCloser(bytes.NewBuffer(randomBytes))

	readBytes, _ := readBody(body)

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, readBytes)
}

func TestReadBody__MockedReadAll__Error(t *testing.T) {
	original := ioutilReadAll

	defer func() {
		ioutilReadAll = original
	}()

	ioutilReadAll = func(r io.Reader) ([]byte, error) {
		return nil, io.EOF
	}

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	body := ioutil.NopCloser(bytes.NewBuffer(randomBytes))

	_, err := readBody(body)

	assert.Error(t, err)
}

func TestReadRequestBody__SameBody__Equal(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(randomBytes))

	readBytes, _ := ReadRequestBody(req)

	assert.Equal(t, randomBytes, readBytes)
}

func TestReadRequestBody__AnotherBytes__NotEqual(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(randomBytes))

	readBytes, _ := ReadRequestBody(req)

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, readBytes)
}

func TestReadResponseBody__SameBody__Equal(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	resp := &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(randomBytes))}

	readBytes, _ := ReadResponseBody(resp)

	assert.Equal(t, randomBytes, readBytes)
}

func TestReadResponseBody__AnotherBytes__Equal(t *testing.T) {
	randomBytes, _ := randutil.GenerateRandomBytes(16)

	resp := &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(randomBytes))}

	readBytes, _ := ReadResponseBody(resp)

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, readBytes)
}
