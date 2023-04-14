package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		e := helpers.ResponseError(fiber.StatusUnprocessableEntity, err.Error())
		return c.JSON(e)
	}

	if user.UserID == uuid.Nil {
		e := helpers.ResponseError(fiber.StatusBadRequest, "You must provide a valid user id!")
		return c.JSON(e)
	}

	if user.DisplayName == "" {
		e := helpers.ResponseError(fiber.StatusBadRequest, "You must provide a valid name!")
		return c.JSON(e)
	}
	if user.Points < 0 {
		e := helpers.ResponseError(fiber.StatusBadRequest, "You can't submit a negative score!")
		return c.JSON(e)
	}

	if err := h.UserService.AddUserToLeaderboard(user); err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)

	}

	if err := h.UserService.SaveUser(user); err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("guid")
	user, err := h.UserService.GetUserFromUserID(userID)
	if err != nil {
		e := helpers.ResponseError(fiber.StatusInternalServerError, err.Error())
		return c.JSON(e)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
