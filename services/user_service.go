package services

import (
	"errors"
	"finance-backend/models"
	"finance-backend/repository"
)

func CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email required")
	}

	if user.Role != "admin" && user.Role != "analyst" && user.Role != "viewer" {
		return errors.New("invalid role")
	}

	return repository.CreateUser(user)
}

func GetUsers() ([]models.User, error) {
	return repository.GetUsers()
}

func UpdateUser(user *models.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}