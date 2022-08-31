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
	ServerDsn string
}

func Config() *Configuration {
	onceConfig.Do(func() {
		configuration = Configuration{
			ServerDsn: os.Getenv("SERVER_DSN"),
		}
	})
	return &configuration
}
