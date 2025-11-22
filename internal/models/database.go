package models

type Model interface {
	ConnectToDatabase() error

	SetUserIsActive(user *User) error
	GetUsersReviews(userID string) ([]PullRequest, error)

	AddTeam(team *Team) error
	GetTeamByName(teamName string) (*Team, error)

	CreatePullRequest(pullRequest *PullRequestShortWith[Array]) error
	SetPullRequestMerged(pullRequest *PullRequestShortWith[ArrayAndMergedAt]) error
	ReassignPullRequest(pullRequestID, oldReviewerId string) (*PullRequest, error)
}
