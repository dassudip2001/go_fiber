package main

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Server is healthy  👋!")
}

func main() {

	// initialized the database connection
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", welcome)

	// load the static files
	app.Static("/", "./public")

	app.Listen(":3000")
}
