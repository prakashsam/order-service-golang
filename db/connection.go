package db

import (
	"fmt"
	"log"
	"orderservice/config"
	"orderservice/models"
	"orderservice/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbconnect *gorm.DB

func InitDBConnection() {
	cfg := config.Load()

	dbPassword, err := utils.GetSecret("DB_PASSWORD", cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to get DB password from Secret Manager: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		dbPassword,
		cfg.DBName,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := db.AutoMigrate(&models.Order{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	dbconnect = db
}

func GetDBConnection() *gorm.DB {
	return dbconnect
}
