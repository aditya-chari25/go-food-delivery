package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	RestaurantName string `json:"rest_name,omitempty"`
	RestaurantID   string `json:"rest_id,omitempty" `
	Name           string `json:"name,omitempty" `
	Quantity       int    `json:"quantity,omitempty"`
	Price          int    `json:"price,omitempty" `
	Address        string `json:"address,omitempty"`
}

type Orders struct {
	ID       primitive.ObjectID `json:"_id,omitempty"bson:"_id,omitempty"`
	Userid int `json:"userid"validate:"required"`
	Username string             `json:"username" validate:"required"`
	Orders   []OrderItem        `json:"orders,omitempty"`
}

type RestMenu struct {
	RestaurantID string `json:"rest_id,omitempty"`
}

type Restaurant struct {
	RestaurantID string `json:"restaurant_id" bson:"restaurant_id"`
	Menus        []Menu `json:"menus" bson:"menus"`
}

type Menu struct {
	Name      string  `json:"name" ,omitempty"`
	Price     float64 `json:"price" ,omitempty`
	Category  string  `json:"category" ,omitempty"`
	Available bool    `json:"available" ,omitempty"`
}