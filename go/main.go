package main

import (
	"log"

	"go-admin/config"
	"go-admin/database"
	"go-admin/models"
	"go-admin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	config.Init()

	// Connect to the database
	database.Init()

	// Automatically migrate table structures
	models.Migrate()

	// Create a Gin engine
	r := gin.Default()

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	log.Printf("Server starting on port %s...", config.AppConfig.ServerPort)
	if err := r.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
