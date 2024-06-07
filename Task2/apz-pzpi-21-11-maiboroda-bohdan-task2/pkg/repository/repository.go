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
	Create(company AlcoSafe.Company) (int, error)
	GetAll() ([]AlcoSafe.Company, error)
	GetByID(companyID int) (AlcoSafe.Company, error)
	Delete(companyID int) error
}

type Location interface {
	Create(location AlcoSafe.Location) (int, error)
	GetByID(locationID int) (AlcoSafe.Location, error)
	Delete(locationID int) error
}

type TestResult interface {
}

type AccessControl interface {
}

type Notification interface {
}
type Admin interface {
	GetUserByID(userID int) (AlcoSafe.User, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GetAllUsers() ([]AlcoSafe.User, error)
	UpdateUser(UserID int, input AlcoSafe.UpdateUserInput) error
	Delete(UserID int) error
}
type Repository struct {
	Authorization
	Company
	Location
	TestResult
	AccessControl
	Notification
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Admin:         NewAdminMysql(db),
		Company:       NewCompanyMysql(db),
		Location:      NewLocationMysql(db),
	}
}
