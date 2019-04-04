package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
	"github.com/stretchr/testify/assert"
)

func TestLogout__LoggedIn__TokenDeleted(t *testing.T) {
	defer recoverRedirect()

	webtokenservice.SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	webtokenservice.SetCryptoServiceIV(IV)

	randomBytes, _ := randutil.GenerateRandomBytes(16)
	testToken := base64util.Encode(randomBytes)

	rw := httptest.NewRecorder()

	webtokenservice.SaveToken(rw, testToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	Logout(rw, req)

	_, err := webtokenservice.GetToken(req)

	assert.Error(t, err)
}

func TestLogout__NotLoggedIn__CookieNotFoundError(t *testing.T) {
	defer recoverRedirect()

	rw := httptest.NewRecorder()

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	Logout(rw, req)

	_, err := webtokenservice.GetToken(req)

	assert.Error(t, err)
}
