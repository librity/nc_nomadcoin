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
		URL:     url("/blokchain"),
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
		Details: "Mine current block.",
	},
	{
		URL:     url("/blocks/{hash}"),
		Method:  "GET",
		Details: "Read a block.",
	},

	{
		URL:     url("/wallet/server"),
		Method:  "GET",
		Details: "Read the balance and transaction outputs of the server's wallet.",
	},
	{
		URL:     url("/wallet/{address}"),
		Method:  "GET",
		Details: "Read the balance and transaction outputs of a wallet.",
	},

	{
		URL:     url("/mempool"),
		Method:  "GET",
		Details: "Get all unconfirmed transactions waiting in the mempool.",
	},

	{
		URL:     url("/transactions"),
		Method:  "POST",
		Details: "Create and add a new transaction to the mempool.",
		Payload: "to:string,amount:int",
	},

	{
		URL:     url("/peers"),
		Method:  "GET",
		Details: "Get all connected peers.",
	},
	{
		URL:     url("/peers"),
		Method:  "POST",
		Details: "Receiver connects to requester as peer though web sockets.",
		Payload: "address:string,port:string",
	},
	{
		URL:     url("/peers/upgrade"),
		Method:  "GET",
		Details: "Upgrade connection to Web Sockets.",
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
