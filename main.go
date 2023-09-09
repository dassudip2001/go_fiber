package main

import (
	"github.com/dassudip2001/webapp/database"
	"github.com/dassudip2001/webapp/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Server is healthy  👋!")
}

func setUpRooutes(app *fiber.App) {
	// server health check
	app.Get("/", welcome)
	// user end points
	app.Post("/api/v1/user", routes.CreateUser)
	app.Get("/api/v1/user", routes.GetUsers)
	app.Get("/api/v1/user/:id", routes.GetUser)
	app.Put("/api/v1/user/:id", routes.UpdateUser)
	app.Delete("/api/v1/user/:id", routes.DeleteUser)
}

func main() {

	// initialized the database connection
	database.ConnectDb()

	app := fiber.New()

	// routes setup
	setUpRooutes(app)

	// load the static files
	app.Static("/", "./public")

	app.Listen(":3000")
}
