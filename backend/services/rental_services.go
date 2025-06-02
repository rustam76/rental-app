package services

import (
	"rental-app/models"
	"rental-app/repositories"
)

type RentalService interface {
	GetAllRentals() ([]models.Rental, error)
	GetRentalByID(id int) (models.Rental, error)
	CreateRental(rental models.Rental) (models.Rental, error)
	UpdateRental(rental models.Rental) (models.Rental, error)
}

type rentalService struct {
	repository repositories.RentalRepository
}

func NewRentalService(repository repositories.RentalRepository) RentalService {
	return &rentalService{repository: repository}
}

func (s *rentalService) GetAllRentals() ([]models.Rental, error) {
	return s.repository.GetAllRentals()
}

func (s *rentalService) GetRentalByID(id int) (models.Rental, error) {
	return s.repository.GetRentalByID(id)
}

func (s *rentalService) CreateRental(rental models.Rental) (models.Rental, error) {
	return s.repository.CreateRental(rental)
}

func (s *rentalService) UpdateRental(rental models.Rental) (models.Rental, error) {
	return s.repository.UpdateRental(rental)
}
