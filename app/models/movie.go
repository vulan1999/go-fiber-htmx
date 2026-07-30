package models

import (
	"encoding/json"
	"regexp"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// helper function to clean byte to get clean int number.
// For Example: "year": "2012è" -> "year": "2012"
type CleanInt string

func (c CleanInt) MarshalJSON() ([]byte, error) {
	// replace all rune that not number to empty string
	reg := regexp.MustCompile("[^0-9]")
	cleaned := reg.ReplaceAllString(string(c), "")
	if cleaned == "" {
		return json.Marshal(0)
	}
	return json.Marshal(cleaned)
}

// Define type
type Movie struct {
	ID     bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string        `json:"title" bson:"title"`
	Plot   string        `json:"plot" bson:"plot"`
	Poster string        `json:"poster" bson:"poster"`
	Year   CleanInt      `json:"year" bson:"year"`
}
