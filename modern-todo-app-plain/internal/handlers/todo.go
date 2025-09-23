package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"modern_todo_plain/internal/store"
	"modern_todo_plain/internal/views"
)

type TodoHandler struct {
	store *store.Store
}

func NewTodoHandler(store *store.Store) *TodoHandler {
	return &TodoHandler{store: store}
}

func (h *TodoHandler) Index(w http.ResponseWriter, r *http.Request) {
	filter := parseFilter(r.URL.Query().Get("filter"))
	data := h.pageData(filter)

	if isHX(r) && r.URL.Query().Get("partial") == "app" {
		writeHTML(w, views.RenderAppShell(data))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprint(w, "<!DOCTYPE html>\n")
	_, _ = fmt.Fprint(w, views.RenderFullPage(data))
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form submission", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}
	description := strings.TrimSpace(r.FormValue("description"))
	priority := parsePriority(r.FormValue("priority"))
	filter := parseFilter(r.FormValue("filter"))

	h.store.Add(title, description, priority)
	h.respondWithApp(w, r, filter)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form submission", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	if title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	description := strings.TrimSpace(r.FormValue("description"))
	priority := parsePriority(r.FormValue("priority"))
	filter := parseFilter(r.FormValue("filter"))

	if _, err := h.store.Update(id, title, description, priority); err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	h.respondWithApp(w, r, filter)
}

func (h *TodoHandler) Toggle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	filter := parseFilter(r.FormValue("filter"))

	if _, err := h.store.Toggle(id); err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	h.respondWithApp(w, r, filter)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	filter := parseFilter(r.FormValue("filter"))

	if err := h.store.Delete(id); err != nil {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}

	h.respondWithApp(w, r, filter)
}

func (h *TodoHandler) respondWithApp(w http.ResponseWriter, r *http.Request, filter store.Filter) {
	data := h.pageData(filter)
	if isHX(r) {
		writeHTML(w, views.RenderAppShell(data))
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/?filter=%s", filter), http.StatusSeeOther)
}

func (h *TodoHandler) pageData(filter store.Filter) views.PageData {
	return views.PageData{
		Todos:  h.store.List(filter),
		Filter: filter,
		Stats:  h.store.Stats(),
	}
}

func parseFilter(raw string) store.Filter {
	switch strings.ToLower(raw) {
	case string(store.FilterActive):
		return store.FilterActive
	case string(store.FilterCompleted):
		return store.FilterCompleted
	default:
		return store.FilterAll
	}
}

func parsePriority(raw string) store.Priority {
	switch strings.ToLower(raw) {
	case string(store.PriorityLow):
		return store.PriorityLow
	case string(store.PriorityHigh):
		return store.PriorityHigh
	default:
		return store.PriorityMedium
	}
}

func isHX(r *http.Request) bool {
	return strings.EqualFold(r.Header.Get("HX-Request"), "true")
}

func writeHTML(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprint(w, html)
}
