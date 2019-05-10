// Package pulls implements function for pull requests page render.
package pulls

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

// Render renders pull requests page.
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn := token.GetTokenService().GetToken()

	err := view.GetTemplates().ExecuteTemplate(rw, "pulls.gohtml", view.AuthState{Auth: tkn != ""})

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Println("Pulls page successfuly rendered")
	}
}
