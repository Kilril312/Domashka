package db

import (
	"domashka/internal/tasksService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user= postgres password=qwerty dbname=postgres port=5438 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&tasksService.RequestBodyTask{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
	return db, nil
}
