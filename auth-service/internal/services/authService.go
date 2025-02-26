// internal/services/authService.go
package services

import (
	"errors"
	"auth-service/internal/utils"
)

// User struct to hold user information
type User struct {
	Password string
	Role     string
}

// Mock user database (Replace with actual DB lookup)
var users = map[string]User{
	"admin": {Password: "password123", Role: "admin"},
	"user":  {Password: "password456", Role: "user"},
}

// AuthenticateUser  checks user credentials and returns a JWT token
func AuthenticateUser(username, password string) (string, error) {
	user, exists := users[username]
	if !exists || user.Password != password {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token with role
	return utils.GenerateToken(username, user.Role)
}
