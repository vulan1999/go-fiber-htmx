package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/app/database"
	"github.com/vulan1999/todo-htmx/app/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetMovies(c fiber.Ctx) error {
	page := fiber.Query[int64](c, "page", 1)
	limit := fiber.Query[int64](c, "limit", 10)

	offset := (page - 1) * limit

	collection := database.DB.Collection("movies")
	queryOptions := options.Find().SetSkip(offset).SetLimit(limit).SetSort(bson.M{"year": -1})
	cursor, err := collection.Find(context.TODO(), bson.M{}, queryOptions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Query film list failed",
		})
	}

	defer cursor.Close(context.TODO())

	var movies []models.Movie
	if err := cursor.All(context.TODO(), &movies); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Error while reading data",
		})
	}
	return c.JSON(movies)
}

type MovieFilterRequest struct {
	Title    string `json:"title"`
	YearFrom int    `json:"year_from"`
	YearTo   int    `json:"year_to"`
}

func GetMoviesByFilter(c fiber.Ctx) error {
	collection := database.DB.Collection("movies")

	queryOptions := options.Find().SetLimit(20).SetSort(bson.M{"year": -1})

	filter := bson.M{
		"year": bson.M{"$gt": 2015},
	}

	cursor, err := collection.Find(context.TODO(), filter, queryOptions)

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.StatusInternalServerError,
		})
	}

	defer cursor.Close(context.TODO())

	results := []models.Movie{}

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Printf("Error: %v", err)
	}

	return c.JSON(results)
}
