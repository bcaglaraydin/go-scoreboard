package handlers

import (
	"net/http"

	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/bcaglaraydin/go-scoreboard/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.UserService.AddUserToLeaderboard(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.UserService.SaveUser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func (h UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("guid")
	user, err := h.UserService.GetUserFromUserID(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
