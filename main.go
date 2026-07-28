package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
	"github.com/vulan1999/todo-htmx/app/routes"
)

func main() {
	engine := html.New("./app/views", ".html")
	engine.Reload(true)
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Fiber Template Example",
			"PageName": "Home",
		}, "layouts/main")
	})

	routes.ApiGroup(app)

	log.Fatal(app.Listen(":3000"))
}
