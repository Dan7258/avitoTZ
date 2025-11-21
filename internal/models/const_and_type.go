package models

type PullRequestStatus string

const (
	Open   PullRequestStatus = "OPEN"
	Merged PullRequestStatus = "MERGED"
)
