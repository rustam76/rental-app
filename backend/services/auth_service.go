package services

import (
	"rental-app/models"
	"rental-app/repositories"
)

type AuthService interface {
	GetUserByEmail(email string) (models.User, error)
	RegisterUser(user models.User) (models.User, error)
}

type authService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (s *authService) GetUserByEmail(email string) (models.User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *authService) RegisterUser(user models.User) (models.User, error) {
	return s.repository.RegisterUser(user)
}
