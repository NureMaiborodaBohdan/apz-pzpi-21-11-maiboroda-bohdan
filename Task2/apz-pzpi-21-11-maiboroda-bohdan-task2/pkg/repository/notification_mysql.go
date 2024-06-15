package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type NotificationMysql struct {
	db *sqlx.DB
}

func NewNotificationMysql(db *sqlx.DB) *NotificationMysql {
	return &NotificationMysql{db: db}
}
func (r *NotificationMysql) GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error) {
	var userNotification []AlcoSafe.Notification
	query := "SELECT * FROM Notification WHERE UserID=?"
	err := r.db.Select(&userNotification, query, userID)
	return userNotification, err
}
