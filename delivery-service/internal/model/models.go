package model

type AuthDelivery struct {
	DeliveryID  string `json:"delivery_id,omitempty"`
	Password    string `json:"password,omitempty"`
	Phonenumber string `json:"phoneno,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
type DeliveryDriver struct {
	DeliveryID  string   `json:"delivery_id"`
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Gender      string   `json:"gender"`
	CurrentLoc  Location `json:"current_loc"`
	Rating      float64  `json:"rating"`
	PhoneNumber string   `json:"phone_number"`
}
