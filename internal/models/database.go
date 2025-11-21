package models

import (
	"gorm.io/gorm"
)

type Model interface {
	GetConn() *gorm.DB
	ConnectToDatabase() error
	Migrate() error

	SetUserActive(userID string) error

	AddTeam(team *Team) error
	GetTeamByName(teamName string) (*Team, error)
}
