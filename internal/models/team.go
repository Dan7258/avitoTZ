package models

type Team struct {
	TeamName string `gorm:"unique" json:"team_name"`
	Members  []User `gorm:"foreignKey:TeamName;references:TeamName" json:"members"`
}
