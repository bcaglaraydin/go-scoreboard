package main

import (
	"github.com/bcaglaraydin/go-scoreboard/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.List)
	app.Post("/create", handlers.CreateUser)
}
