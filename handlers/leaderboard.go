package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/gofiber/fiber/v2"
)

func GetLeaderboard(c *fiber.Ctx) error {

	rdb := database.GetRedisClient()

	// Get the leaderboard as a slice of user IDs, sorted by rank
	result, err := rdb.ZRange(database.Ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Get the rank and score for each user, and return them in a slice of User
	users := make([]models.User, len(result))
	for i, userID := range result {
		rank, err := rdb.ZRank(database.Ctx, "leaderboard", userID).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		score, err := rdb.ZScore(database.Ctx, "leaderboard", userID).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Query the user data from the database
		var user models.User
		if err := database.DB.Db.Where("user_id = ?", userID).First(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Set the user's rank and score
		user.Rank = int(rank) + 1
		user.Points = int(score)

		users[i] = user
	}

	// Return the leaderboard as JSON
	return c.Status(fiber.StatusOK).JSON(users)

}
