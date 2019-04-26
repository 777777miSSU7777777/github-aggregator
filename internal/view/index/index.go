// Package index implements function for index page render.
package index

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

// Render renders index page.
// If user not authenticated, will be rendered form with token field and submit button.
// If user is authenticated, will be rendered user avatar, link to profile,
// button for pop-up window with scopes for provided token  and logout button.
// If occurs any error (besides cookie not found) to client will be returned "Internal server error".
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn := tokenservice.GetToken()

	err := view.GetTemplates().ExecuteTemplate(rw, "index.gohtml", view.AuthState{Auth: tkn != ""})

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Println("Index page successfuly rendered")
	}
}
