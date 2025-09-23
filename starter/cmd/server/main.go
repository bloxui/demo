package main

import (
	"fmt"
	"log"
	"net/http"

	starter "github.com/plainkit/starter/internal"
)

func main() {
	mux := starter.Routes()

	addr := ":8080"
	fmt.Println("🚀 Plain Starter Demo Server starting on :8080")
	fmt.Println("🔗 Open http://localhost:8080 to view the demo")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
