package infra

import (
	"fmt"
	"seckill/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Conf.Redis.Addr, config.Conf.Redis.Port),
		DB:       config.Conf.Redis.Database,
		Password: "",
	})
	return redisClient, nil
}
