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

	collection := database.DB.Collection("movies")

	queryOptions := options.Find().SetLimit(20).SetSort(bson.M{"year": -1})

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
