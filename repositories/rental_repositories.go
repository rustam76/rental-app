package repositories


import (
	"rental-app/models"
	"gorm.io/gorm"
)

type RentalRepository interface {
	GetAllRentals(db *gorm.DB) ([]models.Rental, error)
	GetRentalByID(db *gorm.DB, id int) (models.Rental, error)
	CreateRental(db *gorm.DB, rental models.Rental) (models.Rental, error)
	UpdateRental(db *gorm.DB, rental models.Rental) (models.Rental, error)
	DeleteRental(db *gorm.DB, id int) error
}


type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepository{db: db}
}

func (r *rentalRepository) GetAllRentals(db *gorm.DB) ([]models.Rental, error) {
	var rentals []models.Rental
	if err := db.Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *rentalRepository) GetRentalByID(db *gorm.DB, id int) (models.Rental, error) {
	var rental models.Rental
	if err := db.First(&rental, id).Error; err != nil {
		return models.Rental{}, err
	}
	return rental, nil
}

func (r *rentalRepository) CreateRental(db *gorm.DB, rental models.Rental) (models.Rental, error) {
	if err := db.Create(&rental).Error; err != nil {
		return models.Rental{}, err
	}
	return rental, nil
}

func (r *rentalRepository) UpdateRental(db *gorm.DB, rental models.Rental) (models.Rental, error) {
	if err := db.Save(&rental).Error; err != nil {
		return models.Rental{}, err
	}
	return rental, nil
}

func (r *rentalRepository) DeleteRental(db *gorm.DB, id int) error {
	if err := db.Delete(&models.Rental{}, id).Error; err != nil {
		return err
	}
	return nil
}