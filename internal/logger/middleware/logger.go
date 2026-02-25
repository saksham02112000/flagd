package middleware

import (
	"flagd/internal/logger"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
)


func RequestLogger(log *slog.Logger) fiber.Handler{
	return func(c fiber.Ctx) error {
		start := time.Now()

		ctx := logger.WithContext(c.Context(), log)
		c.SetContext(ctx)
		err := c.Next()
		latency := time.Since(start)
		
		status := c.Response().StatusCode()

        switch {
			case status >= 500:
				log.ErrorContext(ctx, "request",
					"method",   c.Method(),
					"path",     c.Path(),
					"status",   status,
					"duration", latency,
					"ip",       c.IP(),
				)
			case status >= 400:
				log.WarnContext(ctx, "request",
					"method",   c.Method(),
					"path",     c.Path(),
					"status",   status,
					"duration", latency,
					"ip",       c.IP(),
				)
			default:
				log.InfoContext(ctx, "request",
					"method",   c.Method(),
					"path",     c.Path(),
					"status",   status,
					"duration", latency,
				)
        }

        return err
    }
}