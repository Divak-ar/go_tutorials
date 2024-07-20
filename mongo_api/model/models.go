package model

// bson is used for _id ; giving unique id to our data elements

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string	`json:"movie,omitempty"`
	Watched bool `json:"watched,omitempty"`
}