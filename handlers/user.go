package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bcaglaraydin/go-scoreboard/database"
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

	if err := updateUserRank(rdb, user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func addUserToLeaderboard(rdb *redis.Client, user *models.User) error {
	return rdb.ZAdd(database.Ctx, helpers.RedisLeaderboardKey, &redis.Z{
		Score:  float64(user.Points),
		Member: user.UserID.String(),
	}).Err()
}

func updateUserRank(rdb *redis.Client, user *models.User) error {
	rank, err := rdb.ZRevRank(database.Ctx, helpers.RedisLeaderboardKey, user.UserID.String()).Result()
	if err != nil {
		return err
	}
	user.Rank = int(rank) + 1

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// err = rdb.HSet(database.Ctx, helpers.RedisUsersKey, user.UserID.String(), map[string]interface{}{
	// 	"user_id":      user.UserID.String(),
	// 	"display_name": user.DisplayName,
	// 	"points":       user.Points,
	// 	"rank":         user.Rank,
	// 	"country":      user.Country,
	// }).Err()
	// return err

	err = rdb.HSet(database.Ctx, helpers.RedisUsersKey, user.UserID.String(), userBytes).Err()
	return err
}

// func CreateUser(c *fiber.Ctx) error {
// 	// user := new(models.User)
// 	// if err := c.BodyParser(user); err != nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	// 		"message": err.Error(),
// 	// 	})
// 	// }

// 	// database.DB.Db.Create(&user)
// 	// return c.Status(200).JSON(user)

// 	// user := new(models.User)
// 	// if err := c.BodyParser(user); err != nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	// 		"message": err.Error(),
// 	// 	})
// 	// }

// 	// // Connect to Redis
// 	// rdb := database.GetRedisClient()

// 	// // Calculate the user's rank based on their score
// 	// rank, err := rdb.ZRevRank(database.Ctx, "leaderboard", user.UserID.String()).Result()
// 	// if err != nil && err != redis.Nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	// 		"message": err.Error(),
// 	// 	})
// 	// }

// 	// // Add the user to the leaderboard
// 	// if err := rdb.ZAdd(database.Ctx, "leaderboard", &redis.Z{
// 	// 	Score:  float64(user.Points),
// 	// 	Member: user.UserID.String(),
// 	// }).Err(); err != nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	// 		"message": err.Error(),
// 	// 	})
// 	// }

// 	// // Update the user's rank in the database
// 	// user.Rank = int(rank) + 1
// 	// if err := database.DB.Db.Save(&user).Error; err != nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	// 		"message": err.Error(),
// 	// 	})
// 	// }

// 	// // Return the updated user as JSON
// 	// return c.Status(fiber.StatusOK).JSON(user)

// 	user := new(models.User)
// 	if err := c.BodyParser(user); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	userBytes, err := json.Marshal(user)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	// Add user to Redis sorted set based on their points
// 	rdb := database.GetRedisClient()
// 	if err := rdb.ZAdd(database.Ctx, "leaderboard", &redis.Z{
// 		Score:  float64(user.Points),
// 		Member: string(userBytes),
// 	}).Err(); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	// Update user rank in Redis sorted set
// 	rank, err := rdb.ZRevRank(database.Ctx, "users", user.UserID.String()).Result()
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	user.Rank = int(rank) + 1
// 	if err := rdb.HSet(database.Ctx, user.UserID.String(), map[string]interface{}{
// 		"user_id":      user.UserID.String(),
// 		"display_name": user.DisplayName,
// 		"points":       user.Points,
// 		"rank":         user.Rank,
// 		"country":      user.Country,
// 	}).Err(); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	// // Return success response with created user information
// 	return c.Status(http.StatusCreated).JSON(user)
// }

// func CreateRandomUsers(c *fiber.Ctx) error {
// 	// Parse the number of users to create from the request body
// 	var requestBody struct {
// 		NumUsers int `json:"num_users"`
// 	}
// 	if err := c.BodyParser(&requestBody); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "invalid request body",
// 		})
// 	}
// 	numUsers := requestBody.NumUsers

// 	mockCountries := [5]string{"us", "tr", "bd", "gb", "fr"}
// 	// Create the specified number of fake user
// 	for i := 0; i < numUsers; i++ {
// 		user := models.User{
// 			UserID:      uuid.MustParse(faker.UUIDHyphenated()),
// 			DisplayName: faker.Name(),
// 			Points:      rand.Intn(100),
// 			Country:     mockCountries[rand.Intn(len(mockCountries))],
// 			Rank:        rand.Intn(100),
// 		}

// 		// Create a JSON request body
// 		reqBody, err := json.Marshal(user)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 			})
// 		}

// 		// Send the request using fiber client
// 		r, err := http.NewRequest("POST", c.BaseURL()+"/user/create", bytes.NewBuffer(reqBody))
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"message": err.Error(),
// 			})
// 		}

// 		r.Header.Add("Content-Type", "application/json")
// 		client := &http.Client{}
// 		res, err := client.Do(r)
// 		if err != nil {
// 			panic(err)
// 		}

// 		defer res.Body.Close()

// 		// // Check if the request was successful
// 		// if res.StatusCode != http.StatusOK {
// 		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		// 		"message": fmt.Sprintf("failed to create user %d", i+1),
// 		// 	})
// 		// }

// 		// Check if the request was successful
// 		// if res.StatusCode != http.StatusOK {
// 		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		// 		"message": fmt.Sprintf("failed to create user %d", i+1),
// 		// 	})
// 		// }
// 	}

// 	// Return a success response
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": fmt.Sprintf("successfully created %d random users", numUsers),
// 	})
// }
