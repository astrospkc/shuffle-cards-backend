package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	
	Name string `json:"name"`

}

type Game struct{

	ID primitive.ObjectID `json:"_id, omitempty" bson.M:"_id,omitempty"`
	UserID string `json:"user_id"`
	Point int `json:"point"`
}

// constructor function to create a new Game with default values


// type User1 struct {
//     ID      string `json:"id"`
//     Name    string `json:"name"`
//     College string `json:"college"`
// }

// type Marks struct {
//     UserID string `json:"user_id"`
//     Subject string `json:"subject"`
//     Score   int    `json:"score"`
// }


// func (game *Game) MarshalBinary() ([]byte, error){
// 	return json.Marshal(game)
// }