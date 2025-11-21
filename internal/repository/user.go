package repository

import "avito/internal/models"

func (db *PostgresDB) SetUserActive(userID string) error {
	return db.Conn.Model(models.User{}).Where("id = ?", userID).Update("is_active", true).Error
}
