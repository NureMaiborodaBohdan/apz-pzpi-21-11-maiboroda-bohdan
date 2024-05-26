package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type LocationService struct {
	repo repository.Location
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) Create(location AlcoSafe.Location) (int, error) {
	return s.repo.Create(location)
}

func (s *LocationService) GetByID(locationID int) (AlcoSafe.Location, error) {
	return s.repo.GetByID(locationID)
}

func (s *LocationService) Delete(locationID int) error {
	return s.repo.Delete(locationID)
}
