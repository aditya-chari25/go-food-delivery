package services

import(
	"admin-service/internal/database"
	"errors"
	"encoding/json"
	"log"
	"fmt"
)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUsersData()(string, error){
	db := database.New() // Get the DB service instance
	user, err := db.GetAllUsers()
	if err != nil {
		log.Println("Authentication error:", err)
		return "", errors.New("invalid username or password")
	}
	users := []User{}

	// Dynamically add user objects
	for _, userextract := range user {
		fmt.Println("Username:", userextract.Username)
		// fmt.Println("Password:", userextract.Password)
		// fmt.Println("Role:", userextract.Role)
		// fmt.Println("---------------------")
		users = append(users, User{Username: userextract.Username, Password: userextract.Password, Role: userextract.Role})
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil

}

func GetAllRestaurants()(string, error){
	db := database.New() // Get the DB service instance
	restaurants, err := db.GetAllRestaurants();
	if err != nil {
		log.Println("Authentication error:", err)
		return "", errors.New("invalid username or password")
	}
	jsonData,err := json.Marshal(restaurants)
	if err != nil{
		return "",err
	}
	return string(jsonData),nil
}