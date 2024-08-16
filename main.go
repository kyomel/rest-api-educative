package main

import (
	"fmt"
	"os"
	"rest-api-educative/database"
	"rest-api-educative/routes"
	"rest-api-educative/utils"

	"github.com/gofiber/fiber/v2"
)

const DEFAULT_PORT = "8080"

func NewFiberApp() *fiber.App {
	var app *fiber.App = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(app)
	return app
}

func main() {
	var app *fiber.App = NewFiberApp()

	database.InitDatabase(utils.GetValue("DB_NAME"))

	var PORT string = os.Getenv("PORT")

	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))
}
