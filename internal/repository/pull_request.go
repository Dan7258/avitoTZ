package repository

import (
	"avito/internal/models"
)

func (db *PostgresDB) CreatePullRequest(pullRequest *models.PullRequest) error {
	author := new(models.User)
	usersID := make([]string, 0)
	res := db.Conn.Where("user_id = ?", pullRequest.AuthorId).First(author)
	if res.Error != nil {
		return models.NotFoundError
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

func (db *PostgresDB) ReassignPullRequest(pullRequestID, oldReviewerId string) (*models.PullRequest, error) {
	pullRequest := new(models.PullRequest)
	usersID := make([]string, 0)
	user := new(models.User)
	teamName := ""
	res := db.Conn.Model(pullRequest).Where("pull_request_id = ?", pullRequestID).First(pullRequest)
	if res.Error != nil {
		return pullRequest, models.NotFoundError
	}
	if pullRequest.Status == models.Merged {
		return pullRequest, models.NotChangedError
	}
	res.Select("users.team_name").
		Joins("JOIN users ON pull_requests.author_id = users.user_id").
		Find(&teamName)
	if teamName == "" {
		return pullRequest, models.NotFoundError
	}
	var err error
	db.Conn.Model(user).
		Select("user_id").
		Where("team_name = ? AND is_active = ?", teamName, true).
		Where("user_id != ?", pullRequest.AuthorId).
		Where("user_id NOT IN ?", pullRequest.AssignedReviews).
		Order("RANDOM()").
		Limit(1).
		Find(&usersID)
	arr := make([]string, 0)
	for _, val := range pullRequest.AssignedReviews {
		if val != oldReviewerId {
			arr = append(arr, val)
		} else {
			user.UserID = val
			user.IsActive = true
			err = db.SetUserIsActive(user)
		}
	}
	if err != nil {
		return pullRequest, err
	}
	arr = append(arr, usersID...)
	pullRequest.AssignedReviews = pullRequest.AssignedReviews[:0]
	pullRequest.AssignedReviews = append(pullRequest.AssignedReviews, arr...)
	if len(pullRequest.AssignedReviews) > 0 {
		err = db.Conn.Model(&models.User{}).
			Where("user_id IN ?", pullRequest.AssignedReviews).
			Update("is_active", false).Error
		if err != nil {
			return pullRequest, err
		}
	}
	err = db.Conn.Save(pullRequest).Error
	return pullRequest, err
}
