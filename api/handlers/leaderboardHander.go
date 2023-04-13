package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
)

type LeaderBoardHandler struct {
	UserService services.UserService
}

func (h LeaderBoardHandler) GetLeaderboard(c *fiber.Ctx) error {
	leaderboardJson, err := h.UserService.GetLeaderBoard()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(leaderboardJson)
}

func (h LeaderBoardHandler) GetLeaderboardFilterCountry(c *fiber.Ctx) error {
	country := c.Params("country_iso_code")
	leaderboardJson, err := h.UserService.GetLeaderBoard(country)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(leaderboardJson)
}
