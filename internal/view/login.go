package view

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/token"

	log "github.com/sirupsen/logrus"
)

// MakeLoginRenderHandler returns func with logging which renders login page.
func MakeLoginRenderHandler(logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tkn := token.GetTokenService().GetToken()

		if tkn != "" {
			http.Redirect(rw, req, "/", 301)
			return
		}

		err := GetTemplates().ExecuteTemplate(rw, "login.gohtml", AuthState{Auth: tkn != ""})

		if err != nil {
			logger.Warnln(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		} else {
			logger.Infoln("Login page rendered")
		}
	}
}
