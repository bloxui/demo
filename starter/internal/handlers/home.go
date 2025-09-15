package handlers

import (
	"net/http"

	x "github.com/bloxui/blox"
	"github.com/bloxui/starter/internal/views"
)

type Home struct{}

func NewHome() *Home { return &Home{} }

func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	page := views.HomePage()
	doc := views.Layout("BloxUI Starter", page)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte("<!DOCTYPE html>\n" + x.Render(doc)))
}
