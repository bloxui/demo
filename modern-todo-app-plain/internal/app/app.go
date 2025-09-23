package app

import (
	"fmt"
	"net/http"

	"modern_todo_plain/internal/css"
	"modern_todo_plain/internal/handlers"
	"modern_todo_plain/internal/store"
)

type App struct {
	Mux http.Handler
}

func New() *App {
	todoStore := store.New()
	todos := handlers.NewTodoHandler(todoStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/assets/styles.css", cssHandler)
	mux.HandleFunc("/", todos.Index)
	mux.HandleFunc("/todos/create", todos.Create)
	mux.HandleFunc("/todos/update", todos.Update)
	mux.HandleFunc("/todos/toggle", todos.Toggle)
	mux.HandleFunc("/todos/delete", todos.Delete)

	return &App{Mux: mux}
}

func cssHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	_, _ = fmt.Fprint(w, css.TailwindCSS)
}
