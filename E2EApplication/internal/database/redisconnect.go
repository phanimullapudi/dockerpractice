package redisconnect

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func NewRedisConnect(*redis.NewClient, error) {
	log.Info("Settiing Connection to Redis")

	_ = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}
