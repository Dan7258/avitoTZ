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
	return tx.Commit()
}

func (db *PostgresDB) ReassignPullRequest(pullRequestID, oldReviewerId string) (*models.PullRequestShortWith[models.ArrayAndReplaceBy], error) {
	tx, err := db.Conn.Begin()
	pullRequest := &models.PullRequestShortWith[models.ArrayAndReplaceBy]{}
	if err != nil {
		return pullRequest, err
	}
	id := ""
	err = tx.QueryRow("select pull_request_id from pull_request where pull_request_id = $1", pullRequestID).Scan(&id)
	if err != nil {

		return pullRequest, models.NotFoundError
	}
	err = tx.QueryRow("select id from users where id = $1", oldReviewerId).Scan(&id)
	if err != nil {

		return pullRequest, models.NotFoundError
	}
	var text string
	err = tx.QueryRow("select status, author_id, assigned_reviews, pull_request_id, pull_request_name from pull_request where pull_request_id = $1", pullRequestID).
		Scan(&pullRequest.Status, &pullRequest.AuthorId, &text, &pullRequest.PullRequestID, &pullRequest.PullRequestName)
	if err != nil {

		return pullRequest, models.NotFoundError
	} else if pullRequest.Status != models.Open {
		return pullRequest, models.NotChangedError
	}
	pullRequest.Extra.AssignedReviews = append(pullRequest.Extra.AssignedReviews, parseTextToArray(text)...)
	id = ""
	_ = tx.QueryRow(
		`select id from users where team_name = (select team_name from users where id = $1) and is_active = true and id != $2 and id != $3 limit 1`,
		oldReviewerId, pullRequest.AuthorId, oldReviewerId).
		Scan(&id)

	var arr []string
	if id != "" {
		arr = append(arr, id)
	}
	for _, a := range pullRequest.Extra.AssignedReviews {
		if a != oldReviewerId {
			arr = append(arr, a)
		}
	}
	pullRequest.Extra.AssignedReviews = arr
	pullRequest.Extra.ReplaceBy = id

	_, err = tx.Exec(
		"update users set is_active = $1 where id = $2",
		true, oldReviewerId)
	if err != nil {
		log.Println(err)
		return pullRequest, err
	}
	defer tx.Rollback()
	return pullRequest, tx.Commit()
}

func parseTextToArray(text string) []string {
	if text == "" {
		return make([]string, 0)
	}
	prepText := string([]rune(text)[1 : len([]rune(text))-1])
	return strings.Split(prepText, ",")
}
