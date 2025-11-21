package models

type CodeError string
type PullRequestStatus string

const (
	TeamExists  CodeError = "TEAM_EXISTS"
	PrExists    CodeError = "PR_EXISTS"
	PrMerged    CodeError = "PR_MERGED"
	NotAssigned CodeError = "NOT_ASSIGNED"
	NoCandidate CodeError = "NO_CANDIDATE"
	NotFound    CodeError = "NOT_FOUND"

	Open   PullRequestStatus = "OPEN"
	Merged PullRequestStatus = "MERGED"
)
