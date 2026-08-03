package theater

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/vulan1999/todo-htmx/internal/helpers"
)

type TheaterHandler struct {
	repo *TheaterRepository
}

func NewTheaterHandler(repo *TheaterRepository) *TheaterHandler {
	return &TheaterHandler{
		repo: repo,
	}
}

// @Summary Get Theaters
// @Description Get a list of theaters with pagination
// @Tags Theaters
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of theaters per page" default(10)
// @Success 200 {object} helpers.ApiResponse
// @Failure 500 {object} helpers.ApiResponse
// @Router /api/theaters [get]
func (th *TheaterHandler) GetTheaters(c fiber.Ctx) error {
	theaters, err := th.repo.GetTheatersCollection()

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(helpers.ApiResponse{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return c.JSON(helpers.ApiResponse{
		Message: "Get theaters list successfully",
		Data:    theaters,
	})
}
