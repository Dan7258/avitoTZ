package models

type TeamMember struct {
	UserID   string `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}
