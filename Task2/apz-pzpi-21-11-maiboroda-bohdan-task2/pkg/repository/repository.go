package repository

import "github.com/jmoiron/sqlx"

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

type Repository struct {
	Authorization
	Company
	Location
	Thresholds
	TestResult
	AccessControl
	Notification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
