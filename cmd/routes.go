package main

import (
	"github.com/bcaglaraydin/go-scoreboard/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupLeaderboardRoute(app, "/leaderboard")
	setupUserRoute(app, "/user")
}

func setupLeaderboardRoute(app *fiber.App, prefix string) {
	leaderboardRoute := app.Group(prefix)
	leaderboardRoute.Get("/", handlers.GetLeaderboard)
}

func setupUserRoute(app *fiber.App, prefix string) {
	userRoute := app.Group(prefix)
	userRoute.Post("/create", handlers.CreateUser)
}
