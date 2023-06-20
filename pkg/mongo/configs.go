package mongo

import (
	"fmt"
	"os"
)

// Configs MongoDB 설정
type Configs struct {
	Host     string
	Port     string
	UserName string
	Password string
}

// setConfigs MongoDB 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		Host:     os.Getenv("MONGO_HOST"),
		UserName: os.Getenv("MONGO_USER_NAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.Host == "" || configs.UserName == "" || configs.Password == "" {
		return nil, fmt.Errorf("MongoDB 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
