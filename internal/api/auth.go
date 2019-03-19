package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
)


func Auth(rw http.ResponseWriter, req *http.Request) {
	tkn := req.FormValue("access_token")
	resp, err := http.Get( fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		cookieutil.SaveCookie(rw, "access_token", tkn, time.Hour)
		log.Println("Authentication is successful")
	} else if resp.StatusCode == 401 {
		log.Println("Authentication is failed")
	}

	http.Redirect(rw, req, "/", 301)
}
