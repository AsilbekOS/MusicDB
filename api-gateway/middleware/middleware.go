package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

func FiberCORSMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Set the CORS headers to allow any origin to access resources
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the method is OPTIONS, return a status OK response
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}

		// Otherwise, pass the request to the next handler
		return next(c)
	}
}

func FiberRateLimiter(limiter *rate.Limiter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the request is allowed based on the rate limiter
		if !limiter.Allow() {
			// If the request exceeds the rate limit, return a 429 (Too Many Requests) status
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		}

		// Otherwise, pass the request to the next handler
		return c.Next()
	}
}
