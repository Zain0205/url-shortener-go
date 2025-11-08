package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func RateLimiter(maxRequests int, window time.Duration) fiber.Handler {
	requests := make(map[string]int)
	timestamps := make(map[string]time.Time)

	return func(c *fiber.Ctx) error {
		ip := c.IP()

		// reset window
		if t, ok := timestamps[ip]; !ok || time.Since(t) > window {
			timestamps[ip] = time.Now()
			requests[ip] = 0
		}

		requests[ip]++
		if requests[ip] > maxRequests {
			return c.Status(429).JSON(fiber.Map{
				"error": "rate limit exceeded",
			})
		}

		return c.Next()
	}
}
