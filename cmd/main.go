package main

import (
	"rental-app/config"
	"rental-app/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter(config.DB)
	config.RunMigration()
	r.Run(":8080")
}
