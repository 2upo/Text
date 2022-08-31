package utils

import "github.com/go-redis/redis/v8"

func InitClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redisdb:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
