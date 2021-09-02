package rest

import (
	"net/http"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	middleware := func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(middleware)
}
