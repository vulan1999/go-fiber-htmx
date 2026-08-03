package movie

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/database"
)

func ApiMovieGroup(app *fiber.App) {
	movie := app.Group("/api/movies")
	movieRepository := NewMovieRepository(database.DB.Collection("movies"))
	movieController := NewMovieHandler(movieRepository)

	movie.Get("/", movieController.GetMovies)
	movie.Post("/search", movieController.SearchMovies)
}
