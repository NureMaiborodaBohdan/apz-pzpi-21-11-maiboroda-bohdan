package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type Authorization interface {
	CreateAdmin(user AlcoSafe.User) (int, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Delete(UserID int) error
	UpdateUser(UserID int, input AlcoSafe.UpdateUserInput) error
}

type Service struct {
	Authorization
	Company
	Location
	TestResult
	AccessControl
	Notification
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admin:         NewAdminService(repos.Admin),
		Company:       NewCompanyService(repos.Company),
		Location:      NewLocationService(repos.Location),
	}
}
