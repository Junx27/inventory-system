package product

import (
	"database/sql"
	"inventory-system/helper"
	"inventory-system/model"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var productUploadDir = "uploads/products"

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.Preload("Inventory").Preload("Order").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	var productResponse model.ProductResponse
	productResponse.FillFromModel(product)

	c.JSON(http.StatusOK, helper.SuccessResponse("Product fetched successfully", productResponse))
}

func GetProducts(c *gin.Context) {
	var products []model.Product

	if err := model.DB.Preload("Inventory").Preload("Order").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve products"))
		return
	}

	var productResponses []model.ProductResponse
	for _, product := range products {
		var productResponse model.ProductResponse
		productResponse.FillFromModel(product)
		productResponses = append(productResponses, productResponse)
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Products fetched successfully", productResponses))
}

func AddProduct(c *gin.Context) {
	var newProduct model.Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input data"))
		return
	}

	if err := model.DB.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to add product"))
		return
	}

	var productResponse model.ProductResponse
	productResponse.FillFromModel(newProduct)

	c.JSON(http.StatusCreated, helper.SuccessResponse("Product added successfully", productResponse))
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid input data"))
		return
	}

	if err := model.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update product"))
		return
	}

	var productResponse model.ProductResponse
	productResponse.FillFromModel(product)

	c.JSON(http.StatusOK, helper.SuccessResponse("Product updated successfully", productResponse))
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	if product.ImagePath.Valid && product.ImagePath.String != "" {
		if err := os.Remove(product.ImagePath.String); err != nil {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete image"))
			return
		}
	}

	if err := model.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete product"))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Product deleted successfully", id))
}

func UploadProductImage(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Image is required"))
		return
	}

	imagePath := filepath.Join(productUploadDir, id+filepath.Ext(file.Filename))

	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to save image"))
		return
	}

	product.ImagePath = sql.NullString{String: imagePath, Valid: true}
	if err := model.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to save product image path"))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Image uploaded successfully", imagePath))
}

func DeleteProductImage(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	if !product.ImagePath.Valid || product.ImagePath.String == "" {
		c.JSON(http.StatusNotFound, helper.FailedResponse("No image to delete"))
		return
	}

	if err := os.Remove(product.ImagePath.String); err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to delete image"))
		return
	}

	product.ImagePath = sql.NullString{String: "", Valid: false}
	if err := model.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to update product after deleting image"))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("Image deleted successfully", id))
}

func DownloadProductImage(c *gin.Context) {
	id := c.Param("id")
	var product model.Product

	if err := model.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, helper.FailedResponse("Product not found"))
		} else {
			c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to retrieve product"))
		}
		return
	}

	if !product.ImagePath.Valid || product.ImagePath.String == "" {
		c.JSON(http.StatusNotFound, helper.FailedResponse("Image not found"))
		return
	}

	c.File(product.ImagePath.String)
}
