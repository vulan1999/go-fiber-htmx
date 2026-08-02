package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/app/models"
	"github.com/vulan1999/todo-htmx/app/services"
)

type MovieController struct {
	service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{
		service: service,
	}
}

// @Summary Get Movies
// @Description Get a list of movies with pagination
// @Tags Movies
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of movies per page" default(10)
// @Success 200 {array} models.Movie
// @Failure 500 {object} fiber.Map{"Error": "Query film list failed"}
// @Router /api/movies [get]

func (mc *MovieController) GetMovies(c fiber.Ctx) error {
	page := fiber.Query[int64](c, "page", 1)
	limit := fiber.Query[int64](c, "limit", 10)

	filter := models.MovieFilter{
		Page:  page,
		Limit: limit,
	}

	movies, err := mc.service.GetMoviesCollection(filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(movies)
}

// @Summary Search Movies
// @Description Search for movies based on filters
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {array} models.Movie
// @Failure 500 {object} fiber.Map{"Error": "Query film list failed"}
// @Router /api/movies/search [post]
func (mc *MovieController) SearchMovies(c fiber.Ctx) error {
	var request models.MovieFilter
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid request body",
		})
	}

	movies, err := mc.service.GetMoviesCollection(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(movies)
}
