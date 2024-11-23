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
		{Name: "Product A", Price: 100, Category: "Category A", Description: "Description for Product A"},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Failed to insert product data: %v", err)
		}
	}

	inventories := []model.Inventory{
		{ProductID: 1, Quantity: 50, Location: "Warehouse A"},
	}

	for _, inventory := range inventories {
		if err := db.Create(&inventory).Error; err != nil {
			log.Fatalf("Failed to insert inventory data: %v", err)
		}
	}

	orders := []model.Order{
		{ProductID: 1, Quantity: 5, DateOrder: "12 Agustus 2024"},
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
