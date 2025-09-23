package main

import (
	"log"
	"net/http"

	"modern_todo_plain/internal/app"
)

func main() {
	application := app.New()

	addr := ":8080"
	log.Printf("Modern Todo App listening on %s", addr)

	if err := http.ListenAndServe(addr, application.Mux); err != nil {
		log.Fatal(err)
	}
}
