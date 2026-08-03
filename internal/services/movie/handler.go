package movie

import (
	"github.com/gofiber/fiber/v3"
)

type MovieHandler struct {
	repo *MovieRepository
}

func NewMovieHandler(repo *MovieRepository) *MovieHandler {
	return &MovieHandler{
		repo: repo,
	}
}

// @Summary Get Movies
// @Description Get a list of movies with pagination
// @Tags Movies
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of movies per page" default(10)
// @Success 200 {array} Movie
// @Failure 500 {object} fiber.Map{"Error": "Query film list failed"}
// @Router /api/movies [get]

func (mh *MovieHandler) GetMovies(c fiber.Ctx) error {
	page := fiber.Query[int64](c, "page", 1)
	limit := fiber.Query[int64](c, "limit", 10)

	filter := MovieFilter{
		Page:  page,
		Limit: limit,
	}

	movies, err := mh.repo.GetMoviesCollection(filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": movies,
	})
}

// @Summary Search Movies
// @Description Search for movies based on filters
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {array} Movie
// @Failure 500 {object} fiber.Map{"Error": "Query film list failed"}
// @Router /api/movies/search [post]
func (mh *MovieHandler) SearchMovies(c fiber.Ctx) error {
	var request MovieFilter
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Invalid request body",
		})
	}

	movies, err := mh.repo.GetMoviesCollection(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": movies,
	})
}
