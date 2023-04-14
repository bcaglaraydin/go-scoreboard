package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	helpers.HTTPError
//	@Failure		422	{object}	helpers.HTTPError
//	@Failure		500	{object}	helpers.HTTPError
//	@Router			/users/create [post]

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		helpers.ResponseError(c, fiber.StatusUnprocessableEntity, err.Error())
	}
	if user.Points < 0 {
		helpers.ResponseError(c, fiber.StatusBadRequest, "You can't submit a negative score!")
	}

	if err := h.UserService.AddUserToLeaderboard(user); err != nil {
		helpers.ResponseError(c, fiber.StatusInternalServerError, err.Error())

	}

	if err := h.UserService.SaveUser(user); err != nil {
		helpers.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("guid")
	user, err := h.UserService.GetUserFromUserID(userID)
	if err != nil {
		helpers.ResponseError(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
