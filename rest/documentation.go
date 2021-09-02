package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/librity/nc_nomadcoin/utils"
)

type URL string

// TextMarshaler interface: https://pkg.go.dev/encoding#TextMarshaler
func (u URL) MarshalText() ([]byte, error) {
	fullURL := baseURL + string(u)
	return []byte(fullURL), nil
}

type EndpointDescription struct {
	URL     URL    `json:"url"`
	Method  string `json:"method"`
	Details string `json:"details"`
	Payload string `json:"payload,omitempty"`
}

var data = []EndpointDescription{
	{
		URL:     URL("/"),
		Method:  "GET",
		Details: "Read API documentation.",
	},

	{
		URL:     URL("/blocks"),
		Method:  "GET",
		Details: "Read all blocks.",
	},
	{
		URL:     URL("/blocks"),
		Method:  "POST",
		Details: "Create a block.",
		Payload: "data:string",
	},
	{
		URL:     URL("/blocks/:id"),
		Method:  "GET",
		Details: "Read a block.",
	},
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func deprecated(rw http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(data)
	utils.HandleError(err)

	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", bytes)
}
