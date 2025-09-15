package handlers

import (
	"net/http"

	x "github.com/bloxui/blox"
	"github.com/bloxui/starter/internal/service"
	"github.com/bloxui/starter/internal/views"
)

type Users struct{ Svc *service.UserService }

func NewUsers(svc *service.UserService) *Users { return &Users{Svc: svc} }

func (h *Users) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		_ = r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		if name != "" || email != "" {
			_, _ = h.Svc.Create(name, email)
		}
		// Fallthrough to GET rendering
		fallthrough
	default:
		assets := x.NewAssets()
		users, _ := h.Svc.List()
		page := views.UsersPage(users)
		assets.Collect(page)
		doc := views.LayoutWithAssetsProvided("Users", page, assets)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte("<!DOCTYPE html>\n" + x.Render(doc)))
	}
}
