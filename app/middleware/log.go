package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func StructuredLogger() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()

		requestId := requestid.FromContext(c)

		slog.Info("HTTP Request",
			slog.String("requestId", requestId),
			slog.Int("status", status),
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.Duration("latency", latency),
		)

		return err
	}
}
