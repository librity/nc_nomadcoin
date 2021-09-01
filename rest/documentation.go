package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/librity/nc_nomadcoin/utils"
)

type EndpointDescription struct {
	Path    string
	Method  string
	Details string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []EndpointDescription{
		{
			Path:    "/",
			Method:  "GET",
			Details: "Browse API documentation.",
		},
	}
	bytes, err := json.Marshal(data)
	utils.HandleError(err)

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", bytes)
}
