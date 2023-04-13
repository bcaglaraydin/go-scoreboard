package handlers

import (
	"net/http"

	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/handlers/common"
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	rdb := database.GetRedisClient()

	if err := addUserToLeaderboard(rdb, user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := common.UpdateUserPointAndScore(rdb, user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	rdb := database.GetRedisClient()
	userID := c.Params("guid")
	user, err := common.GetUserFromUserID(rdb, userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func addUserToLeaderboard(rdb *redis.Client, user *models.User) error {
	return rdb.ZAdd(database.Ctx, helpers.RedisLeaderboardKey, &redis.Z{
		Score:  float64(user.Points),
		Member: user.UserID.String(),
	}).Err()
}
