package middleware

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

// WithAuthCheck decorator for api, which checks auth state before handle.
func WithAuthCheck(apiHandler httpHandlerFunc) httpHandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if session.GetSessionService().HasActiveSession() {
			apiHandler(rw, req)
		} else {
			_, err := rw.Write([]byte("You are not authorized"))

			if err != nil {
				log.Error.Println(err)
			}

			rw.WriteHeader(401)
			log.Warning.Println("Anauthorized request to api")
		}
		log.Info.Println(session.GetSessionService().HasActiveSession())
	}
}
