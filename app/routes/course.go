package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ApiCourseGroup(app *fiber.App) {
	api := app.Group("/api/courses")

	courses := []Course{
		{Id: 1, Name: "Course1"},
		{Id: 2, Name: "Course2"},
		{Id: 3, Name: "Course3"},
	}

	// ----- Get Path ------
	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(courses)
	})

	api.Get("/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Id")
		}

		for _, course := range courses {
			if course.Id == id {
				return c.JSON(course)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Course Id not Existed")
	})

	// ---  Post path ---
	api.Post("/", func(c fiber.Ctx) error {
		var course Course

		if err := c.Bind().Body(&course); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Body")
		}

		course.Id = len(courses) + 1
		courses = append(courses, course)

		return c.Status(fiber.StatusCreated).SendString("Course Created")
	})

	// --- Put Path ---
	api.Put("/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Invalid Id")
		}

		updateCourse := new(Course)

		for i := range courses {
			if courses[i].Id == id {
				updateCourse = &courses[i]
				break
			}
		}

		if updateCourse == nil {
			return c.Status(fiber.StatusNotFound).SendString("Course id not found")
		}

		var input struct {
			Name string `json:"name"`
		}

		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		updateCourse.Name = input.Name

		return c.JSON(updateCourse)
	})

	// -- Delete path --
	api.Delete("/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Invalid Id")
		}

		for i, course := range courses {
			if course.Id == id {
				courses = append(courses[:i], courses[i+1:]...)
				return c.Status(fiber.StatusOK).SendString("Course Deleted")
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("Course Id does not existed")
	})
}
