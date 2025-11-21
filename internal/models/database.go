package models

import (
	"gorm.io/gorm"
)

type Model interface {
	GetConn() *gorm.DB
	ConnectToDatabase() error
	Migrate() error

	SetUserIsActive(user *User) error
	GetUsersReviews(userID string) ([]PullRequest, error)

	AddTeam(team *Team) error
	GetTeamByName(teamName string) (*Team, error)

	CreatePullRequest(pullRequest *PullRequest) error
	SetPullRequestMerged(pullRequest *PullRequest) error
	ReassignPullRequest(pullRequestID, oldReviewerId string) (*PullRequest, error)
}
