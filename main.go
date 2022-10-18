package main

import (
	"Papatia/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/trendyol", handler.Trendyol)

	app.Listen(":3000")
}
