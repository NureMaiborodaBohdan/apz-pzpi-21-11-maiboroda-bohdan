package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type AdministratorService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdministratorService {
	return &AdministratorService{repo: repo}
}

func (s *AdministratorService) GetUserByID(userID int) (AlcoSafe.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *AdministratorService) CreateUser(user AlcoSafe.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AdministratorService) UpdateUser(UserID int, input AlcoSafe.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateUser(UserID, input)
}
func (s *AdministratorService) GetAllUsers() ([]AlcoSafe.User, error) {
	return s.repo.GetAllUsers()
}
func (s *AdministratorService) Delete(UserID int) error {
	return s.repo.Delete(UserID)
}
