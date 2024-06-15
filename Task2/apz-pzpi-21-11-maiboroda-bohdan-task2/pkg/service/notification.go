package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}
func (s *NotificationService) GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error) {
	return s.repo.GetAllUserNotification(userID)
}
