package main

import "github.com/gofiber/fiber/v2"

func main() {
	var app *fiber.App = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
