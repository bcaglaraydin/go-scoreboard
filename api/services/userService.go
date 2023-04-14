package services

import (
	"encoding/json"

	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
)

type UserServiceDB struct {
	rdb *redis.Client
}

//go:generate mockery --name UserService
type UserService interface {
	AddUserToLeaderboard(user *models.User) error
	GetUserFromUserID(userID string) (*models.User, error)
	SaveUser(user *models.User) error
	GetLeaderBoard(args ...string) ([]*models.User, error)
}

func NewUserServiceDB(rdb *redis.Client) UserServiceDB {
	return UserServiceDB{rdb: rdb}
}

func (u UserServiceDB) AddUserToLeaderboard(user *models.User) error {
	return u.rdb.ZAdd(database.Ctx, database.RedisLeaderboardKey, &redis.Z{
		Score:  float64(user.Points),
		Member: user.UserID.String(),
	}).Err()
}

func (u UserServiceDB) GetUserFromUserID(userID string) (*models.User, error) {
	userJson, err := u.rdb.HGet(database.Ctx, database.RedisUsersKey, userID).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return nil, err
	}

	userP, err := u.setUserPointAndRank(&user)
	if err != nil {
		return nil, err
	}
	return userP, nil
}

func (u UserServiceDB) SaveUser(user *models.User) error {
	user, err := u.setUserPointAndRank(user)
	if err != nil {
		return err
	}
	err = u.saveUserToHashSet(user)
	return err
}

func (u UserServiceDB) GetLeaderBoard(args ...string) ([]*models.User, error) {
	userIDs, err := u.getUsersBelowUser()
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0)
	for _, userID := range userIDs {
		user, err := u.GetUserFromUserID(userID)
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

func (u UserServiceDB) getUsersBelowUser() ([]string, error) {
	userIDs, err := u.rdb.ZRevRange(database.Ctx, database.RedisLeaderboardKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}

func (u UserServiceDB) setUserPointAndRank(user *models.User) (*models.User, error) {
	user, err := u.setUserRank(user)
	if err != nil {
		return nil, err
	}

	user, err = u.setUserPoint(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u UserServiceDB) setUserRank(user *models.User) (*models.User, error) {
	rank, err := u.rdb.ZRevRank(database.Ctx, database.RedisLeaderboardKey, user.UserID.String()).Result()
	if err != nil {
		return nil, err
	}
	user.Rank = int(rank) + 1
	return user, nil
}

func (u UserServiceDB) setUserPoint(user *models.User) (*models.User, error) {
	point, err := u.rdb.ZScore(database.Ctx, database.RedisLeaderboardKey, user.UserID.String()).Result()
	if err != nil {
		return nil, err
	}
	user.Points = int(point)
	return user, nil
}

func (u UserServiceDB) saveUserToHashSet(user *models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.rdb.HSet(database.Ctx, database.RedisUsersKey, user.UserID.String(), userBytes).Err()
}
