// internal/services/authService.go
package services

import (
	"auth-service/internal/database"
	"errors"
	"auth-service/internal/utils"
	"log"
)

// User struct to hold user information
type User struct {
	Username string
	Password string
	Role     string
}

// Mock user database (Replace with actual DB looku p)
var users = map[string]User{
	"admin": {Password: "password123", Role: "admin"},
	"user":  {Password: "password456", Role: "user"},
}

// AuthenticateUser  checks user credentials and returns a JWT token
func AuthenticateUser(username, password string) (string, error) {
	db := database.New() // Get the DB service instance
	user, err := db.GetUser(username)
	if err != nil {
		log.Println("Authentication error:", err)
		return "", errors.New("invalid username or password")
	}

	// Check password (use hashed passwords in production)
	if user.Password != password {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token with role
	return utils.GenerateToken(user.Username, user.Role)
}
