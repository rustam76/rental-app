package routes

import (
	"rental-app/controllers"
	"rental-app/repositories"
	"rental-app/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	// Buat instance service
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemConroller(itemService)

	items := r.Group("/items")
	{
		items.GET("/", itemController.GetAllItems)
		items.POST("/", itemController.CreateItem)
		items.GET("/:id", itemController.GetItemByID)
		items.PUT("/:id", itemController.UpdateItem)
		items.DELETE("/:id", itemController.DeleteItem)
	}

	return r
}