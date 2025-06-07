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

	// item
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemConroller(itemService)

	items := r.Group("/items")
	{
		items.GET("/", itemController.GetAllItems)
		items.POST("/",
			//  middleware.AuthMiddleware(),
			itemController.CreateItem)
		items.GET("/:id", itemController.GetItemByID)
		items.PUT("/:id",
			// middleware.AuthMiddleware(),
			itemController.UpdateItem)
		items.DELETE("/:id",
			//  middleware.AuthMiddleware(),
			itemController.DeleteItem)
	}

	// rental
	rentalRepository := repositories.NewRentalRepository(db)
	rentalService := services.NewRentalService(rentalRepository)
	renralController := controllers.NewRentalController(rentalService)
	rental := r.Group("/rental")
	// rental.Use(middleware.AuthMiddleware())
	{
		rental.POST("/", renralController.CreateRental)
		rental.GET("/", renralController.GetAllRentals)
		rental.GET("/:id", renralController.GetRentalByID)
		rental.PUT("/:id", renralController.UpdateRental)
	}

	// user
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUsersController(userService)

	users := r.Group("/users")
	// users.Use(middleware.AuthMiddleware())
	{
		users.GET("/", userController.GetAllUsers)
		users.POST("/", userController.CreateUser)
		users.GET("/:id", userController.GetUserByID)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.LoginUser)
		auth.POST("/register", authController.RegisterUser)
	}

	return r
}
