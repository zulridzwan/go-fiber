package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)

	app.Static("/", "./public") 

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}