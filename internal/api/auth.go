package api

// Contains function for authentication in github aggregator.

import (
	"fmt"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

// Auth authenticates user with provided Github API access token.
func Auth(rw http.ResponseWriter, req *http.Request) {
	tkn := req.FormValue(constants.AccessToken)
	resp, err := httpGet(fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		log.Error.Println(err)
	}

	if resp.StatusCode == 200 {
		tokenservice.SaveToken(tkn)
		session.GetSessionService().StartSession(tkn)
		log.Info.Println("Authentication is successful")
	} else if resp.StatusCode == 401 {
		log.Info.Println("Authentication is failed")
	}

	http.Redirect(rw, req, "/", 301)
}
