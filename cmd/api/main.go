package main

import (
	"inventory-system/config"
	"inventory-system/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		panic(err)
	}
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.SetTrustedProxies([]string{"127.0.0.1"})
	config.LoadEnv()
	config.InitializeDatabase()
	router.ServerRoutes()

	r.Run(":" + port)
}
