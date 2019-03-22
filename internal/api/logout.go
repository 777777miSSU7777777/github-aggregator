package api

import (
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"

)


// Logout app user logouts.
func Logout(rw http.ResponseWriter, req *http.Request) {
	err := webtokenservice.DeleteToken(rw,req); if err != nil{
		log.Println(err)
	}

	log.Println("Logout")
	
	http.Redirect(rw, req, "/", 301)
}
