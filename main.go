package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vikbert/go-fiber-api/database"
	"github.com/vikbert/go-fiber-api/router"
	"log"
)

func main() {
	database.ConnectDbInstance()

	app := fiber.New()
	router.SetupDefaultRoutes(app)
	router.SetupApiRoutes(app)

	log.Fatal(app.Listen("localhost:3001"))
}
