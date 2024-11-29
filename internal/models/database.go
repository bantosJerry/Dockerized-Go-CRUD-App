package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Retry connection up to 10 times with a delay of 2 seconds between retries
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database")
			break
		}
		log.Printf("Failed to connect to the database: %v. Retrying...\n", err)
		time.Sleep(2 * time.Second)
	}

	// Handle the case if the connection failed after retries
	if err != nil {
		log.Fatalf("Could not connect to the database after multiple attempts: %v", err)
	}

	// Drop the table if needed (be cautious)
	// To drop the "users" table, uncomment the following line
	// err = DB.Migrator().DropTable(&User{})
	// if err != nil {
	// 	log.Fatalf("Failed to drop the table: %v", err)
	// }

	// Run AutoMigrate to sync the User model with the database schema
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	log.Println("Database migration successful")
}
