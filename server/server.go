package server

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func Start() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
