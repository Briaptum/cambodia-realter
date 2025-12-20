package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS configures CORS middleware for the Gin router
func SetupCORS(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000", 
		"http://localhost:5173",
		"http://frontend:3000", // Docker internal network
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{
		"Origin", 
		"Content-Type", 
		"Accept", 
		"Authorization",
		"X-Requested-With",
	}
	config.AllowCredentials = true
	
	r.Use(cors.New(config))
}
