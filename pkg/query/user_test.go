package query

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
)

func TestGetUser__SameBytes__Equal(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	httpGet = func(url string) (*http.Response, error) {
		response := &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(randomBytes))}
		return response, nil
	}

	responseBytes, _ := GetUser("123")

	assert.Equal(t, randomBytes, responseBytes)
}

func TestGetUser__DifferentBytes__NotEqual(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	httpGet = func(url string) (*http.Response, error) {
		response := &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer(randomBytes))}
		return response, nil
	}

	responseBytes, _ := GetUser("123")

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, responseBytes)
}

func TestGetUser__HttpGetOccursError__Error(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		return nil, errors.New("HTTP GET ERROR")
	}

	_, err := GetUser("123")

	assert.Error(t, err)
}

func TestGetUser__ReadResponseBodyOccursError_Error(t *testing.T) {
	original := readResponseBody

	defer func() {
		readResponseBody = original
	}()

	readResponseBody = func(resp *http.Response) ([]byte, error) {
		return nil, errors.New("HTTP response body read error")
	}

	_, err := GetUser("123")

	assert.Error(t, err)
}
