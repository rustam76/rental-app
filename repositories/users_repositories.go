package repositories

import (
	"rental-app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(db *gorm.DB) ([]models.User, error)
	GetUserByID(db *gorm.DB, id int) (models.User, error)
	CreateUser(db *gorm.DB, user models.User) (models.User, error)
	UpdateUser(db *gorm.DB, user models.User) (models.User, error)
	DeleteUser(db *gorm.DB, id int) error
}


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(db *gorm.DB, id int) (models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(db *gorm.DB, user models.User) (models.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(db *gorm.DB, user models.User) (models.User, error) {
	if err := db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(db *gorm.DB, id int) error {
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}