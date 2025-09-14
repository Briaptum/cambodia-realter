package routes

import (
	"manage/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine) {
	// Initialize controllers
	userController := controllers.NewUserController()
	healthController := controllers.NewHealthController()

	// API v1 routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", healthController.HealthCheck)
		
		// User routes
		api.GET("/name", userController.GetName) // Used by frontend
	}

	// Future API versions can be added here
	// v2 := r.Group("/api/v2")
}
