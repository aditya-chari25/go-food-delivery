package utils

import (
	"errors"
	"time"
	"log"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key") // Change this in production

// GenerateToken creates a JWT token for a user
func GenerateToken(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken verifies the JWT token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	log.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	// Handle parsing errors
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Ensure the token is valid
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	// Extract claims safely
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
