package router

import (
	"inventory-system/service/inventory"
	"inventory-system/service/order"
	"inventory-system/service/product"

	"github.com/gin-gonic/gin"
)

func ServerRoutes() {
	r := gin.Default()
	products := r.Group("api/v1/products")
	{
		products.POST("", product.AddProduct)
		products.GET("", product.GetProducts)
		products.GET(":id", product.GetProductByID)
		products.PUT(":id", product.UpdateProduct)
		products.DELETE(":id", product.DeleteProduct)

		products.POST(":id/upload-image", product.UploadProductImage)
		products.DELETE(":id/delete-image", product.DeleteProductImage)
		products.GET(":id/download-image", product.DownloadProductImage)
	}
	inventories := r.Group("api/v1/inventories")
	{
		inventories.POST("", inventory.AddInventory)
		inventories.GET("", inventory.GetInventories)
		inventories.GET(":id", inventory.GetInventoryByID)
		inventories.PUT(":id", inventory.UpdateInventory)
		inventories.DELETE(":id", inventory.DeleteInventory)
	}
	orders := r.Group("api/v1/orders")
	{
		orders.POST("", order.AddOrder)
		orders.GET("", order.GetOrders)
		orders.GET(":id", order.GetOrderByID)
		orders.PUT(":id", order.UpdateOrder)
		orders.DELETE(":id", order.DeleteOrder)
	}
	_ = r.Run()
}
