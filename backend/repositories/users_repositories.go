package repositories

import (
	"rental-app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
