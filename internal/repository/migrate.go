package repository

import "avito/internal/models"

func (db *PostgresDB) Migrate() error {
	return db.Conn.AutoMigrate(&models.Team{}, &models.User{}, &models.PullRequest{})
}
