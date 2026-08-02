package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Define type
type Movie struct {
	ID       bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string        `json:"title" bson:"title"`
	Plot     string        `json:"plot" bson:"plot"`
	Poster   string        `json:"poster" bson:"poster"`
	Year     int           `json:"year" bson:"year"`
	Released bson.DateTime `json:"released" bson:"released"`
	Rated    string        `json:"rated" bson:"rated"`
	IMDB     Imdb          `json:"imdb" bson:"imdb"`
}

type Imdb struct {
	ID     int `json:"id" bson:"id"`
	Rating any `json:"rating" bson:"rating"`
	Votes  any `json:"votes" bson:"votes"`
}

type MovieFilter struct {
	Title    string `json:"title" bson:"title"`
	YearFrom int    `json:"year_from" bson:"year_from"`
	YearTo   int    `json:"year_to" bson:"year_to"`
	Page     int64  `json:"page" bson:"page"`
	Limit    int64  `json:"limit" bson:"limit"`
}
