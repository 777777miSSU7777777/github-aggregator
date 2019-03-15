package query

import (
	"log"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/header"
)

//GetScopes returns scopes for provided token.
func GetScopes(req *http.Request) []string {
	accessToken, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		log.Println(err.Error())
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + accessToken)

	respHeader := header.ReadResponseHeader(resp)

	return strings.Split(respHeader["X-Oauth-Scopes"][0], ",")
}
