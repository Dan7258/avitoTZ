package repository

import "avito/internal/models"

func (db *PostgresDB) AddTeam(team *models.Team) error {
	return db.Conn.Create(team).Error
}

func (db *PostgresDB) GetTeamByName(teamName string) (*models.Team, error) {
	team := new(models.Team)
	err := db.Conn.
		Preload("Members").
		Where("team_name = ?", teamName).First(team).Error
	return team, err
}
