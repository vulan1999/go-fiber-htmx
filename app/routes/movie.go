package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/app/controllers"
)

func ApiMovieGroup(app *fiber.App) {
	movie := app.Group("/api/movies")

	movie.Get("/", controllers.GetMovies)
}
