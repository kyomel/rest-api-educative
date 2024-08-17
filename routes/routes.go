package routes

import (
	"rest-api-educative/handlers"
	"rest-api-educative/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	var publicRoutes fiber.Router = app.Group("/api/v1")
	publicRoutes.Post("/api/v1/signup", handlers.Signup)
	publicRoutes.Post("/api/v1/login", handlers.Login)
	publicRoutes.Get("/api/v1/items", handlers.GetAllItems)
	publicRoutes.Get("/api/v1/items/:id", handlers.GetItemById)

	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())
	privateRoutes.Post("/api/v1/items", handlers.CreateItem)
	privateRoutes.Put("/api/v1/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/api/v1/items/:id", handlers.DeleteItem)
}
