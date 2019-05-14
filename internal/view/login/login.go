// Package login implements function for login page render.
package login

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

// Render renders login page.
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn := token.GetTokenService().GetToken()

	if tkn != "" {
		http.Redirect(rw, req, "/", 301)
		return
	}

	err := view.GetTemplates().ExecuteTemplate(rw, "login.gohtml", view.AuthState{Auth: tkn != ""})

	if err != nil {
		view.Logger().Warnln(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		view.Logger().Infoln("Login page rendered")
	}

}
