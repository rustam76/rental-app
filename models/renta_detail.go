package models

import "time"

type RentalDetail struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RentalID   int       `gorm:"not null" json:"rental_id"`
	ItemID     int       `gorm:"not null" json:"item_id"`
	Quantity   int       `gorm:"type:int;not null" json:"quantity"`
	TotalPrice int       `gorm:"type:int;not null" json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
