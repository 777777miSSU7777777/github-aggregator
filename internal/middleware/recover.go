package middleware

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
)

// WithRecover decorator for api, which recovers from panic.
func WithRecover(apiHandler httpHandlerFunc) httpHandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Warning.Printf("Recovered from %s", r)
			}
		}()

		apiHandler(rw, req)
	}
}
