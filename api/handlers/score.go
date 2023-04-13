package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/handlers/common"
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/gofiber/fiber/v2"
)

type Score struct {
	ScoreWorth float64 `json:"score_worth"`
	UserID     string  `json:"user_id"`
	Timestamp  int64   `json:"timestamp"`
}

func SubmitScore(c *fiber.Ctx) error {
	var score Score
	if err := c.BodyParser(&score); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	if time.Unix(score.Timestamp, 0).Before(time.Now()) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You can't submit a score from past!"})
	}

	if score.ScoreWorth <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You can't submit a negative score!"})
	}

	rdb := database.GetRedisClient()
	newScore := rdb.ZIncrBy(database.Ctx, helpers.RedisLeaderboardKey, score.ScoreWorth, score.UserID)
	user, err := common.GetUserFromUserID(rdb, score.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	common.UpdateUserPointAndScore(rdb, user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fmt.Sprintf("User %s current score: %s", score.UserID, newScore)})
}
