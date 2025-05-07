package routes

import (
	"net/http"

	"github.com/Nyantise/image_provider/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandlers())
}
