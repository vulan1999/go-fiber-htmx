package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Movie struct {
	ID     bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string        `json:"title" bson:"title"`
	Plot   string        `json:"plot" bson:"plot"`
	Poster string        `json:"poster" bson:"poster"`
	Year   string        `json:"year" bson:"year"`
}
