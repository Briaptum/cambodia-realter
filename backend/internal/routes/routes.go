package routes

import (
	"go-vue-base/internal/controllers"
	"go-vue-base/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(r *gin.Engine) {
	// Initialize controllers
	healthController := controllers.NewHealthController()
	authController := controllers.NewAuthController()
	contactRequestController := controllers.NewContactRequestController()

	// Public API routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", healthController.HealthCheck)

		// Auth routes (public)
		api.POST("/auth/login", authController.Login)
		api.GET("/auth/logout", authController.Logout)

		// Contact requests (public)
		api.POST("/contact-requests", contactRequestController.CreateContactRequest)
	}

	// Protected API routes (require authentication)
	protected := api.Group("")
	protected.Use(middleware.Auth0Middleware())
	{
		// Profile endpoint
		protected.GET("/profile", authController.Profile)

		// Contact requests (protected)
		protected.GET("/contact-requests", contactRequestController.GetContactRequests)
		protected.GET("/contact-requests/:id", contactRequestController.GetContactRequest)
		protected.DELETE("/contact-requests/:id", contactRequestController.DeleteContactRequest)
	}

	// Future API versions can be added here
	// v2 := r.Group("/api/v2")
}
