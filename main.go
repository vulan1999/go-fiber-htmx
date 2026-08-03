package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/template/html/v3"
	"github.com/joho/godotenv"
	"github.com/vulan1999/todo-htmx/internal/database"
	"github.com/vulan1999/todo-htmx/internal/helpers"
	"github.com/vulan1999/todo-htmx/internal/middleware"
	"github.com/vulan1999/todo-htmx/internal/services/movie"
	"github.com/vulan1999/todo-htmx/internal/services/theater"
)

func init() {
	// Load enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot find .env file, relying on system enviroment variables")
	} else {
		log.Println(".env file found")
	}
}

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

	// Connect to Mongo Db
	mongoUri := helpers.GetEnv("MONGODB_URI", "mongodb://localhost:27017")
	mongoDbName := helpers.GetEnv("MONGODB_NAME", "sample_mflix")

	if err := database.ConnectMongoDB(mongoUri, mongoDbName); err != nil {
		log.Fatalf("Connect to MongoDb failed: %v", err)
	}

	// Api group
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":    "Fiber Template Example",
			"PageName": "Home",
		}, "layouts/main")
	})

	movie.ApiMovieGroup(app)
	theater.ApiTheaterGroup(app)

	//Run
	log.Fatal(app.Listen(":3000"))
}
