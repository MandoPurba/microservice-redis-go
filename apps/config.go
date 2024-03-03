package apps

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func NewConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort:   8080,
	}

	if redisAddress, exists := os.LookupEnv("REDIS_ADDRESS"); exists {
		cfg.RedisAddress = redisAddress
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 64); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
