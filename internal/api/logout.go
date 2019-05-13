package api

// Contains function for logout from github aggregator.

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
	"github.com/777777miSSU7777777/github-aggregator/pkg/time/timeutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

// Logout logs out user from app.
func Logout(rw http.ResponseWriter, req *http.Request) {
	token.GetTokenService().DeleteToken()

	session.GetSessionService().CloseSession()

	logger.Log(
		"method", "Logout",
		"time", timeutil.GetCurrentTime(),
		"info", "Logout",
	)
}
