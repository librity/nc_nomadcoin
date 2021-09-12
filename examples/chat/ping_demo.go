package main

import (
	"fmt"
	"net/http"

	"github.com/librity/nc_nomadcoin/utils"
)

func pingDemo(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	utils.PanicError(err)

	for {
		fmt.Println("Awaiting message...")
		_, payload, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\"\n", payload)
		fmt.Println("---")
	}
}
