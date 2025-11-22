package repository

import "avito/internal/models"

func (db *PostgresDB) AddTeam(team *models.Team) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("insert into teams (team_name) values ($1)", team.TeamName)
	if err != nil {
		return models.AlreadyExistsError
	}
	for _, member := range team.Members {
		_, err = tx.Exec("insert into users (id, username, is_active, team_name) values ($1, $2, $3, $4)",
			member.UserID, member.Username, member.IsActive, team.TeamName)
		if err != nil {
			return models.AlreadyExistsError
		}
	}
	return tx.Commit()
}

func (db *PostgresDB) GetTeamByName(teamName string) (*models.Team, error) {
	team := new(models.Team)
	rows, err := db.Conn.Query(
		`SELECT t.team_name, u.id, u.username, u.is_active FROM teams t LEFT JOIN users u ON t.team_name = u.team_name WHERE t.team_name = $1`,
		teamName)
	defer rows.Close()
	if err != nil {
		return team, err
	}
	for rows.Next() {
		member := models.TeamMember{}
		err = rows.Scan(&team.TeamName, &member.UserID, &member.Username, &member.IsActive)
		if err != nil {
			return team, err
		}
		team.Members = append(team.Members, member)
	}
	if len(team.Members) == 0 {
		return team, models.NotFoundError
	}
	return team, err
}
