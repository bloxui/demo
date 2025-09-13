package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bloxui/demo/internal"
)

func main() {
	// Register route handlers
	http.HandleFunc("/", internal.HomeHandler)
	http.HandleFunc("/features", internal.FeaturesHandler)
	http.HandleFunc("/docs", internal.DocsHandler)
	http.HandleFunc("/contact", internal.ContactHandler)
	http.HandleFunc("/modal", internal.ModalHandler)
	http.HandleFunc("/tabs", internal.TabsHandler)
	http.HandleFunc("/assets/styles.css", internal.CssHandler)

	fmt.Println("Demo running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
