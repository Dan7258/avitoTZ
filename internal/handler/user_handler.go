package handler

import (
	"avito/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) SetIsActive(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		UserID   string `json:"user_id"`
		IsActive bool   `json:"is_active"`
	}{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		return
	}
	user := new(models.User)
	user.UserID = data.UserID
	user.IsActive = data.IsActive
	err = h.db.SetUserIsActive(user)
	if err != nil {
		jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			UserID       string               `json:"user_id"`
			PullRequests []models.PullRequest `json:"pull_requests"`
		}{
			UserID:       userID,
			PullRequests: []models.PullRequest{},
		})
		return
	}
	reviews, _ := h.db.GetUsersReviews(userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		UserID       string               `json:"user_id"`
		PullRequests []models.PullRequest `json:"pull_requests"`
	}{
		UserID:       userID,
		PullRequests: reviews,
	})
}
