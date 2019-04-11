// Package login implements function for login page render.
package login

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

type LoginState struct {
	Auth bool
}

// Render renders login page.
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn, err := webtokenservice.GetToken(req)

	if err != nil {
		log.Warning.Println(err)
	}

	if tkn != "" {
		http.Redirect(rw, req, "/", 301)
		return
	}

	err = view.GetTemplates().ExecuteTemplate(rw, "login.gohtml", LoginState{Auth: tkn != ""})

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Println("Login page successfuly rendered")
	}

}
