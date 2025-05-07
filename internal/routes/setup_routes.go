package routes

import (
	"net/http"

	"github.com/Nyantise/image_provider/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
}
