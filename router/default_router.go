package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupDefaultRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("it works")
	})
}
