package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	Username string
	Password string
	Role     string
}

type SignUp struct{
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	ID primitive.ObjectID `json:"_id,omitempty"bson:"_id,omitempty"`
	Address string `json:"address,omitempty`
	Rating float64 `json:"rating,omitempty`
	Phonenumber string `json:"phone_number,omitempty"`
}

// username : user2;
// address : 
// rating : 
// phone_number :


