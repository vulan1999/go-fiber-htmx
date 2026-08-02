package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/app/controllers"
	"github.com/vulan1999/todo-htmx/app/database"
	"github.com/vulan1999/todo-htmx/app/services"
)

func ApiMovieGroup(app *fiber.App) {
	movie := app.Group("/api/movies")
	movieService := services.NewMovieService(database.DB.Collection("movies"))
	movieController := controllers.NewMovieController(movieService)

	movie.Get("/", movieController.GetMovies)
	movie.Post("/search", movieController.SearchMovies)
}
