package theater

import "go.mongodb.org/mongo-driver/v2/bson"

type Address struct {
	Street1 string `json:"street1" bson:"street1"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	ZipCode string `json:"zip_code" bson:"zipcode"`
}

type Geo struct {
	Coordinates []float32 `json:"coordinates" bson:"coordinates"`
}

type Location struct {
	Address Address `json:"address" bson:"address"`
	Geo     Geo     `json:"geo" bson:"geo"`
}

type Theater struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TheaterID int           `json:"theater_id" bson:"theaterId"`
	Location  any           `json:"locatiton" bson:"location"`
}
