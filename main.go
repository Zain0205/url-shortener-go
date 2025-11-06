package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/Zain0205/url-shortener-go/app/db"
	"github.com/Zain0205/url-shortener-go/app/handler"
	"github.com/Zain0205/url-shortener-go/app/middleware"
)

func main() {
	godotenv.Load()
	db.Init()

	app := fiber.New()

	app.Use(middleware.RateLimiter(5, 10*time.Second)) // 5 req per 10s

	app.Post("/shorten", handler.Shorten)
	app.Get("/:code", handler.Resolve)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
