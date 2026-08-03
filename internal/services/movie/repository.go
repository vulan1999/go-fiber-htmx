package movie

import (
	"context"
	"log"

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

func (m *MovieRepository) CreateMovie(movie Movie) error {
	cursor, err := m.collection.InsertOne(context.TODO(), movie)

	if err != nil {
		return err
	}

	log.Printf("New Movie insert to db with _id: %v", cursor.InsertedID)

	return nil

}

func (m *MovieRepository) UpdateMovie(updateMovieRequest MovieUpdateRequest, id bson.ObjectID) error {
	filter := bson.M{"_id": id}

	updateQuery := bson.M{}

	if updateMovieRequest.Title != "" {
		updateQuery["title"] = updateMovieRequest.Title
	}

	if updateMovieRequest.Plot != "" {
		updateQuery["plot"] = updateMovieRequest.Plot
	}

	if updateMovieRequest.Poster != "" {
		updateQuery["poster"] = updateMovieRequest.Poster
	}

	if updateMovieRequest.Rated != "" {
		updateQuery["rated"] = updateMovieRequest.Rated
	}

	if updateMovieRequest.Year > 0 {
		updateQuery["year"] = updateMovieRequest.Year
	}

	update := bson.M{
		"$set": updateQuery,
	}

	cursor, err := m.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	log.Printf("Matched Records: %d", cursor.MatchedCount)
	log.Printf("Modified Records: %d", cursor.ModifiedCount)

	return nil
}
