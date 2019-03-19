package api

import (
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
)

func Logout(rw http.ResponseWriter, req *http.Request) {
	err := cookieutil.DeleteCookie(rw, req, "access_token"); if err != nil {
		log.Println(err)
	}

	log.Println("Logout")
	
	http.Redirect(rw, req, "/", 301)
}
