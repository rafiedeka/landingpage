package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
