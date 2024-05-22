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

type Service struct {
	Authorization
	Company
	Location
	Thresholds
	TestResult
	AccessControl
	Notification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
