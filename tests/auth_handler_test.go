package tests

import (
	"bytes"
	"encoding/json"
	"go-crud-app/internal/handlers"
	"go-crud-app/internal/models"
	"go-crud-app/internal/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupTestDB sets up the test database, ensuring a clean state before each test
func setupTestDB() *gorm.DB {
	// Fetch database parameters from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Construct the DSN (Data Source Name) string
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
		" dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Migrate the User model to ensure the `users` table exists
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Start a transaction to ensure changes are discarded after each test
	tx := db.Begin()

	// Check if the `users` table exists before truncating
	if tx.Migrator().HasTable(&models.User{}) {
		// Ensure the Users table is empty and reset auto-incrementing IDs
		tx.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	} else {
		log.Println("Users table does not exist, skipping truncate.")
	}

	// Return the transaction (for use in tests)
	return tx
}

// TestRegisterHandler tests the user registration handler
func TestRegisterHandler(t *testing.T) {
	// Setup the test database with a fresh transaction
	db := setupTestDB()
	defer db.Rollback() // Rollback the transaction after the test

	router := routes.SetupRoutes(db)

	// Create the request body for user registration
	body := map[string]string{
		"username": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
	}
	bodyJSON, _ := json.Marshal(body)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the response status code and check if the token is present in the response
	assert.Equal(t, http.StatusCreated, rec.Code)
	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Contains(t, response, "token")
}

// TestLoginHandler tests the user login handler
func TestLoginHandler(t *testing.T) {
	// Setup the test database with a fresh transaction
	db := setupTestDB()
	defer db.Rollback() // Rollback the transaction after the test

	// Seed a user for login (check if the user already exists, then create if not)
	hashedPassword, _ := handlers.HashPassword("password123")
	db.Where(models.User{Username: "testuser"}).FirstOrCreate(&models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: hashedPassword,
	})

	router := routes.SetupRoutes(db)

	// Create the request body for login
	body := map[string]string{
		"username": "testuser",
		"password": "password123",
	}
	bodyJSON, _ := json.Marshal(body)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert the response status code and check if the token is present in the response
	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Contains(t, response, "token")
}
