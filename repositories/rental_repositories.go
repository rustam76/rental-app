package repositories

import (
	"rental-app/models"

	"gorm.io/gorm"
)

type RentalRepository interface {
	GetAllRentals() ([]models.Rental, error)
	GetRentalByID(id int) (models.Rental, error)
	CreateRental(rental models.Rental) (models.Rental, error)
	UpdateRental(rental models.Rental) (models.Rental, error)
}

type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) RentalRepository {
	return &rentalRepository{db: db}
}

func (r *rentalRepository) GetAllRentals() ([]models.Rental, error) {
	var rentals []models.Rental
	if err := r.db.Preload("RentalDetails").Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}

func (r *rentalRepository) GetRentalByID(id int) (models.Rental, error) {
	var rental models.Rental
	if err := r.db.Preload("RentalDetails").First(&rental, id).Error; err != nil {
		return models.Rental{}, err
	}
	return rental, nil
}

func (r *rentalRepository) CreateRental(rental models.Rental) (models.Rental, error) {
	trx := r.db.Begin()

	if err := trx.Create(&rental).Error; err != nil {
		trx.Rollback()
		return models.Rental{}, err
	}

	for i := range rental.RentalDetails {
		rental.RentalDetails[i].RentalID = rental.ID
	}

	if err := trx.Create(&rental.RentalDetails).Error; err != nil {
		trx.Rollback()
		return models.Rental{}, err
	}

	trx.Commit()
	return rental, nil
}

func (r *rentalRepository) UpdateRental(rental models.Rental) (models.Rental, error) {
	trx := r.db.Begin()

	if err := trx.Model(&models.Rental{}).Where("id = ?", rental.ID).Updates(rental).Error; err != nil {
		trx.Rollback()
		return models.Rental{}, err
	}

	if len(rental.RentalDetails) > 0 {
		if err := trx.Delete(&models.RentalDetail{}, "rental_id = ?", rental.ID).Error; err != nil {
			trx.Rollback()
			return models.Rental{}, err
		}

		for i := range rental.RentalDetails {
			rental.RentalDetails[i].RentalID = rental.ID
		}

		if err := trx.Create(&rental.RentalDetails).Error; err != nil {
			trx.Rollback()
			return models.Rental{}, err
		}

	}

	trx.Commit()
	return rental, nil
}
