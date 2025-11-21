package repository

import "avito/internal/models"

func (db *PostgresDB) AddTeam(team *models.Team) error {
	return db.Conn.Create(team).Error
}
