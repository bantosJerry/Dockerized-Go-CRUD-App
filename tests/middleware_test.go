package tests

import (
	"go-crud-app/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAuthenticateMiddleware tests the behavior of the authentication middleware.
func TestAuthenticateMiddleware(t *testing.T) {
	// Initialize the AuthHandler for handling the authentication
	authHandler := handlers.AuthHandler{}

	// Generate a valid token for testing (using a user ID of 1)
	validToken, _ := handlers.GenerateJWT(1)

	// Define a slice of test cases
	tests := []struct {
		name       string
		token      string
		expectCode int
	}{
		{"ValidToken", "Bearer " + validToken, http.StatusOK},
		{"NoToken", "", http.StatusUnauthorized},
		{"InvalidToken", "Bearer invalid.token", http.StatusUnauthorized},
	}

	// Loop over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a simple next handler that will return HTTP 200 status code
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// Create a new recorder to capture the HTTP response
			rec := httptest.NewRecorder()

			// Create a new HTTP request for the test case with the provided Authorization header
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tt.token)

			// Call the Authenticate middleware with the next handler
			handler := authHandler.Authenticate(nextHandler)

			// Execute the handler and capture the response
			handler.ServeHTTP(rec, req)

			// Assert that the HTTP status code returned matches the expected code
			assert.Equal(t, tt.expectCode, rec.Code)
		})
	}
}
