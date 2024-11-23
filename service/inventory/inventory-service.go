package inventory

import (
	"inventory-system/helper"
	"inventory-system/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.Preload("Product").First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Inventory not found"))
		return
	}
	inventoryResponse := model.InventoryResponse{
		ID:        inventory.ID,
		ProductID: inventory.ProductID,
		Quantity:  inventory.Quantity,
		Location:  inventory.Location,
	}
	c.JSON(http.StatusOK, helper.SuccessResponse("Inventory fetched successfully", inventoryResponse))
}

func GetInventories(c *gin.Context) {
	var inventories []model.Inventory

	if err := model.DB.Preload("Product").Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve inventories"))
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

	c.JSON(http.StatusOK, helper.SuccessResponse("Inventories fetched successfully", inventoryResponses))
}

func AddInventory(c *gin.Context) {
	var input model.Inventory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}

	if err := model.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create inventory"))
		return
	}
	inventoryResponse := model.InventoryResponse{
		ID:        input.ID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Location:  input.Location,
	}

	c.JSON(http.StatusCreated, helper.SuccessResponse("Inventory created successfully", inventoryResponse))
}

func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Inventory not found"))
		return
	}

	var input model.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}

	inventory.ProductID = input.ProductID
	inventory.Quantity = input.Quantity
	inventory.Location = input.Location

	if err := model.DB.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update inventory"))
		return
	}
	inventoryResponse := model.InventoryResponse{
		ID:        input.ID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		Location:  input.Location,
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Inventory updated successfully", inventoryResponse))
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory model.Inventory

	if err := model.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Inventory not found"))
		return
	}

	if err := model.DB.Delete(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete inventory"))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Inventory deleted successfully", id))
}
