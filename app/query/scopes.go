package query

import (
	"net/http"
	"strings"
	"log"
	"github-aggregator-web/app/util/http/header"
	"github-aggregator-web/app/util/http/cookie"
)

func GetScopes(req *http.Request)([]string){
	access_token, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		log.Println(err.Error())
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + access_token)

	respHeader := header.ReadResponseHeader(resp)

	return strings.Split(respHeader["X-Oauth-Scopes"][0],",")
}
