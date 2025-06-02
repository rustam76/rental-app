package config

import (
	"log"
	"rental-app/models"
)

func RunMigration() {
	err := DB.AutoMigrate(&models.Item{}, &models.User{}, &models.Rental{}, &models.RentalDetail{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")
}
