package service

import "apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
