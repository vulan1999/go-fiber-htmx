package movie

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/database"
)

func ApiMovieGroup(app *fiber.App) {
	movie := app.Group("/api/movies")
	movieRepository := NewMovieRepository(database.DB.Collection("movies"))
	movieHandler := NewMovieHandler(movieRepository)

	movie.Get("/", movieHandler.GetMovies)
	movie.Post("/search", movieHandler.SearchMovies)
	movie.Post("/create", movieHandler.CreateMovie)
	movie.Put("/update/:id", movieHandler.UpdateMovie)
}
