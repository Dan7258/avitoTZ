package repository

import (
	"avito/internal/models"
)

func (db *PostgresDB) CreatePullRequest(pullRequest *models.PullRequest) error {
	author := new(models.User)
	usersID := make([]string, 0)
	res := db.Conn.Where("user_id = ?", pullRequest.AuthorId).First(author)
	if res.Error != nil {
		return models.UserNotFoundError
	}
	db.Conn.Model(author).
		Select("user_id").
		Where("team_name = ? AND is_active = ?", author.TeamName, true).
		Not("user_id = ?", pullRequest.AuthorId).
		Order("RANDOM()").
		Limit(2).
		Find(&usersID)

	pullRequest.AssignedReviews = append(pullRequest.AssignedReviews, usersID...)

	if len(usersID) > 0 {
		result := db.Conn.Model(&models.User{}).
			Where("user_id IN ?", usersID).
			Update("is_active", false)

		if result.Error != nil {
			return result.Error
		}
	}
	return db.Conn.Create(pullRequest).Error
}

func (db *PostgresDB) SetPullRequestMerged(pullRequest *models.PullRequest) error {
	res := db.Conn.
		Model(pullRequest).
		Where("pull_request_id = ?", pullRequest.PullRequestID).
		Update("status", models.Merged).
		Scan(pullRequest)
	if res.Error == nil {
		if res.RowsAffected == 0 {
			return models.ZeroUpdatedRowsError
		}
	}
	return res.Error
}
