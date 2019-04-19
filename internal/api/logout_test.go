package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
	"github.com/stretchr/testify/assert"
)

func TestLogout__LoggedIn__TokenDeleted(t *testing.T) {
	defer recoverRedirect()

	randomBytes, _ := randutil.GenerateRandomBytes(16)
	testToken := base64util.Encode(randomBytes)

	rw := httptest.NewRecorder()

	tokenservice.SaveToken(testToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	Logout(rw, req)

	token := tokenservice.GetToken()

	assert.Empty(t, token)
}

func TestLogout__NotLoggedIn__CookieNotFoundError(t *testing.T) {
	defer recoverRedirect()

	rw := httptest.NewRecorder()

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	Logout(rw, req)

	token := tokenservice.GetToken()

	assert.Empty(t, token)
}
