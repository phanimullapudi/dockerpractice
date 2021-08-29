package redisconnect

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func ConnectDB() *redis.Client {
	log.Info("Settiing Connection to Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}
