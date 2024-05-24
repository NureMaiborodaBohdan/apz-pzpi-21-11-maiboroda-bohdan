package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AdminMysql struct {
	db *sqlx.DB
}

func NewAdminMysql(db *sqlx.DB) *AdminMysql {
	return &AdminMysql{db: db}
}

func (r *AdminMysql) GetUserByID(userID int) (AlcoSafe.User, error) {
	var user AlcoSafe.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role FROM User WHERE UserID=?")
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}
