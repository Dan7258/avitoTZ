package routes

import (
	"avito/internal/handler"
	"net/http"
)

func SetRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /team/add", h.AddTeam)
	mux.HandleFunc("GET /team/get", h.GetTeam)

	mux.HandleFunc("POST /users/setIsActive", h.SetIsActive)
	mux.HandleFunc("GET /users/getReview", h.GetReview)

	//mux.HandleFunc("POST /pullRequest/create", h.CreatePullRequest)
	//mux.HandleFunc("POST /pullRequest/merge", h.SetPullRequestMerged)
	//mux.HandleFunc("POST /pullRequest/reassign", h.ReassignPullRequest)
	return mux
}
