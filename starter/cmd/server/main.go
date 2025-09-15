package main

import (
	"log"
	"net/http"

	starter "github.com/plainkit/starter/internal"
)

func main() {
	mux := starter.Routes()

	addr := ":8080"
	log.Printf("bloxui starter listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
