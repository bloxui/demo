package internal

import (
	"net/http"

	"github.com/bloxui/starter/internal/app"
)

// Routes exposes the application handler for external use (e.g., cmd/server).
func Routes() http.Handler {
	return app.NewApp().Mux
}
