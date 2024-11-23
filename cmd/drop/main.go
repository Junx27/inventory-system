package main

import (
	"fmt"
	"log"
	"os"

	"inventory-system/model"
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

	err := db.Migrator().DropTable(&model.Product{}, &model.Inventory{}, &model.Order{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	fmt.Println("Tables dropped successfully!")
}
