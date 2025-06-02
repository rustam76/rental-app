package services

import (
	"rental-app/models"
	"rental-app/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

func (s *userService) GetUserByID(id int) (models.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.repository.CreateUser(user)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	return s.repository.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}