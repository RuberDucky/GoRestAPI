package services

import (
	"my-gin-project/db"
	"my-gin-project/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) error {
	// Hash password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result := db.DB.Create(user)
	return result.Error
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
} 