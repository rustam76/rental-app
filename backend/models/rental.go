package models

import "time"

type Rental struct {
	ID            int            `grom:"primaryKey;autoIncrement" json:"id"`
	UserID        int            `grom:"not null" json:"user_id"`
	ItemID        int            `grom:"not null" json:"item_id"`
	RentalDate    time.Time      `grom:"not null" json:"rental_date"`
	ReturnDate    time.Time      `grom:"not null" json:"return_date"`
	CreatedAt     time.Time      `grom:"not null" json:"created_at"`
	UpdatedAt     time.Time      `grom:"not null" json:"updated_at"`
	RentalDetails []RentalDetail `grom:"foreignKey:RentalID" json:"rental_details"`
}

// join ke table rental_detail
