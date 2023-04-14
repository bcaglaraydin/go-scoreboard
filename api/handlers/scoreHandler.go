package handlers

import (
	"fmt"
	"time"

	"github.com/bcaglaraydin/go-scoreboard/helpers"
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
		e := helpers.ResponseError(fiber.StatusUnprocessableEntity, err.Error())
		return c.JSON(e)
	}

	if time.Unix(score.Timestamp, 0).Before(time.Now()) {
		e := helpers.ResponseError(fiber.StatusBadRequest, "You can't submit a score from past!")
		return c.JSON(e)
	}

	if score.ScoreWorth < 0 {
		e := helpers.ResponseError(fiber.StatusBadRequest, "You can't submit a negative score!")
		return c.JSON(e)
	}

	newScore, err := h.ScoreService.UpdateUserScore(&score)
	if err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)
	}
	user, err := h.UserService.GetUserFromUserID(score.UserID)
	if err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)
	}

	if err := h.UserService.SaveUser(user); err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("User %s current score: %f", score.UserID, newScore)})
}
