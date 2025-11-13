package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("Loaded DATABASE_URL:", dsn)

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	// Ensure sslmode=disable for local development
	if !strings.Contains(dsn, "sslmode=") {
		if strings.Contains(dsn, "?") {
			dsn += "&sslmode=disable"
		} else {
			dsn += "?sslmode=disable"
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Test database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
