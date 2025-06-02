package services

import (
	"rental-app/models"
	"rental-app/repositories"
)

type ItemService interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id int) (models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
	UpdateItem(item models.Item) (models.Item, error)
	DeleteItem(id int) error
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAllItems()
}

func (s *itemService) GetItemByID(id int) (models.Item, error) {
	return s.repo.GetItemByID(id)
}

func (s *itemService) CreateItem(item models.Item) (models.Item, error) {

	
	return s.repo.CreateItem(item)
}

func (s *itemService) UpdateItem(item models.Item) (models.Item, error) {
	return s.repo.UpdateItem(item)
}

func (s *itemService) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}
