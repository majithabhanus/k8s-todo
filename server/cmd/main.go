package main

import (
	"log"
	"github.com/sujin/todo-app/config"
	"github.com/sujin/todo-app/database"
	"github.com/sujin/todo-app/routes"

	"github.com/gin-gonic/gin"
)

// @title Todo API
// @version 1.0
// @description This is a Todo API with JWT authentication
// @host localhost:9000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load config (like DB credentials, JWT secret)
	config.LoadConfig()

	// Connect to  MySQL 
	database.ConnectDB()

	// Init Gin
	r := gin.Default()

	// Setup routes
	routes.RegisterRoutes(r)

	// Run server
	if err := r.Run(":9000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
