package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	value := params[key]

	return value
}

func GetQuery(r *http.Request, key string) string {
	queryParam := r.URL.Query().Get(key)

	return queryParam
}
