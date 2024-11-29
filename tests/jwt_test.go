package tests

import (
	"go-crud-app/internal/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndVerifyJWT(t *testing.T) {
	// Generate a token
	userID := uint(1)
	token, err := handlers.GenerateJWT(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify the token
	verifiedUserID, err := handlers.VerifyJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, verifiedUserID)
}

func TestInvalidJWT(t *testing.T) {
	// Test an invalid token
	invalidToken := "invalid.token.value"
	_, err := handlers.VerifyJWT(invalidToken)
	assert.Error(t, err)
}
