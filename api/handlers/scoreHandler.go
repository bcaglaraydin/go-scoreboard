package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
)

type ScoreHandler struct {
	ScoreService services.ScoreService
	UserService  services.UserService
}

func (h ScoreHandler) SubmitScore(c *fiber.Ctx) error {
	var score models.Score
	if err := c.BodyParser(&score); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if time.Unix(score.Timestamp, 0).Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You can't submit a score from past!"})
	}

	if score.ScoreWorth <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You can't submit a negative score!"})
	}

	newScore, err := h.ScoreService.UpdateUserScore(&score)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.UserService.GetUserFromUserID(score.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	err = h.UserService.SaveUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("User %s current score: %f", score.UserID, newScore)})
}
