package services

import(
	"customer-service/internal/database"
	"fmt",
	"net/http",
	"context"
	"encoding/json"
)

func PlaceOrder(customerJSON string)(){
	db := database.New();
	orderplaced,err := db.PlaceOrder(customerJSON string);


}