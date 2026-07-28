package routes

import "github.com/gofiber/fiber/v3"

type Course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ApiGroup(app *fiber.App) {
	api := app.Group("/api")

	course := []Course{
		{Id: 1, Name: "Course1"},
		{Id: 2, Name: "Course2"},
		{Id: 3, Name: "Course3"},
	}

	api.Get("/courses", func(c fiber.Ctx) error {
		return c.JSON(course)
	})
}
