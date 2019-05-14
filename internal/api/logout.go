package api

// Contains function for logout from github aggregator.

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"

	log "github.com/sirupsen/logrus"
)

// MakeLogoutAPIHandler returns func with logging which logs out user from app.
func MakeLogoutAPIHandler(logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		token.GetTokenService().DeleteToken()

		session.GetSessionService().CloseSession()

		logger.Infoln("Logout")
	}
}
