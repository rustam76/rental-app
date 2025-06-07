package repositories

import (
	"rental-app/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckUser(email, password string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	RegisterUser(user models.User) (models.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) CheckUser(email, password string) (models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}

	// check password hash
	// err != utils.ComparePassword(user.Password, password)

	// if err != nil {
	// 	return models.User{}, err
	// }

	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepository) RegisterUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
