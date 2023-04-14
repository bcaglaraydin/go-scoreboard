package main

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/handlers"
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	rdb := database.GetRedisClient()
	leaderboardHandler := handlers.LeaderBoardHandler{UserService: services.NewUserServiceDB(rdb)}
	userHandler := handlers.UserHandler{UserService: services.NewUserServiceDB(rdb)}
	scoreHandler := handlers.ScoreHandler{UserService: services.NewUserServiceDB(rdb), ScoreService: services.NewScoreServiceDB(rdb)}
	setupLeaderboardRoute(app, "/leaderboard", leaderboardHandler)
	setupUserRoute(app, "/user", userHandler)
	setupScoreRoute(app, "/score", scoreHandler)
}

func setupLeaderboardRoute(app *fiber.App, prefix string, handler handlers.LeaderBoardHandler) {
	leaderboardRoute := app.Group(prefix)
	leaderboardRoute.Get("/", handler.GetLeaderboard)
	leaderboardRoute.Get("/:country_iso_code", handler.GetLeaderboardFilterCountry)
}

func setupUserRoute(app *fiber.App, prefix string, handler handlers.UserHandler) {
	userRoute := app.Group(prefix)
	userRoute.Post("/create", handler.CreateUser)
	userRoute.Get("/profile/:guid", handler.GetUser)
}

func setupScoreRoute(app *fiber.App, prefix string, handler handlers.ScoreHandler) {
	scoreRoute := app.Group(prefix)
	scoreRoute.Post("/submit", handler.SubmitScore)
}
