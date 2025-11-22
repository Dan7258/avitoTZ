package repository

import "avito/internal/models"

func (db *PostgresDB) SetUserIsActive(user *models.User) error {
	return nil
}

func (db *PostgresDB) GetUsersReviews(userID string) ([]models.PullRequest, error) {
	return nil, nil
}
