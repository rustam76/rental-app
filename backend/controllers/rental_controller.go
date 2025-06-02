package controllers

import (
	"net/http"
	"rental-app/models"
	"rental-app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentalController struct {
	service services.RentalService
}

func NewRentalController(service services.RentalService) *RentalController {
	return &RentalController{service: service}
}

func (ctrl *RentalController) GetAllRentals(c *gin.Context) {
	rentals, err := ctrl.service.GetAllRentals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rentals)
}

func (ctrl *RentalController) CreateRental(c *gin.Context) {
	var rental models.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.service.CreateRental(rental)
	c.JSON(http.StatusOK, gin.H{"message": "Rental created successfully"})
}

func (ctrl *RentalController) GetRentalByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rental, err := ctrl.service.GetRentalByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rental)
}

func (ctrl *RentalController) UpdateRental(c *gin.Context) {
	var rental models.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.service.UpdateRental(rental)
	c.JSON(http.StatusOK, gin.H{"message": "Rental updated successfully"})
}
