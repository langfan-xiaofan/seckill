package infra

import (
	"seckill/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		DB:       config.Conf.Redis.Database,
		Password: "",
	})
	return redisClient, nil
}
