package rest

import (
	"fmt"
	"net/http"
)

func jsonContentTypeMw(next http.Handler) http.Handler {
	middleware := func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(middleware)
}

func loggerMw(next http.Handler) http.Handler {
	middleware := func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("🤖", r.Method, r.URL)

		next.ServeHTTP(rw, r)
	}

	return http.HandlerFunc(middleware)
}
