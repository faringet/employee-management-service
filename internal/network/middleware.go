package network

import "net/http"

// ApplyMiddlewares return handler with performed middlewares in the same order.
func ApplyMiddlewares(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
