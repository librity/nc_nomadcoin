package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/librity/nc_nomadcoin/utils"
)

type url string

// TextMarshaler interface: https://pkg.go.dev/encoding#TextMarshaler
func (u url) MarshalText() ([]byte, error) {
	fullURL := baseURL + string(u)
	return []byte(fullURL), nil
}

type endpointDescription struct {
	URL     url    `json:"url"`
	Method  string `json:"method"`
	Details string `json:"details"`
	Payload string `json:"payload,omitempty"`
}

var data = []endpointDescription{
	{
		URL:     url("/"),
		Method:  "GET",
		Details: "Read API documentation.",
	},

	{
		URL:     url("/blocks"),
		Method:  "GET",
		Details: "Read all blocks.",
	},
	{
		URL:     url("/blocks"),
		Method:  "POST",
		Details: "Create a block.",
		Payload: "data:string",
	},
	{
		URL:     url("/blocks/{height}"),
		Method:  "GET",
		Details: "Read a block.",
	},
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(data)
}

func deprecated(rw http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(data)
	utils.HandleError(err)

	fmt.Fprintf(rw, "%s", bytes)
}
