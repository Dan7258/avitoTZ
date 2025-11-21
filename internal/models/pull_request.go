package models

type PullRequest struct {
	PullRequestID   string            `gorm:"primaryKey" json:"pull_request_id"`
	PullRequestName string            `json:"pull_request_name"`
	AuthorId        string            `json:"author_id"`
	Status          PullRequestStatus `json:"status"`
	AssignedReviews []string          `gorm:"serializer:json" json:"assigned_reviews"`
	CreatedAt       string            `json:"created_at"`
	MergedAt        string            `json:"merged_at"`
}
