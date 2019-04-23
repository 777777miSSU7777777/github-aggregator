package api

// Contains function for logout from github aggregator.

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

// Logout logs out user from app.
func Logout(rw http.ResponseWriter, req *http.Request) {
	tokenservice.DeleteToken()

	log.Info.Println("Logout")

	http.Redirect(rw, req, "/", 301)
}
