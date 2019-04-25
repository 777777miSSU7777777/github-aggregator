// Package index implements function for pull requests page render.
package pulls

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

type AuthState struct {
	Auth bool
}

// Render renders pull requests page.
// If occurs any error (besides cookie not found) to client will be returned "Internal server error".
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn := tokenservice.GetToken()

	err := view.GetTemplates().ExecuteTemplate(rw, "pulls.gohtml", AuthState{Auth: tkn != ""})

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Println("Pulls page successfuly rendered")
	}
}
