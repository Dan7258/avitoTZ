package models

type User struct {
	TeamMember
	TeamName string `json:"team_name"`
}
