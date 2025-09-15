package app

import (
	"fmt"
	stdhttp "net/http"

	"github.com/bloxui/starter/internal/css"
	"github.com/bloxui/starter/internal/handlers"
	"github.com/bloxui/starter/internal/httpx"
	"github.com/bloxui/starter/internal/repo"
	"github.com/bloxui/starter/internal/service"
)

type App struct {
	Mux stdhttp.Handler
}

func NewApp() *App {
	// Repos / Services
	userRepo := repo.NewInMemoryUserRepo()
	userSvc := service.NewUserService(userRepo)

	// Handlers
	home := handlers.NewHome()
	users := handlers.NewUsers(userSvc)

	// Router
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/healthz", func(w stdhttp.ResponseWriter, r *stdhttp.Request) { w.WriteHeader(stdhttp.StatusOK) })
	mux.HandleFunc("/assets/styles.css", cssHandler)
	mux.HandleFunc("/robots.txt", robotsHandler)
	mux.HandleFunc("/", home.Index)
	mux.HandleFunc("/users", users.Index)

	// Middleware chain
	h := httpx.Chain(mux, httpx.Recoverer, httpx.Gzip, httpx.Logger)
	return &App{Mux: h}
}

func cssHandler(w stdhttp.ResponseWriter, _ *stdhttp.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Header().Set("Cache-Control", "public, max-age=31536000") // 1 year
	w.Header().Set("ETag", "\"css-v1\"")
	_, _ = fmt.Fprint(w, css.TailwindCSS)
}

func robotsHandler(w stdhttp.ResponseWriter, _ *stdhttp.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = fmt.Fprint(w, "User-agent: *\n")
}
