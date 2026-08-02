package services

import (
	"context"

	"github.com/vulan1999/todo-htmx/app/database"
	"github.com/vulan1999/todo-htmx/app/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MovieService struct {
	collection *mongo.Collection
}

func NewMovieService(collection *mongo.Collection) *MovieService {
	return &MovieService{
		collection: database.DB.Collection("movies"),
	}
}

func (s *MovieService) GetMoviesCollection(filter models.MovieFilter) ([]models.Movie, error) {
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

	cursor, err := s.collection.Find(context.TODO(), query, queryOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	movies := []models.Movie{}
	if err := cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
