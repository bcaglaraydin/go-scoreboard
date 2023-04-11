package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	// user := new(models.User)
	// if err := c.BodyParser(user); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	// database.DB.Db.Create(&user)
	// return c.Status(200).JSON(user)

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Connect to Redis
	rdb := database.GetRedisClient()

	// Calculate the user's rank based on their score
	rank, err := rdb.ZCount(database.Ctx, "leaderboard", "-inf", "+inf").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Add the user to the leaderboard
	if err := rdb.ZAdd(database.Ctx, "leaderboard", &redis.Z{
		Score:  float64(user.Points),
		Member: user.UserID.String(),
	}).Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update the user's rank in the database
	user.Rank = int(rank) + 1
	if err := database.DB.Db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Return the updated user as JSON
	return c.Status(fiber.StatusOK).JSON(user)
}
