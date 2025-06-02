package models

import "time"

type Item struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int       `gorm:"type:int;not null" json:"price"`
	Stock       int       `gorm:"type:int;not null" json:"stock"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	CategoryID  int       `gorm:"not null" json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
