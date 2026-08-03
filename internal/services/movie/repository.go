package movie

import (
	"context"

	"github.com/vulan1999/todo-htmx/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MovieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(collection *mongo.Collection) *MovieRepository {
	return &MovieRepository{
		collection: database.DB.Collection("movies"),
	}
}

func (m *MovieRepository) GetMoviesCollection(filter MovieFilter) ([]Movie, error) {
	query := bson.M{}

	if filter.Title != "" {
		query["title"] = bson.M{"$regex": filter.Title, "$options": "i"}
	}

	if filter.YearFrom != 0 && filter.YearTo != 0 {
		query["year"] = bson.M{"$gte": filter.YearFrom, "$lte": filter.YearTo}
	} else if filter.YearFrom != 0 {
		query["year"] = bson.M{"$gte": filter.YearFrom}
	} else if filter.YearTo != 0 {
		query["year"] = bson.M{"$lte": filter.YearTo}
	}

	offset := (filter.Page - 1) * filter.Limit

	queryOptions := options.Find().SetSkip(offset).SetLimit(filter.Limit).SetSort(bson.M{"year": -1})

	cursor, err := m.collection.Find(context.TODO(), query, queryOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	movies := []Movie{}
	if err := cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
