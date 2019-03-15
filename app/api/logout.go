package api

import (
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
)

//Logout removes Github API access token from cookie.
func Logout(rw http.ResponseWriter, req *http.Request) {
	err := cookie.DeleteCookie(rw, req, "access_token")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Logout")
	http.Redirect(rw, req, "/", 301)
}
