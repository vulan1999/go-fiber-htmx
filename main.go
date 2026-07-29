package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/template/html/v3"
	"github.com/vulan1999/todo-htmx/app/middleware"
	"github.com/vulan1999/todo-htmx/app/routes"
)

func main() {
	// HTML template engine
	engine := html.New("./app/views", ".html")
	engine.Reload(true)
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	// Middleware used
	app.Use(requestid.New())
	app.Use(middleware.StructuredLogger())

	// Api group
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Fiber Template Example",
			"PageName": "Home",
		}, "layouts/main")
	})

	routes.ApiGroup(app)

	//Run
	log.Fatal(app.Listen(":3000"))
}
