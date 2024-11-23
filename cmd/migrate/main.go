package main

import (
	"fmt"
	"inventory-system/model"
	"log"
	"os"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	db := model.ConnectDatabase(user, password, host, port, dbname)
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		sqlDB.Close()
	}()

	err := db.AutoMigrate(&model.Product{}, &model.Inventory{}, &model.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	fmt.Println("Tables migrated successfully!")
}
