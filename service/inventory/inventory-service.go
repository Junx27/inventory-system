package inventory

import (
	"inventory-system/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.Preload("Product").First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventory fetched successfully",
		"inventory": gin.H{
			"id":         inventory.ID,
			"product_id": inventory.ProductID,
			"qty":        inventory.Quantity,
			"location":   inventory.Location,
			"product":    inventory.Product,
		},
	})
}

func GetInventories(c *gin.Context) {
	var inventories []model.Inventory

	if err := model.DB.Preload("Product").Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventories"})
		return
	}

	var inventoryResponses []model.InventoryResponse
	for _, inventory := range inventories {
		inventoryResponses = append(inventoryResponses, model.InventoryResponse{
			ID:        inventory.ID,
			ProductID: inventory.ProductID,
			Quantity:  inventory.Quantity,
			Location:  inventory.Location,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Inventories fetched successfully",
		"inventories": inventoryResponses,
	})
}

func AddInventory(c *gin.Context) {
	var input model.Inventory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := model.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Inventory created successfully",
		"inventory": input,
	})
}

func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	var input model.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	inventory.ProductID = input.ProductID
	inventory.Quantity = input.Quantity
	inventory.Location = input.Location

	if err := model.DB.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Inventory updated successfully",
		"inventory": inventory,
	})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	if err := model.DB.Delete(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inventory deleted successfully",
	})
}
