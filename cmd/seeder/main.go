package main

import (
	"fmt"
	"inventory-system/model"
	"log"
	"os"

	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	products := []model.Product{
		{Name: "Product A", Price: 100, Category: "Category A", Description: "Description for Product A", ImagePath: "image_a.jpg"},
		{Name: "Product B", Price: 200, Category: "Category B", Description: "Description for Product B", ImagePath: "image_b.jpg"},
		{Name: "Product C", Price: 300, Category: "Category C", Description: "Description for Product C", ImagePath: "image_c.jpg"},
		{Name: "Product D", Price: 150, Category: "Category D", Description: "Description for Product D", ImagePath: "image_d.jpg"},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Failed to insert product data: %v", err)
		}
	}

	inventories := []model.Inventory{
		{ProductID: 1, Quantity: 50, Location: "Warehouse A"},
		{ProductID: 2, Quantity: 30, Location: "Warehouse B"},
		{ProductID: 3, Quantity: 100, Location: "Warehouse C"},
		{ProductID: 4, Quantity: 20, Location: "Warehouse D"},
	}

	for _, inventory := range inventories {
		if err := db.Create(&inventory).Error; err != nil {
			log.Fatalf("Failed to insert inventory data: %v", err)
		}
	}

	orders := []model.Order{
		{ProductID: 1, Quantity: 5},
		{ProductID: 1, Quantity: 9},
		{ProductID: 2, Quantity: 2},
		{ProductID: 3, Quantity: 7},
		{ProductID: 3, Quantity: 20},
		{ProductID: 4, Quantity: 1},
	}

	for _, order := range orders {
		if err := db.Create(&order).Error; err != nil {
			log.Fatalf("Failed to insert order data: %v", err)
		}
	}

	fmt.Println("Seeder: Products, Inventories, and Orders added successfully.")
}

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	db := model.ConnectDatabase(user, password, host, port, dbname)

	SeedData(db)
	fmt.Println("Seeder data successfully!")
}
