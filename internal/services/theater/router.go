package theater

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/database"
)

func ApiTheaterGroup(app *fiber.App) {
	theater := app.Group("/api/theaters")
	theaterRepository := NewTheaterRepository(database.DB.Collection("theaters"))
	theaterHandler := NewTheaterHandler(theaterRepository)

	theater.Get("/", theaterHandler.GetTheaters)
}
