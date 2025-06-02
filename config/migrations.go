package config

import (
    "log"
    "rental-app/models"
)

func RunMigration() {
    err := DB.AutoMigrate(&models.Item{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    log.Println("Database migration completed successfully")
}
