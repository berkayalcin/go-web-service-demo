package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Logging(c *fiber.Ctx) error {
	requestBody := string(c.Request().Body())
	start := time.Now()
	err := c.Next()
	responseBody := string(c.Response().Body())
	if err != nil {
		return err
	}

	end := time.Now()
	diff := end.Sub(start).Milliseconds()

	go func(ctx fiber.Ctx, diff int64) {

		contextLogger := log.WithFields(log.Fields{
			"statusCode":   ctx.Response().StatusCode(),
			"path":         ctx.Path(),
			"method":       ctx.Method(),
			"elapsedTime":  fmt.Sprintf("%dms", diff),
			"requestBody":  requestBody,
			"responseBody": responseBody,
		})

		contextLogger.Info("Requests")
	}(*c, diff)

	return nil
}
