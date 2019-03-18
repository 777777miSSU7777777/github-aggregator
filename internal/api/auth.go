package api

import (
	"log"
	"net/http"
	"time"

	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
)

//Auth Saves to cookie Github API access token from form on index page.
func Auth(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		accessToken := req.FormValue("access_token")
		resp, err := http.Get("https://api.github.com/user?access_token=" + accessToken)

		if err != nil {
			log.Println(err)
		}

		if resp.StatusCode == 200 {
			cookie.SaveCookie(rw, "access_token", accessToken, time.Hour)
			log.Println("Authentication is successful")
		} else if resp.StatusCode == 401 {
			log.Println("Authentication is failed")
		}

		http.Redirect(rw, req, "/", 301)
	}
}
