package handlers

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/handlers/common"
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetLeaderboard(c *fiber.Ctx) error {

	rdb := database.GetRedisClient()

	leaderboardJson, err := getLeaderBoardFromDb(rdb)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(leaderboardJson)
}

func GetLeaderboardFilterCountry(c *fiber.Ctx) error {

	rdb := database.GetRedisClient()
	country := c.Params("country_iso_code")
	leaderboardJson, err := getLeaderBoardFromDb(rdb, country)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(leaderboardJson)
}

func getLeaderBoardFromDb(rdb *redis.Client, args ...string) ([]*models.User, error) {
	userIDs, err := rdb.ZRevRange(database.Ctx, helpers.RedisLeaderboardKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0)
	for _, userID := range userIDs {
		user, err := common.GetUserFromUserID(rdb, userID)
		if err != nil {
			return nil, err
		}
		if len(args) > 0 {
			country := args[0]
			if country != "" && user.Country != country {
				continue
			}
		}

		users = append(users, user)
	}
	return users, nil
}
