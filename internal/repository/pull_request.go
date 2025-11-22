package repository

import (
	"avito/internal/models"
)

func (db *PostgresDB) CreatePullRequest(pullRequest *models.PullRequest) error {
	return nil
}

func (db *PostgresDB) SetPullRequestMerged(pullRequest *models.PullRequest) error {
	return nil
}

func (db *PostgresDB) ReassignPullRequest(pullRequestID, oldReviewerId string) (*models.PullRequest, error) {
	return nil, nil
}
