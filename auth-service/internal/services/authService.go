// internal/services/authService.go
package services

import (
	"auth-service/internal/database"
	"errors"
	"auth-service/internal/utils"
	"auth-service/internal/model"
	"log"
)

// User struct to hold user information
// Mock user database (Replace with actual DB looku p)


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

func AuthenticateDriver(username,password string)(string,error){
	db := database.New()
	user, err := db.GetDeliveryDriver(username)
	if err != nil{
		log.Println("Athenticate Driver:",err)
		return "",errors.New("invalid username or password")
	}
	if user.Password != password {
		return "", errors.New("invalid username or password")
	}
	return utils.GenerateToken(user.DeliveryID, "deliverydriver")

}

func VerifyUser(token string)(string, error) {
	claims,err := utils.ValidateToken(token);
	if err != nil{
		log.Println("error:",err)
		return "",err
	}
	log.Println(claims)
	username, ok1 := claims["username"].(string)
	if !ok1{
		log.Fatal("Error extracting JWT claims")
	}
	return username,nil
}

func SignUser(userJson model.SignUp)(string,error){
	db := database.New();
	msg,err := db.Signup(userJson);
	if err != nil{
		return "",err
	}
	return msg,nil
}
