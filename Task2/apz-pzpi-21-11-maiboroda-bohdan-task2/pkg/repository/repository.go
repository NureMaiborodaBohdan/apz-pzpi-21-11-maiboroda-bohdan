package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateAdmin(user AlcoSafe.User) (int, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GetUser(username, password string) (AlcoSafe.User, error)
}

type Company interface {
}

type Location interface {
}

type Thresholds interface {
}

type TestResult interface {
}

type AccessControl interface {
}

type Notification interface {
}
type Admin interface {
	GetUserByID(userID int) (AlcoSafe.User, error)
}
type Repository struct {
	Authorization
	Company
	Location
	Thresholds
	TestResult
	AccessControl
	Notification
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Admin:         NewAdminMysql(db),
	}
}
