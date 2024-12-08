package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL) // Ensure routes package is implemented
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("APP_PORT environment variable not set")
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Printf("Starting server on port %s...", port)
	log.Fatal(app.Listen(port))
}
