package db

import (
	"log"
	"time"

	"my-gin-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbURL string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Get the underlying SQL DB
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
	log.Println("Database migration completed successfully")
} 