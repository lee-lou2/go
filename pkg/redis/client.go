package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type NewClient struct {
	*redis.Client
}

// Client Redis 클라이언트
func Client(configs ...Configs) (*NewClient, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf(
		"%s:%s",
		cfg.redisHost,
		cfg.redisPort,
	)
	db, _ := strconv.Atoi(cfg.redisDbNumber)
	client := &NewClient{redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: cfg.redisPassword,
		DB:       db,
	})}

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, nil
}
