package model

// bson is used for _id ; giving unique id to our data elements

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// if read lines occurs under the imports , then re download that particular package, here go get go.mongodb.org/mongo-driver/mongo and then go mod tidy to neatly bind the code with local cache where the pkg is installed 

type Netflix struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string	`json:"movie,omitempty"`
	Watched bool `json:"watched,omitempty"`
}