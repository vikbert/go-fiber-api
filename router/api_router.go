package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-fiber-api/controller"
)

func SetupApiRoutes(app *fiber.App) {

	api := app.Group("/api")

	userCtrl := controller.NewUserController()
	api.Post("/users", userCtrl.Create)
	api.Get("/users", userCtrl.List)
	api.Get("/users/:id", userCtrl.Read)
	api.Put("/users/:id", userCtrl.Update)
	api.Delete("/users/:id", userCtrl.Delete)

	productCtrl := controller.NewProductController()
	api.Post("/products", productCtrl.Create)
	api.Get("/products", productCtrl.List)
	api.Get("/products/:id", productCtrl.Read)
	api.Put("/products/:id", productCtrl.Update)
	api.Delete("/products/:id", productCtrl.Delete)
}
