package repository

import "avito/internal/models"

func (db *PostgresDB) SetUserIsActive(user *models.User) error {
	res := db.Conn.
		Model(user).
		Where("user_id = ?", user.UserID).
		Update("is_active", user.IsActive).
		Scan(user)
	if res.Error == nil {
		if res.RowsAffected == 0 {
			return ZeroUpdatedRowsError
		}
	}
	return res.Error
}
