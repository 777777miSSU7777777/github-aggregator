// Package index implements function for index page render.
package index

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/time/timeutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

// Render renders index page.
func Render(rw http.ResponseWriter, req *http.Request) {
	tkn := token.GetTokenService().GetToken()

	err := view.GetTemplates().ExecuteTemplate(rw, "index.gohtml", view.AuthState{Auth: tkn != ""})

	if err != nil {
		view.Logger().Log(
			"method", "index.Render",
			"time", timeutil.GetCurrentTime(),
			"err", err,
		)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		view.Logger().Log(
			"method", "index.Render",
			"time", timeutil.GetCurrentTime(),
			"info", "Index page rendered",
		)
	}
}
