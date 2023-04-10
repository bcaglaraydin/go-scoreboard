package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)

	return c.Status(200).JSON(users)

}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}
