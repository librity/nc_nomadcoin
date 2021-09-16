package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetParam gets the route parameter of index "key".
func GetParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	value := params[key]

	return value
}

// GetQuery gets the query parameter of index "key".
func GetQuery(r *http.Request, key string) string {
	queryParam := r.URL.Query().Get(key)

	return queryParam
}
