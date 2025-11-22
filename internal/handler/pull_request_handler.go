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
	pullRequest := &models.PullRequestShortWith[models.Array]{}
	pullRequest.PullRequestID = data.PullRequestID
	pullRequest.PullRequestName = data.PullRequestName
	pullRequest.AuthorId = data.AuthorId
	err = h.db.CreatePullRequest(pullRequest)
	if err != nil {
		if err == models.NotFoundError {
			jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		} else {
			jsonError(w, PrExists, "PR id already exists", http.StatusConflict)
		}
		return
	}
	resp := map[string]interface{}{
		"pull_request_id":   pullRequest.PullRequestID,
		"pull_request_name": pullRequest.PullRequestName,
		"author_id":         pullRequest.AuthorId,
		"status":            pullRequest.Status,
		"assigned_reviews":  pullRequest.Extra.AssignedReviews,
	}
	jsonOK(w, resp, "pr", http.StatusCreated)
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
	pullRequest := &models.PullRequestShortWith[models.ArrayAndMergedAt]{}
	pullRequest.PullRequestID = data.PullRequestID
	err = h.db.SetPullRequestMerged(pullRequest)
	if err != nil {
		jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		return
	}
	resp := map[string]interface{}{
		"pull_request_id":   pullRequest.PullRequestID,
		"pull_request_name": pullRequest.PullRequestName,
		"author_id":         pullRequest.AuthorId,
		"status":            pullRequest.Status,
		"assigned_reviews":  pullRequest.Extra.AssignedReviews,
		"merged_at":         pullRequest.Extra.MergedAt,
	}
	jsonOK(w, resp, "pr", http.StatusOK)
}

func (h *Handler) ReassignPullRequest(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		PullRequestID string `json:"pull_request_id"`
		OldReviewerID string `json:"old_reviewer_id"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		jsonError(w, NotValid, "Not valid data", http.StatusBadRequest)
		return
	}
	pullRequest := new(models.PullRequestShortWith[models.ArrayAndReplaceBy])
	pullRequest, err = h.db.ReassignPullRequest(data.PullRequestID, data.OldReviewerID)
	if err != nil {
		if err == models.NotChangedError {
			jsonError(w, PrMerged, "cannot reassign on merged PR", http.StatusConflict)
		} else {
			jsonError(w, NotFound, "resource not found", http.StatusNotFound)
		}
		return
	}
	resp := map[string]interface{}{
		"pull_request_id":   pullRequest.PullRequestID,
		"pull_request_name": pullRequest.PullRequestName,
		"author_id":         pullRequest.AuthorId,
		"status":            pullRequest.Status,
		"assigned_reviews":  pullRequest.Extra.AssignedReviews,
	}
	resp = createResponse(resp, "pr")
	resp["replaced_by"] = pullRequest.Extra.ReplaceBy
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
