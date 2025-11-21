package models

type User struct {
	TeamMember
	TeamName string `gorm:"unique" json:"team_name"`
}
