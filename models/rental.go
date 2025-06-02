package models

import "time"

type Rental struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ItemID      int    `json:"item_id"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}