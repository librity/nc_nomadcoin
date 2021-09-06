package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	value := params[key]

	return value
}
