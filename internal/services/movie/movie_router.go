package movie

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/database"
)

func ApiMovieGroup(app *fiber.App) {
	movie := app.Group("/api/movies")
	movieService := NewMovieService(database.DB.Collection("movies"))
	movieController := NewMovieController(movieService)

	movie.Get("/", movieController.GetMovies)
	movie.Post("/search", movieController.SearchMovies)
}
