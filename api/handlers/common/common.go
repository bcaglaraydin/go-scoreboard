package common

import (
	"encoding/json"

	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/helpers"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
)

func GetUserFromUserID(rdb *redis.Client, userID string) (*models.User, error) {
	userJson, err := rdb.HGet(database.Ctx, helpers.RedisUsersKey, userID).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return nil, err
	}

	userP, err := setUserPointAndRank(rdb, &user)
	if err != nil {
		return nil, err
	}
	return userP, nil
}

func SaveUser(rdb *redis.Client, user *models.User) error {
	user, err := setUserPointAndRank(rdb, user)
	if err != nil {
		return err
	}
	err = saveUserToHashSet(rdb, user)
	return err
}

func setUserPointAndRank(rdb *redis.Client, user *models.User) (*models.User, error) {
	user, err := setUserRank(rdb, user)
	if err != nil {
		return nil, err
	}

	user, err = setUserPoint(rdb, user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func setUserRank(rdb *redis.Client, user *models.User) (*models.User, error) {
	rank, err := rdb.ZRevRank(database.Ctx, helpers.RedisLeaderboardKey, user.UserID.String()).Result()
	if err != nil {
		return nil, err
	}
	user.Rank = int(rank) + 1
	return user, nil
}

func setUserPoint(rdb *redis.Client, user *models.User) (*models.User, error) {
	point, err := rdb.ZScore(database.Ctx, helpers.RedisLeaderboardKey, user.UserID.String()).Result()
	if err != nil {
		return nil, err
	}
	user.Points = int(point)
	return user, nil
}

func saveUserToHashSet(rdb *redis.Client, user *models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return rdb.HSet(database.Ctx, helpers.RedisUsersKey, user.UserID.String(), userBytes).Err()
}
