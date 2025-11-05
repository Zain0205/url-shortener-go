package handler

import (
	"net/http"

	"github.com/Zain0205/url-shortener-go/app/service"
	"github.com/gofiber/fiber/v2"
)

func Shorten(c *fiber.Ctx) error {
	type Req struct {
		URL string `json:"url"`
	}
	var body Req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	code, err := service.ShortenURL(body.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"short_url": "http://localhost:3000/" + code,
	})
}

func Resolve(c *fiber.Ctx) error {
	code := c.Params("code")

	original, err := service.ResolveURL(code)
	if err != nil {
		return c.Status(404).SendString("URL not found")
	}

	return c.Redirect(original, http.StatusMovedPermanently)
}
