package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() {
	// Build DSN from individual environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "postgres" // Default for Docker Compose
	}
	
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "cambodia_realter"
	}
	
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
		host, user, password, dbname, port, sslmode)
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		// Don't exit, allow the app to start without DB for development
		return
	}

	log.Println("Database connected successfully")
	
	// Auto-migrate models here when we add them
	// DB.AutoMigrate(&models.User{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
