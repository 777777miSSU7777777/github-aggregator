package api

// Contains function for logout from github aggregator.

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

// Logout logs out user from app.
func Logout(rw http.ResponseWriter, req *http.Request) {
	token.GetTokenService().DeleteToken()

	session.GetSessionService().CloseSession()

	log.Info.Println("Logout")
}
