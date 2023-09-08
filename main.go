package main

import (
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Server is healthy  👋!")
}

func main() {
	// database.ConnectDb()
	app := fiber.New()

	app.Get("/", welcome)

	app.Listen(":3000")
}

// how to  connect mysql in golang gorm
// https://www.youtube.com/watch?v=ZzrV7o6t6fE
//
