package middleware

import (
	"net/http"
)

type httpHandlerFunc = func(http.ResponseWriter, *http.Request)
type middlewareFunc = func(httpHandlerFunc) httpHandlerFunc

// ChainMiddleware decorates http handler with auth check and recover middleware.
func ChainMiddleware(decorated httpHandlerFunc) httpHandlerFunc {
	return chainMiddleWare(decorated, WithRecover, WithAuthCheck)
}

func chainMiddleWare(decor httpHandlerFunc, decorators ...middlewareFunc) httpHandlerFunc {
	decorated := decor
	for _, decorator := range decorators {
		decorated = decorator(decorated)
	}
	return decorated
}
