package routes

import (
	"avito/internal/handler"
	"net/http"
)

func SetRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /team/add", h.AddTeam)
	mux.HandleFunc("GET /team/get", h.GetTeam)
	return mux
}
