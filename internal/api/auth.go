package api

// Contains function for authentication in github aggregator.

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

const (
	// OAUTH2_HEADER key for Oauth token in request header.
	OAUTH2_HEADER = "Authorization"

	// OAUTH2_PREFIX prefix for Oauth token.
	OAUTH2_PREFIX = "Bearer "

	// AUTH_URL url to get token auth status.
	AUTH_URL = "https://api.github.com/user"

	// ACCESS_TOKEN key for access token value.
	ACCESS_TOKEN = "access_token"
)

// Auth authenticates user with provided Github API access token.
func Auth(rw http.ResponseWriter, req *http.Request) {
	tkn := req.FormValue(ACCESS_TOKEN)

	tokenservice.GetTokenService().SaveToken(tkn)

	if tokenservice.GetTokenService().GetToken() != "" {
		session.GetSessionService().StartSession(tkn)
		log.Info.Println("Authentication is successful")
		http.Redirect(rw, req, "/", 301)
	} else {
		log.Info.Println("Authentication is failed")
		http.Redirect(rw, req, "/login", 301)
	}

}
