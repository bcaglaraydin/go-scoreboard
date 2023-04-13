package services

import (
	"github.com/bcaglaraydin/go-scoreboard/database"
	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
)

type ScoreServiceDB struct {
	rdb *redis.Client
}

type ScoreService interface {
	UpdateUserScore(score *models.Score) (float64, error)
}

func NewScoreServiceDB(rdb *redis.Client) ScoreServiceDB {
	return ScoreServiceDB{rdb: rdb}
}

func (s ScoreServiceDB) UpdateUserScore(score *models.Score) (float64, error) {
	newScore, err := s.rdb.ZIncrBy(database.Ctx, database.RedisLeaderboardKey, score.ScoreWorth, score.UserID).Result()
	if err != nil {
		return 0.0, err
	}
	return newScore, nil
}
