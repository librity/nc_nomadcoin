package rest

import (
	"encoding/json"
	"net/http"
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
		URL:     url("/blokchain/status"),
		Method:  "GET",
		Details: "See the status of the blockchain.",
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
		URL:     url("/blocks/{hash}"),
		Method:  "GET",
		Details: "Read a block.",
	},

	{
		URL:     url("/wallet/{address}"),
		Method:  "GET",
		Details: "Read the balance and transaction outputs of a wallet.",
	},
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(data)
}

/*
func deprecated(rw http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(data)
	utils.HandleError(err)

	fmt.Fprintf(rw, "%s", bytes)
}
*/
