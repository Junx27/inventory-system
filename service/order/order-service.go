package order

import (
	"inventory-system/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.Preload("Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order fetched successfully",
		"order": gin.H{
			"id":         order.ID,
			"product_id": order.ProductID,
			"qty":        order.Quantity,
			"product":    order.Product,
		},
	})
}

func GetOrders(c *gin.Context) {
	var orders []model.Order

	if err := model.DB.Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	var orderResponses []model.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, model.OrderResponse{
			ID:        order.ID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All orders fetched successfully",
		"orders":  orderResponses,
	})
}

func AddOrder(c *gin.Context) {
	var input model.Order

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := model.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order": gin.H{
			"id":         input.ID,
			"product_id": input.ProductID,
			"qty":        input.Quantity,
		},
	})
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input model.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order.ProductID = input.ProductID
	order.Quantity = input.Quantity

	if err := model.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order": gin.H{
			"id":         order.ID,
			"product_id": order.ProductID,
			"qty":        order.Quantity,
		},
	})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order model.Order

	if err := model.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := model.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
