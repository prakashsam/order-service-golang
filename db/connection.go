package db

import (
	"fmt"
	"log"
	"orderservice/models"
	"os"

	spannergorm "github.com/googleapis/go-gorm-spanner"
	// "gorm.io/driver/postgres" // For Local development
	"gorm.io/gorm"
)

var (
	dbconnect *gorm.DB
)

func InitDBConnection() {
	projectID := os.Getenv("GCP_PROJECT")
	instanceID := os.Getenv("SPANNER_INSTANCE_ID")
	databaseID := os.Getenv("SPANNER_DATABASE_ID")
	dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instanceID, databaseID)
	db, err := gorm.Open(spannergorm.New(spannergorm.Config{DriverName: "spanner", DSN: dsn}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Spanner: %v", err)
		return
	}
	err = db.AutoMigrate(
		&models.Order{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return
	}
	dbconnect = db
}

//****For Local development*********

// func InitDBConnection() {
// 	dsn := "host=localhost user=postgres password=passpass dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to Spanner: %v", err)
// 		return
// 	}
// 	err = db.AutoMigrate(
// 		&models.Order{},
// 	)
// 	if err != nil {
// 		log.Fatalf("Migration failed: %v", err)
// 		return
// 	}
// 	dbconnect = db
// }

func GetDBConnection() *gorm.DB {
	return dbconnect
}
