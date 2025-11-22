package repository

import (
	"avito/internal/models"
	"log"
	"strings"
)

func (db *PostgresDB) CreatePullRequest(pullRequest *models.PullRequestShortWith[models.Array]) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	pullRequest.Status = models.Open
	id := ""
	err = tx.QueryRow("select id from users where id = $1", pullRequest.AuthorId).Scan(&id)
	if err != nil {
		return models.NotFoundError
	}
	var text string
	err = tx.QueryRow(
		`insert into pull_request (pull_request_id, pull_request_name, author_id, status, assigned_reviews) 
				(select $1, $2, $3, $4, array(select id from users where team_name = (select team_name from users where id = $5) and is_active = true and id != $6 limit 2))
				returning pull_request_id, pull_request_name, author_id, status, assigned_reviews`,
		pullRequest.PullRequestID, pullRequest.PullRequestName, pullRequest.AuthorId, pullRequest.Status, pullRequest.AuthorId, pullRequest.AuthorId).
		Scan(&pullRequest.PullRequestID, &pullRequest.PullRequestName, &pullRequest.AuthorId, &pullRequest.Status, &text)
	if err != nil {
		log.Println(err)
		return err
	}
	pullRequest.Extra.AssignedReviews = append(pullRequest.Extra.AssignedReviews, parseTextToArray(text)...)
	_, err = tx.Exec("update users set is_active = false where id = any($1)", pullRequest.Extra)
	if err != nil {
		log.Println(err)
		return err
	}
	return tx.Commit()
}

func (db *PostgresDB) SetPullRequestMerged(pullRequest *models.PullRequestShortWith[models.ArrayAndMergedAt]) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	id := ""
	err = tx.QueryRow("select pull_request_id from pull_request where pull_request_id = $1", pullRequest.PullRequestID).Scan(&id)
	if err != nil {
		log.Println(err)
		return models.NotFoundError
	}
	var text string
	pullRequest.Status = models.Merged
	err = tx.QueryRow(
		`update pull_request set status = $1, merged_at = now() where pull_request_id = $2
               returning pull_request_id, pull_request_name, author_id, status, assigned_reviews, merged_at`,
		pullRequest.Status, pullRequest.PullRequestID).
		Scan(&pullRequest.PullRequestID, &pullRequest.PullRequestName, &pullRequest.AuthorId, &pullRequest.Status, &text, &pullRequest.Extra.MergedAt)
	if err != nil {
		log.Println(err)
		return err
	}
	pullRequest.Extra.AssignedReviews = append(pullRequest.Extra.AssignedReviews, parseTextToArray(text)...)
	return nil
}

func (db *PostgresDB) ReassignPullRequest(pullRequestID, oldReviewerId string) (*models.PullRequest, error) {
	return nil, nil
}

func parseTextToArray(text string) []string {
	if text == "" {
		return make([]string, 0)
	}
	prepText := string([]rune(text)[1 : len([]rune(text))-1])
	return strings.Split(prepText, ",")
}
