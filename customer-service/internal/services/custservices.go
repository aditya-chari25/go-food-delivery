package services

import(
	"customer-service/internal/database"
	"customer-service/internal/model"
	// "fmt"
	// "net/http"
	// "context"
	// "encoding/json"
)

func PlaceOrder(customerJSON model.Orders)(string,error){
	db := database.New();
	orderplaced,error := db.PlaceOrder(customerJSON)
	if error != nil{
		return "",error
	}
	return orderplaced,nil
}

func RestMenu (RestMenuJSON model.RestMenu)(*model.Restaurant,error){
	db := database.New();
	menu,error := db.RetRestMenu(RestMenuJSON)
	if error != nil{
		return  nil,error
	}
	return menu,nil
}