package models

type PullRequestShort struct {
	PullRequestID   string            `json:"pull_request_id"`
	PullRequestName string            `json:"pull_request_name"`
	AuthorId        string            `json:"author_id"`
	Status          PullRequestStatus `json:"status"`
}

type PullRequestShortWith[T any] struct {
	PullRequestID   string            `json:"pull_request_id"`
	PullRequestName string            `json:"pull_request_name"`
	AuthorId        string            `json:"author_id"`
	Status          PullRequestStatus `json:"status"`
	Extra           T                 `json:",inline"`
}
