package email

import (
	"fmt"
	"os"
)

// Configs 이메일 설정
type Configs struct {
	Host     string
	Port     string
	User     string
	Password string
}

// setConfigs 이메일 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		User:     os.Getenv("EMAIL_USER"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.Host == "" || configs.Port == "" || configs.User == "" || configs.Password == "" {
		return nil, fmt.Errorf("이메일 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
