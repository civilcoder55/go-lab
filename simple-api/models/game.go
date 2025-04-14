package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
	Id   bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
	Year int           `json:"year" bson:"year"`
}

func (g *Game) IsEmpty() bool {
	return g.Name == ""
}
