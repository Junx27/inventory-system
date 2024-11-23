package order

import (
	"inventory-system/helper"
	"inventory-system/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.Preload("Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Order not found"))
		return
	}
	orderResponse := model.OrderResponse{

		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		DateOrder: order.DateOrder,
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Order fetched successfully", orderResponse))
}

func GetOrders(c *gin.Context) {
	var orders []model.Order

	if err := model.DB.Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to fetch orders"))
		return
	}

	var orderResponses []model.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, model.OrderResponse{
			ID:        order.ID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			DateOrder: order.DateOrder,
		})
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Order fetched successfully", orderResponses))
}

func AddOrder(c *gin.Context) {
	var input model.Order

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}

	if err := model.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to create order"))
		return
	}
	newOrder := model.OrderResponse{
		ID:        input.ID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		DateOrder: input.DateOrder,
	}

	c.JSON(http.StatusCreated, helper.SuccessResponse("Order created successfully", newOrder))
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Order not found"))
		return
	}

	var input model.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input"))
		return
	}

	order.ProductID = input.ProductID
	order.Quantity = input.Quantity
	order.DateOrder = input.DateOrder

	if err := model.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update order"))
		return
	}
	newOrder := model.OrderResponse{
		ID:        input.ID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		DateOrder: input.DateOrder,
	}

	c.JSON(http.StatusCreated, helper.SuccessResponse("Order updated successfully", newOrder))

}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Order not found"))
		return
	}

	if err := model.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete order"))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Order deleted successfully", id))
}
