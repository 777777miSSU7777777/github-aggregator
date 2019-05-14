package view

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/token"

	log "github.com/sirupsen/logrus"
)

// MakeIndexHandler returns func with logging which renders index page.
func MakeIndexHandler(logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tkn := token.GetTokenService().GetToken()

		err := GetTemplates().ExecuteTemplate(rw, "index.gohtml", AuthState{Auth: tkn != ""})

		if err != nil {
			logger.Warnln(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		} else {
			logger.Infoln("Index page rendered")
		}
	}
}
