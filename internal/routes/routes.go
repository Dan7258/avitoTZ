package routes

import (
	"avito/internal/handler"
	"net/http"
)

func SetRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
