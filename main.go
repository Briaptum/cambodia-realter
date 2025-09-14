package main

import (
	"log"
	"os"

	"manage/internal/config"
	"manage/internal/middleware"
	"manage/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.InitDatabase()

	// Initialize Gin router
	r := gin.New()

	// Add middleware
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// Configure CORS
	config.SetupCORS(r)

	// Setup routes
	routes.SetupRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📚 API Documentation available at http://localhost:%s/api/health", port)
	
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
