package main

import (
	"fmt"
	"log"
	"net/http"

	"modern_todo_plain/internal/app"
)

func main() {
	application := app.New()

	addr := ":8080"
	fmt.Println("🚀 Modern Todo App Demo Server starting on :8080")
	fmt.Println("🔗 Open http://localhost:8080 to view the demo")

	if err := http.ListenAndServe(addr, application.Mux); err != nil {
		log.Fatal(err)
	}
}
