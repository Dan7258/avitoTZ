package repository

import (
	"avito/internal/models"
)

func (db *PostgresDB) SetUserIsActive(user *models.User) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.QueryRow(
		"update users set is_active = $1 where id = $2 returning id, username, is_active, team_name",
		user.IsActive, user.UserID).Scan(&user.UserID, &user.Username, &user.IsActive, &user.TeamName)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (db *PostgresDB) GetUsersReviews(userID string) ([]models.PullRequestShort, error) {
	pr := make([]models.PullRequestShort, 0)
	rows, err := db.Conn.Query("select pull_request_id, pull_request_name, author_id, status from pull_request where $1 = any(assigned_reviews)", userID)
	if err != nil {
		return pr, err
	}
	defer rows.Close()
	for rows.Next() {
		pullRequest := models.PullRequestShort{}
		err = rows.Scan(&pullRequest.PullRequestID, &pullRequest.PullRequestName, &pullRequest.AuthorId, &pullRequest.Status)
		if err != nil {
			return pr, err
		}
		pr = append(pr, pullRequest)
	}
	return pr, nil
}
