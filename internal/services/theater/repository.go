package theater

import (
	"context"

	"github.com/vulan1999/todo-htmx/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TheaterRepository struct {
	collection *mongo.Collection
}

func NewTheaterRepository(collection *mongo.Collection) *TheaterRepository {
	return &TheaterRepository{
		collection: database.DB.Collection("theaters"),
	}
}

func (t *TheaterRepository) GetTheatersCollection() ([]Theater, error) {
	queryOptions := options.Find().SetSort(bson.M{"theaterId": -1})

	cursor, err := t.collection.Find(context.TODO(), bson.M{}, queryOptions)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	theaters := []Theater{}

	if err := cursor.All(context.TODO(), &theaters); err != nil {
		return nil, err
	}

	return theaters, nil
}
