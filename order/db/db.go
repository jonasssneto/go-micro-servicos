package db

import (
	"os"

	"github.com/go-redis/redis/v7"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return client
}
