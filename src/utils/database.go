package utils

import "github.com/go-redis/redis/v8"

func InitClient() *redis.Client {
	config := Config()

	client := redis.NewClient(&redis.Options{
		Addr: config.RedisDsn,
		DB:   1,
	})

	return client
}
// // TODO: make the same as in Jobs-api

// func RedisClient() -> возвращяет экземпляр

// func InitClient() -> инициализирует клиент синглтон

// CloseClient -> закрывает соединение
