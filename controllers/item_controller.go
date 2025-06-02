package controllers

import (
	"net/http"
	"path/filepath"
	"rental-app/models"
	"rental-app/services"

	"rental-app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	service services.ItemService
}

func NewItemConroller(service services.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (ctrl *ItemController) GetAllItems(c *gin.Context) {
	items, err := ctrl.service.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (ctrl *ItemController) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadDir := "images"
	filePath, err := utils.UploadFile(file, uploadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item.Image = filepath.Base(filePath)
	ctrl.service.CreateItem(item)
	c.JSON(http.StatusOK, gin.H{"message": "Item created successfully"})
}

func (ctrl *ItemController) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.service.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *ItemController) UpdateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.service.UpdateItem(item)
	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func (ctrl *ItemController) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.service.DeleteItem(id)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
