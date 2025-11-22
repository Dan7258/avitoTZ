package models

type PullRequest struct {
	PullRequestShort
	AssignedReviews []string `json:"assigned_reviews"`
	CreatedAt       string   `json:"created_at"`
	MergedAt        string   `json:"merged_at"`
}
