package movie

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
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
// @Success 200 {object} helpers.ApiResponse
// @Failure 500 {object} helpers.ApiResponse
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
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Get movies failed",
			Data:    nil,
		})
	}

	return c.JSON(helpers.ApiResponse{
		Message: "Get movies successfully",
		Data:    movies,
	})
}

// @Summary Search Movies
// @Description Search for movies based on filters
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {object} helpers.ApiResponse
// @Failure 500 {object} helpers.ApiResponse
// @Router /api/movies/search [post]
func (mh *MovieHandler) SearchMovies(c fiber.Ctx) error {
	var request MovieFilter
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.ApiResponse{
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	movies, err := mh.repo.GetMoviesCollection(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Search movies failed",
			Data:    nil,
		})
	}

	return c.JSON(helpers.ApiResponse{
		Message: "Search movies successfully",
		Data:    movies,
	})
}

// @Summary Create Movie
// @Description Create Movie Record on movie collection
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {object} helpers.ApiResponse
// @Failure 500 {object} helpers.ApiResponse
// @Router /api/movies/create [post]
func (mh *MovieHandler) CreateMovie(c fiber.Ctx) error {
	var request Movie
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.ApiResponse{
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	if err := mh.repo.CreateMovie(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Error while creating movie",
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(helpers.ApiResponse{
		Message: "Move Created",
		Data:    nil,
	})
}

// @Summary Update Movie
// @Description Update Movie with object id
// @Tags Movies
// @Accept json
// @Produce json
// @Success 200 {object} helpers.ApiResponse
// @Failure 500 {object} helpers.ApiResponse
// @Router /api/movies/update/:id [put]
func (mh *MovieHandler) UpdateMovie(c fiber.Ctx) error {
	id := fiber.Params[string](c, "id")
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Error while converting Object Id",
			Data:    nil,
		})
	}

	var request MovieUpdateRequest

	if err := c.Bind().JSON(&request); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(helpers.ApiResponse{
			Message: "Bad Request",
			Data:    nil,
		})
	}

	if err := mh.repo.UpdateMovie(request, objectId); err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Error while updating record",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(helpers.ApiResponse{
		Message: "Update record successfully",
		Data:    nil,
	})
}
