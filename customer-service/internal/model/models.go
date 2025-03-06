package model

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct{
	RestaurantName string `json:"rest_name,omitempty"`
	RestaurantID   string `json:"rest_id,omitempty" `
	Name           string `json:"name,omitempty" `
	Quantity       int    `json:"quantity,omitempty"`
	Price          string `json:"price,omitempty" `
	Address        string `json:"address,omitempty"`
}

type Orders struct{
	ID primitive.ObjectID `json:"_id,omitempty"bson:"_id,omitempty"`
	username string `json:"username,omitempty"`
	Orders   []OrderItem  `json:"orders,omitempty"`
}