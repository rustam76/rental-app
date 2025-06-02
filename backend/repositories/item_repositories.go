package repositories

import (
	"rental-app/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id int) (models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
	UpdateItem(item models.Item) (models.Item, error)
	DeleteItem(id int) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) GetItemByID(id int) (models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, id).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (r *itemRepository) CreateItem(item models.Item) (models.Item, error) {
	if err := r.db.Create(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (r *itemRepository) UpdateItem(item models.Item) (models.Item, error) {
	if err := r.db.Save(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (r *itemRepository) DeleteItem(id int) error {
	if err := r.db.Delete(&models.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}
