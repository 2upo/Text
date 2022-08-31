package utils

import (
	"os"
	"sync"
)

var (
	configuration Configuration
	onceConfig    sync.Once
)

type Configuration struct {
	ServerDsn     string
	RedisPass     string
	RedisDsn      string
	RedisUsername string
}

func Config() *Configuration {
	onceConfig.Do(func() {
		configuration = Configuration{
			ServerDsn:     os.Getenv("SERVER_DSN"),
			RedisPass:     os.Getenv("REDIS_PASS"),
			RedisDsn:      os.Getenv("REDIS_DNS"),
			RedisUsername: os.Getenv("REDIS_USERNAME"),
		}
	})
	return &configuration
}
