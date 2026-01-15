package database

import (
	"fmt"
	"log"

	"go-admin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	log.Printf("Connecting to database: host=%s port=%s user=%s dbname=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBName,
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}
