package theater

import (
	"log"

	"github.com/gofiber/fiber/v3"
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
// @Success 200 {array} Theater
// @Failure 500 {object} fiber.Map{"Error": "Query film list failed"}
// @Router /api/theaters [get]
func (th *TheaterHandler) GetTheaters(c fiber.Ctx) error {
	theaters, err := th.repo.GetTheatersCollection()

	if err != nil {
		log.Printf("Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON((fiber.Map{
			"Error": "Internal Server Error",
		}))
	}

	return c.JSON(fiber.Map{
		"data": theaters,
	})
}
