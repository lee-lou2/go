package redis

import (
	"fmt"
	"os"
)

// Configs Redis 설정
type Configs struct {
	redisHost     string
	redisPort     string
	redisPassword string
	redisDbNumber string
}

// setConfigs Redis 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		redisHost:     os.Getenv("REDIS_HOST"),
		redisPort:     os.Getenv("REDIS_PORT"),
		redisPassword: os.Getenv("REDIS_PASSWORD"),
		redisDbNumber: os.Getenv("REDIS_DB_NUMBER"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.redisHost == "" || configs.redisPort == "" || configs.redisPassword == "" || configs.redisDbNumber == "" {
		return nil, fmt.Errorf("Redis 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
