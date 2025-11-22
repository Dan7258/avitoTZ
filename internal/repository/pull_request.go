package repository

import (
	"avito/internal/models"
	"log"
	"strings"
)

func (db *PostgresDB) CreatePullRequest(pullRequest *models.PullRequestShortWith[[]string]) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	pullRequest.Status = models.Open
	var text string
	err = tx.QueryRow(
		`insert into pull_request (pull_request_id, pull_request_name, author_id, status, assigned_reviews) 
				(select $1, $2, $3, $4, array_agg(id) from users where team_name = (select team_name from users where id = $5) and is_active = true and id != $6 limit 2) 
				returning pull_request_id, pull_request_name, author_id, status, assigned_reviews`,
		pullRequest.PullRequestID, pullRequest.PullRequestName, pullRequest.AuthorId, pullRequest.Status, pullRequest.AuthorId, pullRequest.AuthorId).
		Scan(&pullRequest.PullRequestID, &pullRequest.PullRequestName, &pullRequest.AuthorId, &pullRequest.Status, &text)
	if err != nil {
		log.Println(err)
		return err
	}
	prepText := string([]rune(text)[1 : len([]rune(text))-1])
	pullRequest.Extra = append(pullRequest.Extra, strings.Split(prepText, ",")...)
	_, err = tx.Exec("update users set is_active = false where id = any($1)", pullRequest.Extra)
	if err != nil {
		log.Println(err)
		return err
	}
	return tx.Commit()
}

func (db *PostgresDB) SetPullRequestMerged(pullRequest *models.PullRequest) error {
	return nil
}

func (db *PostgresDB) ReassignPullRequest(pullRequestID, oldReviewerId string) (*models.PullRequest, error) {
	return nil, nil
}
