package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	
}

type Game struct{

	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	UserID *User `json:"user_id"`
	Point int `json:"point"`
}