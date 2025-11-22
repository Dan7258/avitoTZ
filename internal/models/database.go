package models

type Model interface {
	ConnectToDatabase() error

	SetUserIsActive(user *User) error
	GetUsersReviews(userID string) ([]PullRequestShort, error)

	AddTeam(team *Team) error
	GetTeamByName(teamName string) (*Team, error)

	CreatePullRequest(pullRequest *PullRequestShortWith[Array]) error
	SetPullRequestMerged(pullRequest *PullRequestShortWith[ArrayAndMergedAt]) error
	ReassignPullRequest(pullRequestID, oldReviewerId string) (*PullRequestShortWith[ArrayAndReplaceBy], error)
}
