package handler

import (
	"avito/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreatePullRequest(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		PullRequestID   string `json:"pull_request_id"`
		PullRequestName string `json:"pull_request_name"`
		AuthorId        string `json:"author_id"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		return
	}
	pullRequest := &models.PullRequest{
		PullRequestID:   data.PullRequestID,
		PullRequestName: data.PullRequestName,
		AuthorId:        data.AuthorId,
		Status:          models.Open,
	}
	err = h.db.CreatePullRequest(pullRequest)
	if err != nil {
		if err == models.UserNotFoundError {
			jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		} else {
			jsonError(w, PrExists, "PR id already exists", http.StatusConflict)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pullRequest)

}

func (h *Handler) SetPullRequestMerged(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		PullRequestID string `json:"pull_request_id"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		jsonError(w, NotValid, "Not valid data", http.StatusBadRequest)
		return
	}
	pullRequest := &models.PullRequest{
		PullRequestID: data.PullRequestID,
	}
	err = h.db.SetPullRequestMerged(pullRequest)
	if err != nil {
		jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pullRequest)
}
