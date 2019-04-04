package query

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetScopes__SameScopes__Equal(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		respHeader := map[string][]string{}

		respHeader[constants.Scopes] = []string{"user,repository"}

		resp := &http.Response{Header: http.Header(respHeader)}

		return resp, nil
	}

	scopes, _ := GetScopes("123")

	header := map[string][]string{}

	header[constants.Scopes] = []string{"user,repository"}

	splittedHeaderScopes := strings.Split(header[constants.Scopes][0], ",")

	assert.Equal(t, splittedHeaderScopes, scopes)
}

func TestGetScopes__DifferentScopes__NotEqual(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		respHeader := map[string][]string{}

		respHeader[constants.Scopes] = []string{"user,repository"}

		resp := &http.Response{Header: http.Header(respHeader)}

		return resp, nil
	}

	scopes, _ := GetScopes("123")

	header := map[string][]string{}

	header[constants.Scopes] = []string{"user"}

	splittedHeaderScopes := strings.Split(header[constants.Scopes][0], ",")

	assert.NotEqual(t, splittedHeaderScopes, scopes)
}

func TestGetScopes__HttpGetOccursError_Error(t *testing.T) {
	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		return nil, errors.New("HTTP GET ERROR")
	}

	_, err := GetScopes("123")

	assert.Error(t, err)
}
