package index

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/profilefactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Render renders index page.
// If user not authenticated, will be rendered form with token field and submit button.
// If user is authenticated, will be rendered user avatar, link to profile,
// button for pop-up window with scopes for provided token  and logout button.
// If occurs any error (besides cookie not found) to client will be returned "Internal server error".
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn, err := webtokenservice.GetToken(req)
	if err != nil {
		log.Warning.Println(err)
	}

	if tkn != "" {
		userBytes, err := query.GetUser(tkn)
		if err != nil {
			log.Warning.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}

		scopes, err := query.GetScopes(tkn)
		if err != nil {
			log.Warning.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}

		profile, err := profilefactory.New(userBytes, scopes)
		if err != nil {
			log.Warning.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}

		err = view.GetTemplates().ExecuteTemplate(rw, "index-authorized.gohtml", profile)
		if err != nil {
			log.Error.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}
	} else {
		err = view.GetTemplates().ExecuteTemplate(rw, "index.gohtml", nil)
		if err != nil {
			log.Error.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}
	}
}
