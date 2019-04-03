package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/stretchr/testify/assert"
)

func recoverToken() {
	if r := recover(); r != nil {
		log.Info.Println("recovered from", r)
	}
}

func TestAuth__ValidToken__Saved(t *testing.T) {
	defer recoverToken()

	webtokenservice.SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceIV(IV)

	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		resp := &http.Response{}
		resp.StatusCode = 200
		return resp, nil
	}

	randomBytes, _ := randutil.GenerateRandomBytes(16)
	testToken := base64util.Encode(randomBytes)
	req := &http.Request{Form: url.Values{}}
	req.Form.Add("access_token", testToken)

	rw := httptest.NewRecorder()

	Auth(rw, req)

	cookieToken, _ := webtokenservice.GetToken(req)

	assert.Equal(t, testToken, cookieToken)
}

func TestAuth__InvalidToken__NotSaved(t *testing.T) {
	defer recoverToken()

	webtokenservice.SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceIV(IV)

	original := httpGet

	defer func() {
		httpGet = original
	}()

	httpGet = func(url string) (*http.Response, error) {
		resp := &http.Response{}
		resp.StatusCode = 401
		return resp, nil
	}

	randomBytes, _ := randutil.GenerateRandomBytes(16)
	testToken := base64util.Encode(randomBytes)
	req := &http.Request{Form: url.Values{}}
	req.Form.Add("access_token", testToken)

	rw := httptest.NewRecorder()

	Auth(rw, req)

	_, err := webtokenservice.GetToken(req)

	assert.Error(t, err)
}
